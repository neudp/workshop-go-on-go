package envByCaarlos0

import "github.com/caarlos0/env"

/**
github.com/caarlos0/env - это библиотека для чтения переменных окружения в структуру.
Она использует теги структуры для определения переменных окружения.
Тег env используется для указания имени переменной окружения.
Тег envDefault используется для указания значения по умолчанию.

env.Parse() принимает указатель на структуру и заполняет ее значениями из переменных окружения
в соответствии с тегами
*/

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
