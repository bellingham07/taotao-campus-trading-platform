package utils

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type Mysql struct {
	Dsn string
}

func InitXorm(dbtype string, mc Mysql) *xorm.Engine {
	engine, err := xorm.NewEngine(dbtype, mc.Dsn)
	fmt.Printf("[XORM CONNECTING] InitXorm DSN: %v\n", mc.Dsn)
	if err != nil {
		panic("[XORM ERROR] NewServiceContext 获取mysql连接错误 " + err.Error())
	}
	err = engine.Ping()
	if err != nil {
		panic("[XORM ERROR] NewServiceContext ping mysql 失败" + err.Error())
	}
	return engine
}
