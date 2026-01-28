package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort string
	DBHost  string
}

var Cfg Config

func Load() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("no config file found, using env only")
	}

	Cfg = Config{
		AppPort: viper.GetString("APP_PORT"),
		DBHost:  viper.GetString("DB_HOST"),
	}
}
