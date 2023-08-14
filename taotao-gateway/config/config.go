package config

import (
	"gateway/registry"
	"gopkg.in/yaml.v3"
	"os"
)

var ProxyConfig map[string]string

type GatewayConf struct {
	Name     string `yaml:"Name"`
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
	ListenOn string `yaml:"ListenOn"`
}

type Conf struct {
	GatewayConf  GatewayConf   `yaml:"GatewayConf"`
	RegistryConf registry.Conf `yaml:"RegistryConf"`
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
