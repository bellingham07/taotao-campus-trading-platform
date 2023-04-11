package initial

import (
	"com.xpwk/go-gin/cache"
	"com.xpwk/go-gin/config"
	"com.xpwk/go-gin/repository"
	"github.com/yitter/idgenerator-go/idgen"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
)

func Initializer() {
	file, err := os.ReadFile("a.txt")
	if err != nil || err != io.EOF {
		log.Printf("atxt %v", file)
		panic("配置文件读取错误2222：" + err.Error())
	}
	yamlFile, err := os.ReadFile("../config/config.yaml")
	if err != nil || err != io.EOF {
		log.Printf("sajkhdfsdkjfhsdkafjhsdk%v", yamlFile)
		panic("配置文件读取错误：" + err.Error())
	}
	var _config config.Config
	err = yaml.Unmarshal(yamlFile, &_config)
	if err != nil {
		panic("配置文件解析错误：" + err.Error())
	}
	repository.InitMysql(_config.MysqlConfig)
	cache.InitRedis(_config.RedisConfig)
	var options = idgen.NewIdGeneratorOptions(20)
	idgen.SetIdGenerator(options)
}
