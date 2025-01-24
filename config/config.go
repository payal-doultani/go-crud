package config

import (
	"log"

	"github.com/caarlos0/env"
)

type Mysql struct {
	DB_HOST     string `env:"DB_HOST" envDefault:"localhost"`
	DB_USER     string `env:"DB_USER" envDefault:"payal"`
	DB_PASSWORD string `env:"DB_PASSWORD" envDefault:"payal"`
	DB_NAME     string `env:"DB_NAME" envDefault:"go-crud"`
	DB_PORT     string `env:"DB_PORT" envDefault:"3306"`
}

func LoadConfig() (*Mysql, error) {
	var cfg Mysql
	if err := env.Parse(&cfg); err != nil {
		log.Printf("Failed to load MySQL config: %v", err)
		return nil, err
	}
	return &cfg, nil
}
