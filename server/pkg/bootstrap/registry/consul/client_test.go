package consul

import (
	confV1 "mosong/api/gen/go/conf/service/v1"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConsulRegistry(t *testing.T) {
	var cfg confV1.Registry
	cfg.Consul.Scheme = "http"
	cfg.Consul.Address = "localhost:8500"
	cfg.Consul.HealthCheck = false

	reg := NewRegistry(&cfg)
	assert.Nil(t, reg)
}
