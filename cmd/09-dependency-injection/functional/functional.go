package functional

import (
	"goOnGo/internal/swapi-func/model/domain/character"
	"goOnGo/internal/swapi-func/use-case/swapi"
	"strconv"
)

func GetCharacter(id string) (*character.Character, error) {
	idInt, err := strconv.Atoi(id)

	if err != nil {
		return nil, err
	}

	query := swapi.NewGetCharacterQuery(idInt)

	return swapi.GetCharacter(query)
}
