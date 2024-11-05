package swapi

import (
	"encoding/json"
	"errors"
	"fmt"
	getCharacter "goOnGo/internal/swapi/application/get-character"
	"goOnGo/internal/swapi/model/domain/character"
	"goOnGo/internal/swapi/model/domain/color"
	"goOnGo/internal/swapi/model/logging"
	"net/http"
)

// interface definition
var _ getCharacter.Repository = new(CharactersClient)

type CharactersClient struct {
	client        *Client
	planetsClient *PlanetsClient
	logger        logging.Logger
}

func NewCharactersClient(client *Client, planetsClient *PlanetsClient, logger logging.Logger) *CharactersClient {
	return &CharactersClient{
		client:        client,
		planetsClient: planetsClient,
		logger:        logger,
	}
}

func (clnt *CharactersClient) GetCharacter(id int) (*character.Character, error) {
	return clnt.getCharacter(fmt.Sprintf("/api/people/%d/", id))
}

func (clnt *CharactersClient) GetPeople() ([]*character.Character, error) {
	return clnt.getPeople("/api/people/")
}

func (clnt *CharactersClient) FindById(id int) (*character.Character, error) {
	return clnt.GetCharacter(id)
}

func (clnt *CharactersClient) toCharacter(dto *characterDto) (*character.Character, error) {
	height, err := parseInt(dto.Height)
	if err != nil {
		clnt.logger.Error(fmt.Sprintf("error height: %v", err))

		return nil, err
	}

	mass, err := parseFloat(dto.Mass)
	if err != nil {
		clnt.logger.Error(fmt.Sprintf("error mass: %v", err))

		return nil, err
	}

	homeworld, err := clnt.planetsClient.getPlanet(dto.Homeworld)

	if err != nil {
		clnt.logger.Error(fmt.Sprintf("error homeworld: %v", err))

		return nil, err
	}

	return character.NewCharacter(
		dto.Name,
		height,
		mass,
		color.Color(dto.HairColor),
		color.Color(dto.SkinColor),
		color.Color(dto.EyeColor),
		character.BirthYear(dto.BirthYear),
		character.Gender(dto.Gender),
		homeworld,
	), nil
}

func (clnt *CharactersClient) getCharacter(url string) (*character.Character, error) {
	response, err := clnt.client.get(url)
	if err != nil {
		clnt.logger.Error(fmt.Sprintf("error getting character (%s): %v", url, err))

		return nil, err
	}

	clnt.logger.Info("getCharacter response done")

	defer func() {
		err = errors.Join(err, response.Body.Close())
	}()

	if response.StatusCode != http.StatusOK {
		clnt.logger.Error(fmt.Sprintf("error getting character (%s): %s", url, response.Status))

		return nil, err
	}

	clnt.logger.Info("getCharacter response status OK")

	dto := new(characterDto)

	if err = json.NewDecoder(response.Body).Decode(dto); err != nil {
		clnt.logger.Error(fmt.Sprintf("error decode character (%s): %v", url, err))

		return nil, err
	}

	clnt.logger.Info("getCharacter decode done")

	chrctr, err := clnt.toCharacter(dto)

	if err != nil {
		clnt.logger.Error(fmt.Sprintf("error toCharacter (%s): %v", url, err))

		return nil, err
	}

	clnt.logger.Info("getCharacter toCharacter done")

	return chrctr, nil
}

func (clnt *CharactersClient) getPeople(url string) ([]*character.Character, error) {
	characters := make([]*character.Character, 0)

	var next = &url
	for next != nil {
		response, err := clnt.client.get(*next)
		if err != nil {
			clnt.logger.Error(fmt.Sprintf("error getting people (%s): %v", *next, err))

			return nil, err
		}

		clnt.logger.Info("getPeople response done")

		if response.StatusCode != http.StatusOK {
			clnt.logger.Error(fmt.Sprintf("error getting people (%s): %s", *next, response.Status))

			return nil, err
		}

		clnt.logger.Info("getPeople response status OK")

		dto := new(peopleDto)

		if err = json.NewDecoder(response.Body).Decode(dto); err != nil {
			clnt.logger.Error(fmt.Sprintf("error decode people (%s): %v", *next, err))

			return nil, err
		}

		clnt.logger.Info("getPeople decode done")

		for index, item := range dto.Results {
			var chrctr *character.Character
			chrctr, err = clnt.toCharacter(&item)

			if err != nil {
				clnt.logger.Error(fmt.Sprintf("error toCharacter (%s)[%d]: %v", *next, index, err))

				return nil, err
			}

			clnt.logger.Info("getPeople toCharacter done")

			characters = append(characters, chrctr)
		}

		if err = response.Body.Close(); err != nil {
			clnt.logger.Error(fmt.Sprintf("error close people (%s): %v", *next, err))

			return nil, err
		}

		next = dto.Next
	}

	return characters, nil
}
