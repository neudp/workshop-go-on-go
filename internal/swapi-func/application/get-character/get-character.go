package getCharacter

import (
	"fmt"
	"goOnGo/internal/swapi-func/model/domain/character"
	"goOnGo/internal/swapi-func/model/logging"
)

type Query interface {
	Id() int
}

var cache = make(map[int]*character.Character)

type GetCharacter func(query Query) (*character.Character, error)
type Find func(id int) (*character.Character, error)

func New(find Find, logInfo logging.Log) GetCharacter {
	return func(query Query) (*character.Character, error) {
		id := query.Id()
		logInfo(fmt.Sprintf("Looking for character with id %d", id))

		var chrctr *character.Character
		var ok bool
		if chrctr, ok = cache[id]; ok {
			return chrctr, nil
		}

		var err error
		chrctr, err = find(id)
		if err != nil {
			return nil, err
		}

		cache[id] = chrctr

		return chrctr, nil
	}
}
