package initial

import (
	"com.xpwk/go-gin/cache"
	"com.xpwk/go-gin/config"
	ossLogic "com.xpwk/go-gin/logic/oss"
	"com.xpwk/go-gin/repository"
	"github.com/yitter/idgenerator-go/idgen"
	"gopkg.in/yaml.v3"
	"os"
)

func Initializer() {
	yamlFile, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		panic("配置文件读取错误：" + err.Error())
	}
	var _config config.Config
	err = yaml.Unmarshal(yamlFile, &_config)
	if err != nil {
		panic("配置文件解析错误：" + err.Error())
	}
	var options = idgen.NewIdGeneratorOptions(20)
	{
		go repository.InitMysql(_config.MysqlConfig)
		go cache.InitRedis(_config.RedisConfig)
		go ossLogic.InitOSS(_config.OSSConfig)
		go idgen.SetIdGenerator(options)
	}
}
