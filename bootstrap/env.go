package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	APP_ENV                   string `mapstructure:"APP_ENV"`
	SERVER_ADDRESS            string `mapstructure:"SERVER_ADDRESS"`
	CONTEXT_TIMEOUT           int    `mapstructure:"CONTEXT_TIMEOUT"`
	DB_HOST                   string `mapstructure:"DB_HOST"`
	DB_PORT                   string `mapstructure:"DB_PORT"`
	DB_USER                   string `mapstructure:"DB_USER"`
	DB_PASS                   string `mapstructure:"DB_PASS"`
	DB_NAME                   string `mapstructure:"DB_NAME"`
	ACCESS_TOKEN_EXPIRY_HOUR  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	REFRESH_TOKEN_EXPIRY_HOUR int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	ACCESS_TOKEN_SECRET       string `mapstructure:"ACCESS_TOKEN_SECRET"`
	REFRESH_TOKEN_SECRET      string `mapstructure:"REFRESH_TOKEN_SECRET"`
}

func NewEnv() *Env {
	env := Env{}

	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.APP_ENV == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
