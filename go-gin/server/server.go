package server

import (
	"com.xpdj/go-gin/config"
	ossLogic "com.xpdj/go-gin/logic/oss"
	mqLogic "com.xpdj/go-gin/logic/rabbitmq"
	"com.xpdj/go-gin/repository"
	"com.xpdj/go-gin/router"
	"com.xpdj/go-gin/utils/cache"
	"fmt"
	"github.com/yitter/idgenerator-go/idgen"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func Initializer() {
	yamlFile, err := os.ReadFile("config/config.yaml")
	if err != nil {
		panic("配置文件读取错误：" + err.Error())
	}
	_config := &config.Config{}
	err = yaml.Unmarshal(yamlFile, _config)
	if err != nil {
		panic("配置文件解析错误：" + err.Error())
	}
	var options = idgen.NewIdGeneratorOptions(20)
	fmt.Println("_config.MysqlConfig", _config.MysqlConfig)
	{
		go repository.InitMysql(_config.MysqlConfig)
		go cache.InitRedis(_config.RedisConfig)
		go ossLogic.InitOSS(_config.OSSConfig)
		go mqLogic.InitRabbitMQ(_config.RabbitMQConfig)
		go idgen.SetIdGenerator(options)
	}
}

func ListenAndServe(port string) {

	e := router.Routers()

	err := e.Run(":" + port)
	if err != nil {
		log.Printf("服务启动错误！ error：%s", err.Error())
	}
}
