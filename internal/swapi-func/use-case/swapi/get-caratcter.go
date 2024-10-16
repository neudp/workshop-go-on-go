package swapi

import (
	appContext "goOnGo/internal/swapi-func/infrastructure/app-context"
	"goOnGo/internal/swapi-func/model/domain/character"
	"goOnGo/internal/swapi-func/swapi"
)

type GetCharacterQuery struct {
	id int
}

func NewGetCharacterQuery(id int) *GetCharacterQuery {
	return &GetCharacterQuery{id: id}
}

func GetCharacter(ctx *appContext.AppContext, query *GetCharacterQuery) (*character.Character, error) {
	return swapi.GetCharacter(ctx.SwapiContext(), query.id)
}
