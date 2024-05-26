//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"

	confV1 "mosong/api/gen/go/conf/service/v1"
	"mosong/app/admin/service/internal/server"
	"mosong/app/admin/service/internal/service"
)

// initApp init kratos application.
func initApp(log.Logger, registry.Registrar, *confV1.Bootstrap) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, service.ProviderSet, newApp))
}
