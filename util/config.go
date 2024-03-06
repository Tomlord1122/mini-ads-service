package util

import "github.com/spf13/viper"

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	RedisAddr     string `mapstructure:"REDIS_ADDR"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	RedisDB       int    `mapstructure:"REDIS_DB"`
}

func LoadConfig(path string) (config Config, err error) {

	viper.AddConfigPath(path)

	viper.SetConfigName("app")

	viper.SetConfigType("env")

	// 讓 viper 自動讀取環境變量，環境變量將覆蓋配置文件中相同的鍵
	viper.AutomaticEnv()
	// 讀取配置文件到 viper 中
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	// 將讀取到的配置信息解析並存儲到 config 結構體中
	err = viper.Unmarshal(&config)
	return
}
