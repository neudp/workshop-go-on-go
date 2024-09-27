package swapi

import (
	"encoding/json"
	"errors"
	"fmt"
	model2 "goOnGo/internal/swapi/model"
	"net/http"
	"strings"
)

type Transport interface {
	Do(request *http.Request) (*http.Response, error)
}

type Swapi struct {
	transport Transport
	logger    model2.Logger
}

func New(transport Transport, logger model2.Logger) *Swapi {
	return &Swapi{transport: transport, logger: logger}
}

func (swapi *Swapi) GetPlanets() ([]*model2.Planet, error) {
	return swapi.getPlanets("https://swapi.dev/api/planets/")
}

func (swapi *Swapi) GetPlanet(id int) (*model2.Planet, error) {
	return swapi.getPlanet(fmt.Sprintf("https://swapi.dev/api/planets/%d/", id))
}

func (swapi *Swapi) GetCharacter(id int) (*model2.Character, error) {
	return swapi.getCharacter(fmt.Sprintf("https://swapi.dev/api/people/%d/", id))
}

func (swapi *Swapi) GetPeople() ([]*model2.Character, error) {
	return swapi.getPeople("https://swapi.dev/api/people/")
}

func (swapi *Swapi) get(url string) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		swapi.logger.Errorf("error creating request for %s: %v", url, err)

		return nil, err
	}

	return swapi.transport.Do(request)
}

func (swapi *Swapi) getPlanet(url string) (planet *model2.Planet, err error) {
	response, err := swapi.get(url)
	if err != nil {
		swapi.logger.Errorf("error getting planet (%s): %v", url, err)

		return nil, err
	}

	swapi.logger.Infof("getPlanet response done")

	defer func() {
		err = errors.Join(err, response.Body.Close())
	}()

	if response.StatusCode != http.StatusOK {
		swapi.logger.Errorf("error getting planet (%s): %s", url, response.Status)

		return nil, err
	}

	swapi.logger.Infof("getPlanet response status OK")

	dto := new(planetDto)
	if err = json.NewDecoder(response.Body).Decode(dto); err != nil {
		swapi.logger.Errorf("error decode planet (%s): %v", url, err)

		return nil, err
	}

	swapi.logger.Infof("getPlanet decode done")

	planet, err = swapi.toPlanet(dto)

	if err != nil {
		swapi.logger.Errorf("error toPlanet (%s): %v", url, err)

		return nil, err
	}

	swapi.logger.Infof("getPlanet toPlanet done")

	return planet, nil
}

func (swapi *Swapi) toPlanet(dto *planetDto) (*model2.Planet, error) {
	rotationPeriod, err := parseInt(dto.RotationPeriod)

	if err != nil {
		swapi.logger.Errorf("error rotationPeriod: %v", err)

		return nil, err
	}

	orbitalPeriod, err := parseInt(dto.OrbitalPeriod)

	if err != nil {
		swapi.logger.Errorf("error orbitalPeriod: %v", err)

		return nil, err
	}

	diameter, err := parseInt(dto.Diameter)

	if err != nil {
		swapi.logger.Errorf("error diameter: %v", err)

		return nil, err
	}

	terrains := strings.Split(dto.Terrain, ",")
	bioms := make([]model2.Biom, len(terrains))

	for i, terrain := range terrains {
		bioms[i] = model2.Biom(strings.TrimSpace(terrain))
	}

	surfaceWater, err := parseFloat(dto.SurfaceWater)

	if err != nil {
		swapi.logger.Errorf("error surfaceWater: %v", err)

		return nil, err
	}

	population, err := parseInt(dto.Population)

	if err != nil {
		swapi.logger.Errorf("error population: %v", err)

		return nil, err
	}

	return model2.NewPlanet(
		dto.Name,
		rotationPeriod,
		orbitalPeriod,
		diameter,
		model2.Climate(dto.Climate),
		model2.Gravity(dto.Gravity),
		bioms,
		surfaceWater,
		population,
	), nil
}

func (swapi *Swapi) getPlanets(url string) ([]*model2.Planet, error) {
	planets := make([]*model2.Planet, 0)

	var next = &url
	for next != nil {
		response, err := swapi.get(*next)
		if err != nil {
			swapi.logger.Errorf("error getting planets (%s): %v", *next, err)

			return nil, err
		}

		swapi.logger.Infof("getPlanets response done")

		if response.StatusCode != http.StatusOK {
			swapi.logger.Errorf("error getting planets (%s): %s", *next, response.Status)

			return nil, err
		}

		swapi.logger.Infof("getPlanets response status OK")

		dto := new(planetsDto)

		if err = json.NewDecoder(response.Body).Decode(dto); err != nil {
			swapi.logger.Errorf("error decode planets (%s): %v", *next, err)

			return nil, err
		}

		swapi.logger.Infof("getPlanets decode done")

		for index, item := range dto.Results {
			var planet *model2.Planet
			planet, err = swapi.toPlanet(&item)

			if err != nil {
				swapi.logger.Errorf("error toPlanet (%s)[%d]: %v", *next, index, err)

				return nil, err
			}

			swapi.logger.Infof("getPlanets toPlanet done")

			planets = append(planets, planet)
		}

		if err = response.Body.Close(); err != nil {
			swapi.logger.Errorf("error close planets (%s): %v", *next, err)

			return nil, err
		}

		next = dto.Next
	}

	return planets, nil
}

func (swapi *Swapi) getCharacter(url string) (character *model2.Character, err error) {
	response, err := swapi.get(url)
	if err != nil {
		swapi.logger.Errorf("error getting character (%s): %v", url, err)

		return nil, err
	}

	swapi.logger.Infof("getCharacter response done")

	defer func() {
		err = errors.Join(err, response.Body.Close())
	}()

	if response.StatusCode != http.StatusOK {
		swapi.logger.Errorf("error getting character (%s): %s", url, response.Status)

		return nil, err
	}

	swapi.logger.Infof("getCharacter response status OK")

	dto := new(characterDto)

	if err = json.NewDecoder(response.Body).Decode(dto); err != nil {
		swapi.logger.Errorf("error decode character (%s): %v", url, err)

		return nil, err
	}

	swapi.logger.Infof("getCharacter decode done")

	character, err = swapi.toCharacter(dto)

	if err != nil {
		swapi.logger.Errorf("error toCharacter (%s): %v", url, err)

		return nil, err
	}

	swapi.logger.Infof("getCharacter toCharacter done")

	return character, nil
}

func (swapi *Swapi) toCharacter(dto *characterDto) (*model2.Character, error) {
	height, err := parseInt(dto.Height)
	if err != nil {
		swapi.logger.Errorf("error height: %v", err)

		return nil, err
	}

	mass, err := parseFloat(dto.Mass)
	if err != nil {
		swapi.logger.Errorf("error mass: %v", err)

		return nil, err
	}

	homeworld, err := swapi.getPlanet(dto.Homeworld)

	if err != nil {
		swapi.logger.Errorf("error homeworld: %v", err)

		return nil, err
	}

	return model2.NewCharacter(
		dto.Name,
		height,
		mass,
		model2.Color(dto.HairColor),
		model2.Color(dto.SkinColor),
		model2.Color(dto.EyeColor),
		model2.BirthYear(dto.BirthYear),
		model2.Gender(dto.Gender),
		homeworld,
	), nil
}

func (swapi *Swapi) getPeople(url string) ([]*model2.Character, error) {
	characters := make([]*model2.Character, 0)

	var next = &url
	for next != nil {
		response, err := swapi.get(*next)
		if err != nil {
			swapi.logger.Errorf("error getting people (%s): %v", *next, err)

			return nil, err
		}

		swapi.logger.Infof("getPeople response done")

		if response.StatusCode != http.StatusOK {
			swapi.logger.Errorf("error getting people (%s): %s", *next, response.Status)

			return nil, err
		}

		swapi.logger.Infof("getPeople response status OK")

		dto := new(peopleDto)

		if err = json.NewDecoder(response.Body).Decode(dto); err != nil {
			swapi.logger.Errorf("error decode people (%s): %v", *next, err)

			return nil, err
		}

		swapi.logger.Infof("getPeople decode done")

		for index, item := range dto.Results {
			var character *model2.Character
			character, err = swapi.toCharacter(&item)

			if err != nil {
				swapi.logger.Errorf("error toCharacter (%s)[%d]: %v", *next, index, err)

				return nil, err
			}

			swapi.logger.Infof("getPeople toCharacter done")

			characters = append(characters, character)
		}

		if err = response.Body.Close(); err != nil {
			swapi.logger.Errorf("error close people (%s): %v", *next, err)

			return nil, err
		}

		next = dto.Next
	}

	return characters, nil
}
