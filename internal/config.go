package internal

import (
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURL  string
	STMPHost     string
	STMPPort     string
	STMPUsername string
	STMPPassword string
}

func LoadConfig() Config {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	return Config{
		DatabaseURL:  viper.GetString("DATABASE_URL"),
		STMPHost:     viper.GetString("STMP_HOST"),
		STMPPort:     viper.GetString("STMP_PORT"),
		STMPUsername: viper.GetString("STMP_USERNAME"),
		STMPPassword: viper.GetString("STMP_PASSWORD"),
	}
}
