package config

import (
	"github.com/caarlos0/env/v7"
)

type config struct {
	Permission                   string `env:"PERMISSION" envDefault:"user"`
	ApplicationName              string `env:"APPLICATION_NAME" envDefault:"app"`
	Domain                       string `env:"DOMAIN" envDefault:"localhost"`
	Environment                  string `env:"ENVIRONMENT" envDefault:"local"`
	DBName                       string `env:"DB_NAME" envDefault:"proto"`
	DBUser                       string `env:"DB_USER" envDefault:"proto"`
	DBPassword                   string `env:"DB_PASSWORD" envDefault:"password"`
	DBHost                       string `env:"DB_HOST" envDefault:"mysql:3306"` // ローカルでサーバを起動する場合は 127.0.0.1:3306 に変更
	RedisHost                    string `env:"REDIS_HOST" envDefault:"redis:6379"`
	GoogleApplicationCredentials string `env:"GOOGLE_APPLICATION_CREDENTIALS" envDefault:"./credentials/xxx.json"`
}

var EnvConfig = config{}

func init() {
	if err := env.Parse(&EnvConfig); err != nil {
		panic(err)
	}
}
