package useCase

import (
	"goOnGo/internal/swapi/model"
)

type GetCharacterQuery struct {
	id int
}

func NewGetCharacterQuery(id int) *GetCharacterQuery {
	return &GetCharacterQuery{id: id}
}

type GetCharacterClient interface {
	GetCharacter(id int) (*model.Character, error)
}

type GetCharacterHandler struct {
	client GetCharacterClient
	logger model.Logger
}

func NewGetCharacterHandler(client GetCharacterClient, logger model.Logger) *GetCharacterHandler {
	return &GetCharacterHandler{
		client: client,
		logger: logger,
	}
}

func (handler *GetCharacterHandler) Handle(query *GetCharacterQuery) (*CharacterDto, error) {
	handler.logger.Infof("Get character with id %d", query.id)

	chrctr, err := handler.client.GetCharacter(query.id)

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
