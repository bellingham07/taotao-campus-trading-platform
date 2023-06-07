package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	MysqlConf struct {
		Dsn string
	}

	RedisConf struct {
		Addr     string
		Password string
		Db       int
	}

	Idgen struct {
		WorkerId uint16
	}
}
