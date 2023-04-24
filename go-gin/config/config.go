package config

type Config struct {
	MysqlConfig `yaml:"mysql"`
	RedisConfig `yaml:"redis"`
	OSSConfig   `yaml:"oss"`
}

type MysqlConfig struct {
	Url      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

type RedisConfig struct {
	Url      string `yaml:"url"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

type OSSConfig struct {
	AccessKeyId     string `yaml:"accessKeyId"`
	AccessKeySecret string `yaml:"accessKeySecret"`
	EndPoint        string `yaml:"endPoint"`
	BucketName      string `yaml:"bucketName"`
}
