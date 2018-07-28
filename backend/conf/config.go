package conf

import (
	"github.com/jinzhu/configor"
)

type Config struct {
	DB struct {
		Name     string `default:"sharqen"`
		User     string `default:"root" env:"DB_USER"`
		Password string `default:"" env:"DB_PASSWORD"`
		Host     string `default:"localhost" env:"DB_HOST"`
		Port     int    `default:"3306" env:"DB_PORT"`
	}
}

func GetConfig() Config {
	config := Config{}
	configor.Load(&config)

	return config
}
