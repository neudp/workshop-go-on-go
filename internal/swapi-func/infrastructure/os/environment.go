package os

import (
	"goOnGo/internal/environment"
	"goOnGo/internal/swapi-func/model/config"
	"goOnGo/internal/swapi-func/model/logging"
	"strings"
)

type appEnvValues struct {
	SwapiURL    string `env:"SWAPI_URL" default:"https://swapi.dev/api/"`
	MinLogLevel string `env:"MIN_LOG_LEVEL" default:"INFO"`
}

func ConfigFromEnv(overrides ...EnvironmentOverride) (*config.Config, error) {
	env := new(appEnvValues)

	if err := environment.Read(env); err != nil {
		return nil, err
	}

	for _, override := range overrides {
		env = override(env)
	}

	return config.New(
		env.SwapiURL,
		parseLogLevel(env.MinLogLevel),
	), nil
}

func parseLogLevel(level string) logging.Level {
	switch strings.ToUpper(level) {
	case "INFO", "INF", "I", "0":
		return logging.Info
	case "ERROR", "ERR", "E", "1":
		return logging.Error
	default:
		return logging.Info
	}
}
