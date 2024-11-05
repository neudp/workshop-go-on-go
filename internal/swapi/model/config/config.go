package config

import (
	"goOnGo/internal/swapi/model/logging"
)

type Config struct {
	swapiURL    string
	minLoglevel logging.Level
}

func New(swapiURL string, minLoglevel logging.Level) *Config {
	return &Config{
		swapiURL:    swapiURL,
		minLoglevel: minLoglevel,
	}
}

func (config *Config) MinLoglevel() logging.Level {
	return config.minLoglevel
}

func (config *Config) SwapiURL() string {
	return config.swapiURL
}
