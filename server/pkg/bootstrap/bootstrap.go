package bootstrap

import (
	"fmt"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	kratosRegistry "github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport"

	confV1 "mosong/api/gen/go/conf/service/v1"
	"mosong/pkg/bootstrap/config"
	"mosong/pkg/bootstrap/logger"
	"mosong/pkg/bootstrap/registry"
	"mosong/pkg/bootstrap/tracer"
)

var (
	Service = config.NewServiceInfo(
		"",
		"1.0.0",
		"",
	)
)

type InitApp func(logger log.Logger, registrar kratosRegistry.Registrar, bootstrap *confV1.Bootstrap) (*kratos.App, func(), error)

// Bootstrap 应用引导启动
func Bootstrap(initApp InitApp, serviceName, version *string) {
	if serviceName != nil && len(*serviceName) != 0 {
		Service.Name = *serviceName
	}
	if version != nil && len(*version) != 0 {
		Service.Version = *version
	}

	// bootstrap
	cfg, ll, reg := DoBootstrap(Service)

	// init app
	app, cleanup, err := initApp(ll, reg, cfg)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// run the app.
	if err = app.Run(); err != nil {
		panic(err)
	}
}

// NewApp 创建 kratos 应用程序
func NewAPP(logger log.Logger, registrar kratosRegistry.Registrar, srv ...transport.Server) *kratos.App {
	return kratos.New(
		kratos.ID(Service.GetInstanceId()),
		kratos.Name(Service.Name),
		kratos.Version(Service.Version),
		kratos.Metadata(Service.Metadata),
		kratos.Logger(logger),
		kratos.Server(srv...),
		kratos.Registrar(registrar),
	)
}

// DoBootstrap 执行引导
func DoBootstrap(serviceInfo *config.ServiceInfo) (*confV1.Bootstrap, log.Logger, kratosRegistry.Registrar) {
	// inject command flags
	Flags := config.NewCommandFlags()
	Flags.Init()

	var err error

	// load configs
	if err = config.LoadBootstrapConfig(Flags.Conf); err != nil {
		panic(fmt.Sprintf("load config failed: %v", err))
	}

	// init logger
	ll := logger.NewLoggerProvider(config.GetBootstrapConfig().Logger, serviceInfo)

	// init registrar
	reg := registry.NewRegistry(config.GetBootstrapConfig().Registry)

	// init tracer
	if err = tracer.NewTracerProvider(config.GetBootstrapConfig().Tracer, serviceInfo); err != nil {
		panic(fmt.Sprintf("init tracer failed: %v", err))
	}

	return config.GetBootstrapConfig(), ll, reg
}
