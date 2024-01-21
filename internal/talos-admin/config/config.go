package config

import (
	"github.com/caarlos0/env/v10"
)

type Сonfig struct {
	LISTEN_ADDR string `env:"LISTEN_ADDR" envDefault:"127.0.0.1:8080"`
	LOG_LEVEL   string `env:"LOG_LEVEL" envDefault:"info"`
	LOG_FORMAT  string `env:"LOG_FORMAT" envDefault:"text"`
	DB_HOST     string `env:"DB_HOST" envDefault:"127.0.0.1"`
	DB_PORT     int    `env:"DB_PORT" envDefault:"5432"`
	DB_NAME     string `env:"DB_NAME"`
	DB_USER     string `env:"DB_USER"`
	DB_PASS     string `env:"DB_PASS"`
	// Port         int            `env:"PORT" envDefault:"3000"`
	// Password     string         `env:"PASSWORD,unset"`
	// IsProduction bool           `env:"PRODUCTION"`
	// Duration     time.Duration  `env:"DURATION"`
	// Hosts        []string       `env:"HOSTS" envSeparator:":"`
	// TempFolder   string         `env:"TEMP_FOLDER,expand" envDefault:"${HOME}/tmp"`
	// StringInts   map[string]int `env:"MAP_STRING_INT"`
}

func Get() (Сonfig, error) {
	cfg := Сonfig{}

	if err := env.Parse(&cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}
