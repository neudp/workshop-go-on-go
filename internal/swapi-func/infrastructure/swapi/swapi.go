package swapi

import (
	"fmt"
	"goOnGo/internal/swapi-func/model/domain/character"
	"goOnGo/internal/swapi-func/model/domain/planet"
	"goOnGo/internal/swapi-func/model/logging"
)

type GetPlanets func() ([]*planet.Planet, error)
type GetPlanet func(id int) (*planet.Planet, error)
type GetCharacter func(id int) (*character.Character, error)
type GetPeople func() ([]*character.Character, error)

func NewGetPlanets(doGetRequest DoGetRequest, logger *logging.Logger) GetPlanets {
	doGetPlanets := newGetPlanets(doGetRequest, logger)

	return func() ([]*planet.Planet, error) {
		return doGetPlanets("/api/planets/")
	}
}

func NewGetPlanet(doGetRequest DoGetRequest, logger *logging.Logger) GetPlanet {
	doGetPlanet := newGetPlanet(doGetRequest, logger)

	return func(id int) (*planet.Planet, error) {
		return doGetPlanet(fmt.Sprintf("/api/planets/%d/", id))
	}
}

func NewGetCharacter(doGetRequest DoGetRequest, logger *logging.Logger) GetCharacter {
	doGetCharacter := newGetCharacter(doGetRequest, newGetPlanet(doGetRequest, logger), logger)

	return func(id int) (*character.Character, error) {
		return doGetCharacter(fmt.Sprintf("/api/people/%d/", id))
	}
}

func NewGetPeople(doRequest DoGetRequest, logger *logging.Logger) GetPeople {
	doGetPeople := newGetPeople(doRequest, newGetPlanet(doRequest, logger), logger)

	return func() ([]*character.Character, error) {
		return doGetPeople("/api/people/")
	}
}
