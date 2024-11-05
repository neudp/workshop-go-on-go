package useCase

import (
	"goOnGo/internal/swapi/application/get-character"
	"goOnGo/internal/swapi/model/logging"
)

type GetCharacterQuery struct {
	IdValue int `json:"id"`
}

func (query *GetCharacterQuery) Id() int {
	return query.IdValue
}

type GetCharacterHandler struct {
	getCharacter *getCharacter.Handler
}

func NewGetCharacterHandler(repository getCharacter.Repository, logger logging.Logger) *GetCharacterHandler {
	return &GetCharacterHandler{
		getCharacter: getCharacter.NewHandler(repository, logger),
	}
}

func (handler *GetCharacterHandler) Handle(query *GetCharacterQuery) (*CharacterDto, error) {
	chrctr, err := handler.getCharacter.Handle(query)

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
