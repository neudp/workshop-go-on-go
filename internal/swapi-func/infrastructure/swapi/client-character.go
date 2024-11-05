package swapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"goOnGo/internal/swapi-func/model/domain/character"
	"goOnGo/internal/swapi-func/model/domain/color"
	"goOnGo/internal/swapi-func/model/logging"
	"net/http"
)

type toCharacter func(dto *characterDto) (*character.Character, error)
type getCharacter func(url string) (*character.Character, error)
type getPeople func(url string) ([]*character.Character, error)

func newToCharacter(doGetPlanet getPlanet, logError logging.Log) toCharacter {
	return func(dto *characterDto) (*character.Character, error) {
		height, err := parseInt(dto.Height)
		if err != nil {
			logError(fmt.Sprintf("error height: %v", err))

			return nil, err
		}

		mass, err := parseFloat(dto.Mass)
		if err != nil {
			logError(fmt.Sprintf("error mass: %v", err))

			return nil, err
		}

		homeworld, err := doGetPlanet(dto.Homeworld)

		if err != nil {
			logError(fmt.Sprintf("error getPlanet: %v", err))

			return nil, err
		}

		return character.New(
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
}

func newGetCharacter(doGetRequest DoGetRequest, doGetPlanet getPlanet, logger *logging.Logger) getCharacter {
	convertToCharacter := newToCharacter(doGetPlanet, logger.Error)

	return func(url string) (chrctr *character.Character, err error) {
		response, err := doGetRequest(url)
		if err != nil {
			logger.Error(fmt.Sprintf("error getting character (%s): %v", url, err))

			fmt.Println("error getting character", err)

			return nil, err
		}

		logger.Info("getCharacter response done")

		defer func() {
			err = errors.Join(err, response.Body.Close())
		}()

		if response.StatusCode != http.StatusOK {
			logger.Error(fmt.Sprintf("error getting character (%s): %s", url, response.Status))

			return nil, err
		}

		logger.Info("getCharacter response status OK")

		dto := new(characterDto)

		if err = json.NewDecoder(response.Body).Decode(dto); err != nil {
			logger.Error(fmt.Sprintf("error decode character (%s): %v", url, err))

			return nil, err
		}

		logger.Info("getCharacter decode done")

		chrctr, err = convertToCharacter(dto)

		if err != nil {
			logger.Error(fmt.Sprintf("error toCharacter (%s): %v", url, err))

			return nil, err
		}

		logger.Info("getCharacter toCharacter done")

		return chrctr, nil
	}
}

func newGetPeople(doGetRequest DoGetRequest, doGetPlanet getPlanet, logger *logging.Logger) getPeople {
	convertToCharacter := newToCharacter(doGetPlanet, logger.Error)

	return func(url string) ([]*character.Character, error) {
		people := make([]*character.Character, 0)

		var next = &url
		for next != nil {
			response, err := doGetRequest(*next)
			if err != nil {
				logger.Error(fmt.Sprintf("error getting people (%s): %v", *next, err))

				return nil, err
			}

			logger.Info("getPeople response done")

			if response.StatusCode != http.StatusOK {
				logger.Error(fmt.Sprintf("error getting people (%s): %s", *next, response.Status))

				return nil, err
			}

			logger.Info("getPeople response status OK")

			dto := new(peopleDto)

			if err = json.NewDecoder(response.Body).Decode(dto); err != nil {
				logger.Error(fmt.Sprintf("error decode people (%s): %v", *next, err))

				return nil, err
			}

			logger.Info("getPeople decode done")

			for index, item := range dto.Results {
				var crctr *character.Character
				crctr, err = convertToCharacter(&item)

				if err != nil {
					logger.Error(fmt.Sprintf("error toCharacter (%s)[%d]: %v", *next, index, err))

					return nil, err
				}

				logger.Info("getPeople toCharacter done")

				people = append(people, crctr)
			}

			if err = response.Body.Close(); err != nil {
				logger.Error(fmt.Sprintf("error close people (%s): %v", *next, err))

				return nil, err
			}

			next = dto.Next
		}

		return people, nil
	}
}
