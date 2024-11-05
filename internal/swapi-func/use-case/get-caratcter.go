package use_case

import (
	getCharacter "goOnGo/internal/swapi-func/application/get-character"
	"goOnGo/internal/swapi-func/model/logging"
)

type GetCharacterQuery struct {
	IdValue int `json:"id"`
}

func (query *GetCharacterQuery) Id() int {
	return query.IdValue
}

type GetCharacter = func(query *GetCharacterQuery) (*CharacterDto, error)

func NewGetCharacter(find getCharacter.Find, logInfo logging.Log) GetCharacter {
	doGetCharacter := getCharacter.New(find, logInfo)

	return func(query *GetCharacterQuery) (*CharacterDto, error) {
		chrctr, err := doGetCharacter(query)

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
