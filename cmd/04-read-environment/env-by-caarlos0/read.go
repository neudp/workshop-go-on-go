package envByCaarlos0

import "github.com/caarlos0/env"

type Environment struct {
	Host string `env:"HOST" envDefault:"localhost"`
	Port int    `env:"PORT" envDefault:"8080"`
}

func ReadEnv() (*Environment, error) {
	_env := new(Environment)

	if err := env.Parse(_env); err != nil {
		return nil, err
	}

	return _env, nil
}
