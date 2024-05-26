package registry

import (
	"github.com/go-kratos/kratos/v2/registry"

	confV1 "mosong/api/gen/go/conf/service/v1"
	"mosong/pkg/bootstrap/registry/consul"
)

// NewRegistry 创建一个注册客户端
func NewRegistry(cfg *confV1.Registry) registry.Registrar {
	if cfg == nil {
		return nil
	}

	switch Type(cfg.Type) {
	case Consul:
		return consul.NewRegistry(cfg)
	}

	return nil
}

// NewDiscovery 创建一个发现客户端
func NewDiscovery(cfg *confV1.Registry) registry.Discovery {
	if cfg == nil {
		return nil
	}

	switch Type(cfg.Type) {
	case Consul:
		return consul.NewRegistry(cfg)
	}

	return nil
}
