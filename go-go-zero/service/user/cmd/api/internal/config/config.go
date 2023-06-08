package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	Mysql struct {
		Dsn string
	}

	Redis struct {
		Addr     string
		Password string
		Db       int
	}

	Idgen struct {
		WorkerId uint16
	}
}
