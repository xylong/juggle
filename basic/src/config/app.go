package config

// Config 配置
type Config struct {
	AppName  string
	LogLevel string

	Mysql MysqlConfig
	Redis RedisConfig
}
