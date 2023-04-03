package config

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	DbHost       string `mapstructure:"DATABASE_HOST"`
	DbName       string `mapstructure:"DATABASE_NAME"`
	DbUser       string `mapstructure:"DATABASE_USERNAME"`
	DbPassword   string `mapstructure:"DATABASE_PASSWORD"`
	DbPort       string `mapstructure:"DATABASE_PORT"`
	JwtSecretKey string `mapstucture:"JWT_SECRET_KEY"`
}

func LoadConfig() (config Config, err error) {
	mode := os.Getenv("APP_ENV")

	viper.AddConfigPath(".")
	if mode == "prod" {
		viper.SetConfigFile(".env")
	} else {
		viper.SetConfigFile(".env-dev.env")
	}
	
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
