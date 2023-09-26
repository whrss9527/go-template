package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go-template/internal/conf"
	"go-template/internal/pkg/trace"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	zaplog "go-template/internal/pkg/log"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.1
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
	)
}

func customMiddleware(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		if tr, ok := transport.FromServerContext(ctx); ok {
			fmt.Println("operation:", tr.Operation())
		}
		reply, err = handler(ctx, req)
		return
	}
}

func main() {
	// 配置，启动链路追踪
	Name = "go-template"
	id = "go-template"
	Version = "test-V0.0.1"
	traceConf := trace.NewConf(Name, id, Version)
	tp, _ := traceConf.TracerProvider()
	otel.SetTracerProvider(tp) // 为全局链路追踪
	flag.Parse()
	encoder := zapcore.EncoderConfig{
		TimeKey:    "t",
		LevelKey:   "level",
		NameKey:    "logger",
		CallerKey:  "caller",
		MessageKey: "msg",
		// StacktraceKey: "stack",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		// EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	zlogger := zaplog.NewZapLogger(
		"./log.txt",
		encoder,
		zap.NewAtomicLevelAt(zapcore.DebugLevel),
		zap.AddStacktrace(
			zap.NewAtomicLevelAt(zapcore.ErrorLevel)),
		zap.AddCaller(),
		zap.AddCallerSkip(0),
		zap.Development(),
	)

	// 原有的输出到控制台上
	// logger := log.With(log.NewStdLogger(os.Stdout),
	// 输出到日志中
	logger := log.With(zlogger,
		// "ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		// "service.id", id,
		// "service.name", Name,
		// "service.version", Version,
		"trace_id", tracing.TraceID(),
		// "span_id", tracing.SpanID(),
	)

	log.SetLogger(logger)
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(bc.Server, bc.Biz, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// 连接数据库
	//db.InitDbConfig(bc.Data)
	//if err != nil {
	//	panic(err)
	//}

	//redis.Init(bc.Data)

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}

}
