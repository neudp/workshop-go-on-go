package useCase

import (
	model2 "goOnGo/internal/swapi/model"
)

type GetCharacterQuery struct {
	id int
}

func NewGetCharacterQuery(id int) *GetCharacterQuery {
	return &GetCharacterQuery{id: id}
}

type GetCharacterClient interface {
	GetCharacter(id int) (*model2.Character, error)
}

type GetCharacterHandler struct {
	client GetCharacterClient
	logger model2.Logger
}

func NewGetCharacterHandler(client GetCharacterClient, logger model2.Logger) *GetCharacterHandler {
	return &GetCharacterHandler{
		client: client,
		logger: logger,
	}
}

func (handler *GetCharacterHandler) Handle(query *GetCharacterQuery) (*model2.Character, error) {
	handler.logger.Infof("Get character with id %d", query.id)

	return handler.client.GetCharacter(query.id)
}
