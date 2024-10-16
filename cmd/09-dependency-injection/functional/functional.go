package functional

import (
	appContext "goOnGo/internal/swapi-func/infrastructure/app-context"
	"goOnGo/internal/swapi-func/model/domain/character"
	"goOnGo/internal/swapi-func/use-case/swapi"
	"strconv"
)

func GetCharacter(id string) (*character.Character, error) {
	idInt, err := strconv.Atoi(id)

	if err != nil {
		return nil, err
	}

	ctx, err := appContext.New()

	if err != nil {
		return nil, err
	}

	return swapi.GetCharacter(ctx, swapi.NewGetCharacterQuery(idInt))
}
