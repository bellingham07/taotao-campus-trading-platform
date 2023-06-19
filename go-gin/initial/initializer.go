package initial

import (
	"com.xpdj/go-gin/config"
	ossLogic "com.xpdj/go-gin/logic/oss"
	"com.xpdj/go-gin/logic/rabbitmq"
	"com.xpdj/go-gin/repository"
	"com.xpdj/go-gin/utils/cache"
	"fmt"
	"github.com/yitter/idgenerator-go/idgen"
	"gopkg.in/yaml.v3"
	"os"
)

func Initializer() {
	yamlFile, err := os.ReadFile("./config/config.yaml")
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
