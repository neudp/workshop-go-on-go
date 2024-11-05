package character

import (
	"goOnGo/internal/swapi/model/domain/color"
	"goOnGo/internal/swapi/model/domain/planet"
)

type BirthYear string
type Gender string

type Character struct {
	name      string
	height    *int
	mass      *float64
	hairColor color.Color
	skinColor color.Color
	eyeColor  color.Color
	birthYear BirthYear
	gender    Gender
	homeworld *planet.Planet
}

func NewCharacter(
	name string,
	height *int,
	mass *float64,
	hairColor,
	skinColor,
	eyeColor color.Color,
	birthYear BirthYear,
	gender Gender,
	homeworld *planet.Planet,
) *Character {
	return &Character{
		name:      name,
		height:    height,
		mass:      mass,
		hairColor: hairColor,
		skinColor: skinColor,
		eyeColor:  eyeColor,
		birthYear: birthYear,
		gender:    gender,
		homeworld: homeworld,
	}
}

func (character *Character) Name() string {
	return character.name
}

func (character *Character) Height() *int {
	return character.height
}

func (character *Character) Mass() *float64 {
	return character.mass
}

func (character *Character) HairColor() color.Color {
	return character.hairColor
}

func (character *Character) SkinColor() color.Color {
	return character.skinColor
}

func (character *Character) EyeColor() color.Color {
	return character.eyeColor
}

func (character *Character) BirthYear() BirthYear {
	return character.birthYear
}

func (character *Character) Gender() Gender {
	return character.gender
}

func (character *Character) Homeworld() *planet.Planet {
	return character.homeworld
}
