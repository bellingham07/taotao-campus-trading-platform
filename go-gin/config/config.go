package config

type Config struct {
	MysqlConfig `yaml:"mysql"`
	RedisConfig `yaml:"redis"`
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
}
