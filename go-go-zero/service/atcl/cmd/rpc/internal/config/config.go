package config

import (
	"go-go-zero/common/config"
)

type Config struct {
	Consul    config.Consul `yaml:"Consul"`
	TaoTaoRpc config.TaoTaoRpc
}
