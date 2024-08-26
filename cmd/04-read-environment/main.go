package main

import (
	"fmt"
	"goOnGo/cmd/04-read-environment/dotenv"
	envByCaarlos0 "goOnGo/cmd/04-read-environment/env-by-caarlos0"
	"goOnGo/cmd/04-read-environment/reflection"
	"goOnGo/cmd/04-read-environment/vanila"
	"os"
)

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
