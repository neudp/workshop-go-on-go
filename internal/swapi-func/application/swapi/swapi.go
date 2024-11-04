package swapi

import (
	"fmt"
	"goOnGo/internal/swapi-func/model/domain/character"
	"goOnGo/internal/swapi-func/model/domain/planet"
	"goOnGo/internal/swapi-func/model/logging"
)

type GetPlanets = func() ([]*planet.Planet, error)
type GetPlanet = func(id int) (*planet.Planet, error)
type GetCharacter = func(id int) (*character.Character, error)
type GetPeople = func() ([]*character.Character, error)

func NewGetPlanets(logLevel logging.LogLevel, doRequest DoRequest) GetPlanets {
	logError := logging.NewLog(logLevel, logging.Error)
	doGetPlanets := newGetPlanets(logLevel, newGet(logError, doRequest))

	return func() ([]*planet.Planet, error) {
		return doGetPlanets("https://swapi.dev/api/planets/")
	}
}

func NewGetPlanet(logLevel logging.LogLevel, doRequest DoRequest) GetPlanet {
	logError := logging.NewLog(logLevel, logging.Error)
	doGetPlanet := newGetPlanet(logLevel, newGet(logError, doRequest))

	return func(id int) (*planet.Planet, error) {
		return doGetPlanet(fmt.Sprintf("https://swapi.dev/api/planets/%d/", id))
	}
}

func NewGetCharacter(logLevel logging.LogLevel, doRequest DoRequest) GetCharacter {
	logError := logging.NewLog(logLevel, logging.Error)
	doGetCharacter := newGetCharacter(logLevel, newGet(logError, doRequest))

	return func(id int) (*character.Character, error) {
		return doGetCharacter(fmt.Sprintf("https://swapi.dev/api/people/%d/", id))
	}
}

func NewGetPeople(logLevel logging.LogLevel, doRequest DoRequest) GetPeople {
	logError := logging.NewLog(logLevel, logging.Error)
	doGetPeople := newGetPeople(logLevel, newGet(logError, doRequest))

	return func() ([]*character.Character, error) {
		return doGetPeople("https://swapi.dev/api/people/")
	}
}
