package main

import (
	"encoding/json"
	"goOnGo/cmd/09-dependency-injection/functional"
	googleWire "goOnGo/cmd/09-dependency-injection/google-wire"
	uberFx "goOnGo/cmd/09-dependency-injection/uber-fx"
	"goOnGo/cmd/09-dependency-injection/vanila"
	"os"
)

/*
Dependency injection - это процесс предоставления зависимостей объекту.

В Go это можно сделать с помощью структур, в которых зависимости являются полями.
Или с помощью интерфейсов, которые описывают требуемые методы. Поскольку в Go
не существует наследования и как следствие полиморфизма, то способ
с использованием структур является более понятным, но, как и в классических ООП
языках, это может привести к проблеме с циклическими зависимостями и сложности
в тестировании.

Мы рассмотрим пример 3 способов реализации dependency injection:
- Ванильный метод ручного создания зависимостей
- Google Wire - библиотека от Google, по сути автоматизирующая ванильный метод,
  но с некоторыми ограничениями
- Uber FX - библиотека от Uber, так же основанная на рефлексии, но работающая
  с функциями вместо структур
- Функциональный метод (тот же ванильный метод, но с использованием функционального стиля)
*/

func main() {
	var result any
	var err error

	switch os.Args[1] {
	case "vanila":
		var app *vanila.App
		app, err = vanila.NewApp()

		if err != nil {
			panic(err)
		}

		result, err = app.Hadle(os.Args[2])
		if err != nil {
			panic(err)
		}
	case "wire":
		var app *googleWire.App
		app, err = googleWire.NewApp()

		if err != nil {
			panic(err)
		}

		result, err = app.Handle(os.Args[2])
		if err != nil {
			panic(err)
		}
	case "fx":
		err = uberFx.Do(func(app *uberFx.App) error {
			var err error
			result, err = app.Handle(os.Args[2])

			return err
		})

		if err != nil {
			panic(err)
		}
	case "func":
		result, err = functional.GetCharacter(os.Args[2])

		if err != nil {
			panic(err)
		}
	default:
		println("Invalid argument")
	}

	if result == nil {
		panic("character not found")
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(result); err != nil {
		panic(err)
	}
}
