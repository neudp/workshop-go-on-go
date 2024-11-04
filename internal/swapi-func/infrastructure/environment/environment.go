package environment

import (
	"errors"
	"goOnGo/internal/environment"
	"goOnGo/internal/swapi-func/model/logging"
)

var ErrInvalidLogLevelValue = errors.New("invalid log level value")

type appEnvValues struct {
	SwapiURL    string `env:"SWAPI_URL" default:"https://swapi.dev/api/"`
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

func (env *Environment) SwapiURL() string {
	return env.values.SwapiURL
}

func (env *Environment) MinLogLevel() (logging.Level, error) {
	switch env.values.MinLogLevel {
	case "INFO":
		return logging.Info, nil
	case "ERROR":
		return logging.Error, nil
	default:
		return -1, ErrInvalidLogLevelValue
	}
}
