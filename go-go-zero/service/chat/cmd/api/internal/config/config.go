package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	Mysql struct {
		Dsn string
	}

	Mongo struct {
		Url string
	}

	Idgen struct {
		WorkerId uint16
	}
}
