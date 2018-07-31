package config

import "github.com/gobuffalo/envy"

type Configuration struct {
	DatabaseURL string
}

var config *Configuration

func GetConfig() *Configuration {
	if config == nil {
		config = &Configuration{
			DatabaseURL: envy.Get("DATABASE_URL", "postgres://postgres:@postgres:5432/postgres?sslmode=disable"),
		}
	}
	return config
}
