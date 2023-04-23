package main

import (
	"flag"
	"go-template/internal/conf"
	"go-template/internal/pkg/db"
	"go-template/internal/pkg/redis"
	"os"

	logdef "github.com/go-kratos/kratos/contrib/log/logrus/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/sirupsen/logrus"
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

func main() {
	flag.Parse()

	logFmt := logrus.New()
	logFmt.Formatter = &logrus.JSONFormatter{}
	logFmt.SetLevel(logrus.InfoLevel)
	logger := logdef.NewLogger(logFmt)
	logger = log.With(logger,
		"caller", log.DefaultCaller,
	)
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()
	logFmt.Info(logrus.Fields{
		"common": "this is a common filed",
		"other":  "i also should be logged always",
	})
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
	db.InitDbConfig(bc.Data)
	if err != nil {
		panic(err)
	}

	redis.Init(bc.Data)

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}

}
