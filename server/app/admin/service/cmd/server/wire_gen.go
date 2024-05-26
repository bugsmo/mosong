// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"mosong/api/gen/go/conf/service/v1"
	"mosong/app/admin/service/internal/server"
	"mosong/app/admin/service/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(logger log.Logger, registrar registry.Registrar, bootstrap *servicev1.Bootstrap) (*kratos.App, func(), error) {
	helloService := service.NewHelloService()
	httpServer := server.NewRESTServer(bootstrap, logger, helloService)
	app := newApp(logger, registrar, httpServer)
	return app, func() {
	}, nil
}
