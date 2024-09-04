package main

import (
	"fmt"
	"goOnGo/cmd/04-read-environment/dotenv"
	envByCaarlos0 "goOnGo/cmd/04-read-environment/env-by-caarlos0"
	"goOnGo/cmd/04-read-environment/reflection"
	"goOnGo/cmd/04-read-environment/vanila"
	"os"
)

/*
Чтение переменных окружения - это must-have для сервисов распределенных систем.
В Go есть внутренний пакет os, который позволяет работать с переменными окружения,
а также с другими системными свойствами.
Однако, в Go нет встроенного метода для чтения переменных окружения в готовую структуру.
Данный пример демонстрирует различные способы чтения переменных окружения в структуру.
- vanila - самый простой способ, который использует встроенный пакет os
- carlos0-lib - способ, использующий библиотеку от caarlos0
- reflection - способ, использующий рефлексию (технически это вкратце реализация библиотеки от caarlos0)
- dotenv - способ, использующий библиотеку от joho/godotenv для чтения переменных из .env файла
*/
func main() {
	var readEnv func() (any, error)
	switch os.Args[1] {
	case "vanila":
		readEnv = func() (any, error) {
			return vanila.ReadEnv()
		}
	case "carlos0-lib":
		readEnv = func() (any, error) {
			return envByCaarlos0.ReadEnv()
		}
	case "reflection":
		readEnv = func() (any, error) {
			return reflection.ReadEnv()
		}
	case "dotenv":
		readEnv = func() (any, error) {
			return dotenv.ReadEnv()
		}
	default:
		fmt.Println("Invalid argument")
	}

	env, err := readEnv()

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%+v\n", env)
	}
}
