package dotenv

import (
	"github.com/joho/godotenv"
	"goOnGo/internal/environment"
)

/**
godotenv.Load() - функция, которая загружает переменные окружения из переданных файлов
непосредственно в переменные окружения процесса. Переменные окружения системы
при этом не меняются.

Важно учитывать, что переменные окружения которые уже были установлены для процесса
не будут перезаписаны, то есть при вызове godotenv.Load(".local.env", ".env") приоритет
будет у переменных из os.Getenv() затем из .local.env и в конце из .env

Поскольку godotenv загружает переменные из файла в переменные окружения процесса,
то после вызова godotenv.Load() необходимо перечитать переменные окружения в структуру.
*/

type Environment struct {
	Host string `env:"HOST" default:"localhost"`
	Port int    `env:"PORT" default:"8080"`
}

func ReadEnv() (*Environment, error) {
	env := new(Environment)

	if err := godotenv.Load(".env", ".local.env"); err != nil {
		return nil, err
	}

	if err := environment.Read(env); err != nil {
		return nil, err
	}

	return env, nil
}
