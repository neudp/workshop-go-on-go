package getCharacter

import (
	"fmt"
	"goOnGo/internal/swapi-func/model/domain/character"
	"goOnGo/internal/swapi-func/model/logging"
)

type Query interface {
	Id() int
}

type GetCharacter func(query Query) (*character.Character, error)
type Find func(id int) (*character.Character, error)

func New(find Find, logInfo logging.Log) GetCharacter {
	return func(query Query) (*character.Character, error) {
		id := query.Id()

		logInfo(fmt.Sprintf("Looking for character with id %d", id))

		return find(id)
	}
}
