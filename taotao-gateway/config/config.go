package config

import (
	"gateway/gateway"
	"gateway/registry"
	"gopkg.in/yaml.v3"
	"os"
)

var ProxyConfig map[string]string

type Conf struct {
	GatewayConf  gateway.Conf
	RegistryConf registry.Conf
}

func MustLoad(path string, v any) {
	content, err := os.ReadFile(path)
	if err != nil {
		panic("[FATAL] 配置文件打开错误：" + err.Error())
	}

	if err = yaml.Unmarshal(content, v); err != nil {
		panic("[FATAL] 配置文件解析错误：" + err.Error())
	}
}
