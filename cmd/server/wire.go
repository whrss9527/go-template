//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"go-template/internal/biz"
	"go-template/internal/conf"
	"go-template/internal/data/mysql"
	"go-template/internal/interfaces"
	"go-template/internal/server"
	"go-template/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Biz, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, biz.ProviderSet, service.ProviderSet, mysql.ProviderSet, interfaces.ProviderSet, newApp))
}
