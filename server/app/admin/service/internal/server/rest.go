package server

import (
	"mosong/app/admin/service/cmd/server/assets"
	"mosong/app/admin/service/internal/service"
	"mosong/pkg/bootstrap/rpc"
	swaggerUI "mosong/pkg/swagger-ui"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"

	adminV1 "mosong/api/gen/go/admin/service/v1"
	confV1 "mosong/api/gen/go/conf/service/v1"
)

// NewRESTServer new an HTTP server.
func NewRESTServer(
	cfg *confV1.Bootstrap, logger log.Logger,
	helloSvc *service.HelloService,
) *http.Server {
	srv := rpc.CreateRestServer(cfg, newRestMiddleware(logger)...)

	adminV1.RegisterHelloServiceHTTPServer(srv, helloSvc)

	if cfg.GetServer().GetRest().GetEnableSwagger() {
		swaggerUI.RegisterSwaggerUIServerWithOption(
			srv,
			swaggerUI.WithTitle("Kratos巨石应用实践"),
			swaggerUI.WithMemoryData(assets.OpenApiData, "yaml"),
			swaggerUI.WithBasePath("/docs/"),
		)
	}

	return srv
}

// NewMiddleware 创建中间件
func newRestMiddleware(logger log.Logger) []middleware.Middleware {
	var ms []middleware.Middleware
	ms = append(ms, logging.Server(logger))
	ms = append(ms, tracing.Server())
	return ms
}
