package environment

import (
	"goOnGo/internal/environment"
)

type appEnvValues struct {
	SwapiURL    string `env:"SWAPI_URL" default:"https://swapi.dev/api/"`
	MinLogLevel string `env:"MIN_LOG_LEVEL" default:"INFO"`
}

type Environment struct {
	values *appEnvValues
}

func Read(overrides ...Override) (*Environment, error) {
	env := new(appEnvValues)

	if err := environment.Read(env); err != nil {
		return nil, err
	}

	for _, override := range overrides {
		env = override(env)
	}

	return &Environment{values: env}, nil
}

func (env *Environment) SwapiURL() string {
	return env.values.SwapiURL
}

func (env *Environment) MinLogLevel() string {
	return env.values.MinLogLevel
}
