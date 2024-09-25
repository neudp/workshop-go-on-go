package useCase

import (
	"goOnGo/cmd/09-dependency-injection/model"
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

func (handler *GetCharacterHandler) Handle(query *GetCharacterQuery) (*model.Character, error) {
	handler.logger.Infof("Get character with id %d", query.id)

	return handler.client.GetCharacter(query.id)
}
