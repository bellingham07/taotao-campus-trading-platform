package initial

import (
	"com.xpwk/go-gin/internal/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func initMysql(config config.MysqlConfig) {
	username := config.Username //账号
	password := config.Password //密码
	url := config.Url           //数据库地址
	Dbname := config.Dbname     //数据库名
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, url, Dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("连接数据库失败：" + err.Error())
	}
	DB = db
}

func getDB() *gorm.DB {
	return DB
}
