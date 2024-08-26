package dotenv

import (
	"github.com/joho/godotenv"
	"goOnGo/internal/environment"
)

type Environment struct {
	Host string `env:"HOST" default:"localhost"`
	Port int    `env:"PORT" default:"8080"`
}

func ReadEnv() (*Environment, error) {
	env := new(Environment)

	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	if err := environment.Read(env); err != nil {
		return nil, err
	}

	return env, nil
}
