package main

import (
	"encoding/json"
	"errors"
	"goOnGo/cmd/09-dependency-injection/functional"
	googleWire "goOnGo/cmd/09-dependency-injection/google-wire"
	uberFx "goOnGo/cmd/09-dependency-injection/uber-fx"
	"goOnGo/cmd/09-dependency-injection/vanila"
	"goOnGo/internal/swapi/model"
	"os"
	"strings"
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
*/

type GetCharacterApp interface {
	GetCharacter(id string) (*model.Character, error)
}

func GetCharacterCommand(app GetCharacterApp) error {
	character, err := app.GetCharacter(os.Args[2])

	if err != nil {
		return err
	}

	if character == nil {
		return errors.New("character not found")
	}

	terrains := make([]string, len(character.Homeworld().Terrains()))

	for i, terrain := range character.Homeworld().Terrains() {
		terrains[i] = string(terrain)
	}

	dto := CharacterDto{
		Name:      character.Name(),
		Height:    character.Height(),
		Mass:      character.Mass(),
		HairColor: string(character.HairColor()),
		SkinColor: string(character.SkinColor()),
		EyeColor:  string(character.EyeColor()),
		BirthYear: string(character.BirthYear()),
		Gender:    string(character.Gender()),
		Homeworld: HomeworldDto{
			Name:    character.Homeworld().Name(),
			Climate: string(character.Homeworld().Climate()),
			Terrain: strings.Join(terrains, ", "),
		},
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")

	if err = encoder.Encode(dto); err != nil {
		return err
	}

	return nil
}

func main() {
	switch os.Args[1] {
	case "vanila":
		app, err := vanila.NewApp()

		if err != nil {
			panic(err)
		}

		if err := GetCharacterCommand(app); err != nil {
			panic(err)
		}
	case "wire":
		app, err := googleWire.NewApp()

		if err != nil {
			panic(err)
		}

		if err := GetCharacterCommand(app); err != nil {
			panic(err)
		}
	case "fx":
		err := uberFx.Run(func(app *uberFx.App) error {
			return GetCharacterCommand(app)
		})

		if err != nil {
			panic(err)
		}
	case "func":
		character, err := functional.GetCharacter(os.Args[2])

		if err != nil {
			panic(err)
		}

		if character == nil {
			panic("character not found")
		}

		terrains := make([]string, len(character.Homeworld().Terrains()))

		for i, terrain := range character.Homeworld().Terrains() {
			terrains[i] = string(terrain)
		}

		dto := CharacterDto{
			Name:      character.Name(),
			Height:    character.Height(),
			Mass:      character.Mass(),
			HairColor: string(character.HairColor()),
			SkinColor: string(character.SkinColor()),
			EyeColor:  string(character.EyeColor()),
			BirthYear: string(character.BirthYear()),
			Gender:    string(character.Gender()),
			Homeworld: HomeworldDto{
				Name:    character.Homeworld().Name(),
				Climate: string(character.Homeworld().Climate()),
				Terrain: strings.Join(terrains, ", "),
			},
		}

		encoder := json.NewEncoder(os.Stdout)
		encoder.SetIndent("", "  ")

		if err = encoder.Encode(dto); err != nil {
			panic(err)
		}
	default:
		println("Invalid argument")
	}

}
