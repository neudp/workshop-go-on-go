package swapi

import (
	"goOnGo/internal/swapi-func/application/swapi"
	"goOnGo/internal/swapi-func/model/logging"
)

type GetCharacterQuery struct {
	Id int `json:"id"`
}

type GetCharacter = func(query *GetCharacterQuery) (*CharacterDto, error)

func NewGetCharacter(logLevel logging.LogLevel, request swapi.DoRequest) GetCharacter {
	getSwapiCharacter := swapi.NewGetCharacter(logLevel, request)

	return func(query *GetCharacterQuery) (*CharacterDto, error) {
		chrctr, err := getSwapiCharacter(query.Id)

		if err != nil {
			return nil, err
		}

		return &CharacterDto{
			Name:      chrctr.Name(),
			Height:    chrctr.Height(),
			Mass:      chrctr.Mass(),
			HairColor: string(chrctr.HairColor()),
			SkinColor: string(chrctr.SkinColor()),
			EyeColor:  string(chrctr.EyeColor()),
			BirthYear: string(chrctr.BirthYear()),
			Gender:    string(chrctr.Gender()),
			Homeworld: chrctr.Homeworld().Name(),
		}, nil
	}
}
