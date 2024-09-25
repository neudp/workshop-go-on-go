package config

import (
	"goOnGo/cmd/09-dependency-injection/model"
	"goOnGo/internal/environment"
)

type Config struct {
	SwapiURL string `env:"SWAPI_URL" default:"https://swapi.dev/api/"`
}

func Build(logger model.Logger) (*Config, error) {
	env := new(Config)

	if err := environment.Read(env); err != nil {
		logger.Errorf("Failed to load config: %v", err)

		return nil, err
	}
	logger.Infof("Config loaded")

	return env, nil
}
