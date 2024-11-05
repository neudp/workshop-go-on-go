package getCharacter

import (
	"fmt"
	"goOnGo/internal/swapi/model/domain/character"
	"goOnGo/internal/swapi/model/logging"
)

type Query interface {
	Id() int
}

type Repository interface {
	FindById(id int) (*character.Character, error)
}

type Handler struct {
	client Repository
	logger logging.Logger
}

func NewHandler(repository Repository, logger logging.Logger) *Handler {
	return &Handler{
		client: repository,
		logger: logger,
	}
}

func (handler *Handler) Handle(query Query) (*character.Character, error) {
	id := query.Id()

	handler.logger.Info(fmt.Sprintf("Get character with id %d", id))

	return handler.client.FindById(id)
}
