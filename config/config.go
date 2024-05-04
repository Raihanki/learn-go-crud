package config

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

type Config struct {
	APP_NAME string `mapstructure:"APP_NAME"`
	APP_PORT string `mapstructure:"APP_PORT"`
	APP_ENV  string `mapstructure:"APP_ENV"`

	DB_CONNECTION string `mapstructure:"DB_CONNECTION"`
	DB_HOST       string `mapstructure:"DB_HOST"`
	DB_PORT       string `mapstructure:"DB_PORT"`
	DB_DATABASE   string `mapstructure:"DB_DATABASE"`
	DB_USERNAME   string `mapstructure:"DB_USERNAME"`
	DB_PASSWORD   string `mapstructure:"DB_PASSWORD"`
}

var ENV Config

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	errRead := viper.ReadInConfig()
	if errRead != nil {
		log.Fatal("Error reading config file")
	}

	errUnmarshal := viper.Unmarshal(&ENV)
	if errUnmarshal != nil {
		log.Fatal("Unable to decode into struct")
	}
}
