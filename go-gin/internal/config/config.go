package config

type Config struct {
	MysqlConfig
	RedisConfig
}

type MysqlConfig struct {
	Url      string
	Username string
	Password string
	Dbname   string
}

type RedisConfig struct {
	Url      string
	Password string
}
