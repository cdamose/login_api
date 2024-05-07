package config

import (
	"fmt"

	"github.com/caarlos0/env/v7"
)

type Config struct {
	IsProduction  bool   `env:"PRODUCTION"`
	MYSQLHost     string `env:"MYSQL_ADDR"`
	MYSQLUser     string `env:"MYSQL_USER"`
	MYSQLPassword string `env:"MYSQL_PASSWORD"`
	MYSQLDatabase string `env:"MYSQL_DATABASE"`
	BASE_URL      string `env:"BASE_URL"`
	Port          string `env:"PORT"`
}

func InitConfig() Config {
	cfg := Config{}
	err := env.Parse(&cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", err)
	return cfg

}
