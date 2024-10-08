package vanila

import (
	"os"
	"strconv"
)

/*
os.Getenv() - функция, которая возвращает значение переменной окружения по ключу.
Если переменной с таким ключом нет, то возвращается пустая строка.
*/

type Environment struct {
	Host string
	Port int
}

func ReadEnv() (*Environment, error) {
	port, err := ReadInt("PORT")

	if err != nil {
		return nil, err
	}

	return &Environment{
		Host: ReadString("HOST"),
		Port: port,
	}, nil
}

func ReadString(key string) string {
	return os.Getenv(key)
}

func ReadInt(key string) (int, error) {
	strValue := ReadString(key)

	return strconv.Atoi(strValue)
}
