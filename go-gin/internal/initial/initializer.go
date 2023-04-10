package initial

import (
	"com.xpwk/go-gin/internal/config"
	"gopkg.in/yaml.v3"
	"os"
)

func Initializer() {
	yamlFile, err := os.ReadFile("../../etc/config.yaml")
	if err != nil {
		panic("配置文件错误：" + err.Error())
	}
	var _config *config.Config
	err = yaml.Unmarshal(yamlFile, _config)
	if err != nil {
		panic("配置文件解析错误：" + err.Error())
	}
	initMysql(_config.MysqlConfig)
	initRedis(_config.RedisConfig)
}
