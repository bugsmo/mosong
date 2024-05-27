package config

import confV1 "mosong/api/gen/go/conf/service/v1"

var configList []interface{}

var commonConfig = &confV1.Bootstrap{}

func GetBootstrapConfig() *confV1.Bootstrap {
	return commonConfig
}

// RegisterConfig 注册配置
func RegisterConfig(c interface{}) {
	initBootstrapConfig()
	configList = append(configList, c)
}

func initBootstrapConfig() {
	if len(configList) > 0 {
		return
	}

	configList = append(configList, commonConfig)

	if commonConfig.Server == nil {
		commonConfig.Server = &confV1.Server{}
		configList = append(configList, commonConfig.Server)
	}

	if commonConfig.Data == nil {
		commonConfig.Data = &confV1.Data{}
		configList = append(configList, commonConfig.Data)
	}

	if commonConfig.Logger == nil {
		commonConfig.Logger = &confV1.Logger{}
		configList = append(configList, commonConfig.Logger)
	}

	if commonConfig.Registry == nil {
		commonConfig.Registry = &confV1.Registry{}
		configList = append(configList, commonConfig.Registry)
	}

	if commonConfig.Tracer == nil {
		commonConfig.Tracer = &confV1.Tracer{}
		configList = append(configList, commonConfig.Tracer)
	}
}
