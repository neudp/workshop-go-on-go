package environment

import (
	"errors"
	"goOnGo/internal/environment"
	"goOnGo/internal/swapi/model/config"
	"goOnGo/internal/swapi/model/logging"
)

var ErrInvalidLogLevelValue = errors.New("invalid log level value")

type appEnvValues struct {
	SwapiURL    string `env:"SWAPI_URL" default:"https://swapi.dev"`
	MinLogLevel string `env:"MIN_LOG_LEVEL" default:"INFO"`
}

type Environment struct {
	values *appEnvValues
}

func Read() (*Environment, error) {
	env := new(appEnvValues)

	if err := environment.Read(env); err != nil {
		return nil, err
	}

	return &Environment{values: env}, nil
}

func (env *Environment) ToConfig() (*config.Config, error) {
	minLogLevel, err := parseLogLevel(env.values.MinLogLevel)

	if err != nil {
		return nil, err
	}

	return config.New(env.values.SwapiURL, minLogLevel), nil
}

func parseLogLevel(level string) (logging.Level, error) {
	switch level {
	case "INFO":
		return logging.Info, nil
	case "ERROR":
		return logging.Error, nil
	default:
		return logging.Info, ErrInvalidLogLevelValue
	}
}
