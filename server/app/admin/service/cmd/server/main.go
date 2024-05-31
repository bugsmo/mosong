package main

import (
	"mosong/pkg/service"

	"github.com/tx7do/go-utils/trans"

	"github.com/bugsmo/kratos-bootstrap"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/http"
)

var version string

// go build -ldflags "-X main.version=x.y.z"

// wire.go
func newApp(ll log.Logger, rr registry.Registrar, hs *http.Server) *kratos.App {
	return bootstrap.NewAPP(ll, rr, hs)
}

func main() {
	bootstrap.Bootstrap(initApp, trans.Ptr(service.AdminService), trans.Ptr(version))
}
