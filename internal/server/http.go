package server

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	e "go-template/api/template_proj/errors"
	v1 "go-template/api/template_proj/v1"
	"go-template/internal/conf"
	redisClient "go-template/internal/pkg/redis"
	"go-template/internal/service"
	"go-template/internal/util"
	"strconv"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtv4 "github.com/golang-jwt/jwt/v4"
)

func NewWhiteListMatcher() selector.MatchFunc {

	whiteList := make(map[string]struct{})
	// rules： /package.service/function in your .proto file
	whiteList["/template_proj.v1.TemplateProj/Login"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(cs *conf.Server, cb *conf.Biz, service *service.TemplateService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			tracing.Server(), // 链路追踪
			recovery.Recovery(),
			DetermineMinimumVersion(cb),
			selector.Server(
				jwt.Server(
					func(token *jwtv4.Token) (interface{}, error) {
						key, _ := base64.StdEncoding.DecodeString(cs.Http.JwtSecret)
						return key, nil
					},
					jwt.WithSigningMethod(jwtv4.SigningMethodHS512),
					jwt.WithClaims(func() jwtv4.Claims {
						var claims = &util.AccountClaims{}

						return claims
					})),

				ratelimit.Server(),
				UserLock(),
				SetUserInfoToContext(),
			).
				Match(NewWhiteListMatcher()).
				Build(),
			CustomerLogger(logger)),
	}
	if cs.Http.Network != "" {
		opts = append(opts, http.Network(cs.Http.Network))
	}
	if cs.Http.Addr != "" {
		opts = append(opts, http.Address(cs.Http.Addr))
	}
	if cs.Http.Timeout != nil {
		opts = append(opts, http.Timeout(cs.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterTemplateProjHTTPServer(srv, service)
	return srv
}
func SetUserInfoToContext() middleware.Middleware {

	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			var simbaId string
			var source string
			if claims, ok := jwt.FromContext(ctx); ok {
				simbaId = strings.Split(claims.(*util.AccountClaims).Subject, "_")[1]
				source = strings.Split(claims.(*util.AccountClaims).Subject, "_")[0]
			}
			rdb := redisClient.RedisClient

			key := "user:" + simbaId
			userId, _ := rdb.Get(ctx, key).Result()
			ctx = context.WithValue(ctx, "userId", userId)
			ctx = context.WithValue(ctx, "source", source)
			return handler(ctx, req)
		}
	}
}
func DetermineMinimumVersion(cb *conf.Biz) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				if ht, ok := tr.(*http.Transport); ok {
					clientVersion := ht.RequestHeader().Get("client_version")
					if CompareVersions(clientVersion, cb.MinClientVersion) < 0 {
						return nil, e.ErrorClientVersionNeedsToBeUpgraded("客户端版本过低，请升级客户端")
					}
				}
			}
			return handler(ctx, req)
		}
	}
}
func CompareVersions(clientVersion string, minClientVersion string) int {
	// 将版本号按 "." 分割为数组
	clientVersions := strings.Split(clientVersion, ".")
	minClientVersions := strings.Split(minClientVersion, ".")

	// 取两个版本号中长度较小的那个作为循环次数
	minLen := len(clientVersions)
	if len(minClientVersions) < minLen {
		minLen = len(minClientVersions)
	}

	for i := 0; i < minLen; i++ {
		// 将分割出来的每一段转换成数字进行比较
		num1, _ := strconv.Atoi(clientVersions[i])
		num2, _ := strconv.Atoi(minClientVersions[i])

		if num1 > num2 {
			return 1 // clientVersion 大于 minClientVersion
		} else if num1 < num2 {
			return -1 // clientVersion 小于 minClientVersion
		}
	}

	// 如果以上条件都不满足，则说明两个版本号相等
	return 0 // clientVersion 等于 minClientVersion
}

func UserLock() middleware.Middleware {

	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			var simbaId string
			var source string
			if claims, ok := jwt.FromContext(ctx); ok {
				simbaId = strings.Split(claims.(*util.AccountClaims).Subject, "_")[1]
				source = strings.Split(claims.(*util.AccountClaims).Subject, "_")[0]
			}

			rdb := redisClient.RedisClient
			pool := goredis.NewPool(rdb)
			rs := redsync.New(pool)

			var lockKey = "user-lock:" + simbaId + "-" + source + "-lock"
			mutex := rs.NewMutex(lockKey)

			if err := mutex.Lock(); err != nil {
				return nil, e.ErrorRequestBusy("请求繁忙")
			}

			defer func() {
				if ok, err := mutex.Unlock(); !ok || err != nil {
					fmt.Println("unlock failed")
					panic("unlock failed")
				}
			}()

			return handler(ctx, req)
		}
	}
}
func CustomerLogger(logger log.Logger) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			var (
				code      int32
				reason    string
				kind      string
				operation string
			)
			startTime := time.Now()
			if info, ok := transport.FromServerContext(ctx); ok {
				kind = info.Kind().String()
				operation = info.Operation()
			}
			reply, err = handler(ctx, req)
			if se := errors.FromError(err); se != nil {
				code = se.Code
				reason = se.Reason
			}
			userId := ctx.Value("userId")
			level, stack := extractError(err)
			_ = log.WithContext(ctx, logger).Log(
				level,
				"userId", userId,
				"kind", "server",
				"component", kind,
				"operation", operation,
				"args", extractArgs(req),
				"code", code,
				"reason", reason,
				"stack", stack,
				"latency", time.Since(startTime).Seconds(),
			)
			return
		}
	}
}

// extractArgs returns the string of the req
func extractArgs(req interface{}) string {
	if stringer, ok := req.(fmt.Stringer); ok {
		return stringer.String()
	}
	return fmt.Sprintf("%+v", req)
}

// extractError returns the string of the error
func extractError(err error) (log.Level, string) {
	if err != nil {
		return log.LevelError, fmt.Sprintf("%+v", err)
	}
	return log.LevelInfo, ""
}
