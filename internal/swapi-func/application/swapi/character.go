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

type toCharacter = func(dto *characterDto) (*character.Character, error)
type getCharacter = func(url string) (*character.Character, error)
type getPeople = func(url string) ([]*character.Character, error)

func newToCharacter(logError logging.Log, doGetPlanet getPlanet) toCharacter {
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

func newGetCharacter(logLevel logging.LogLevel, getResponse doGet) getCharacter {
	logError := logging.NewLog(logLevel, logging.Error)
	logInfo := logging.NewLog(logLevel, logging.Info)
	convertToCharacter := newToCharacter(logError, newGetPlanet(logLevel, getResponse))

	return func(url string) (*character.Character, error) {
		response, err := getResponse(url)
		if err != nil {
			logError(fmt.Sprintf("error getting character (%s): %v", url, err))

			return nil, err
		}

		logInfo("getCharacter response done")

		defer func() {
			err = errors.Join(err, response.Body.Close())
		}()

		if response.StatusCode != http.StatusOK {
			logError(fmt.Sprintf("error getting character (%s): %s", url, response.Status))

			return nil, err
		}

		logInfo("getCharacter response status OK")

		dto := new(characterDto)

		if err = json.NewDecoder(response.Body).Decode(dto); err != nil {
			logError(fmt.Sprintf("error decode character (%s): %v", url, err))

			return nil, err
		}

		logInfo("getCharacter decode done")

		chrctr, err := convertToCharacter(dto)

		if err != nil {
			logError(fmt.Sprintf("error toCharacter (%s): %v", url, err))

			return nil, err
		}

		logInfo("getCharacter toCharacter done")

		return chrctr, nil
	}
}

func newGetPeople(logLevel logging.LogLevel, getResponse doGet) getPeople {
	logError := logging.NewLog(logLevel, logging.Error)
	logInfo := logging.NewLog(logLevel, logging.Info)
	convertToCharacter := newToCharacter(logError, newGetPlanet(logLevel, getResponse))

	return func(url string) ([]*character.Character, error) {
		people := make([]*character.Character, 0)

		var next = &url
		for next != nil {
			response, err := getResponse(*next)
			if err != nil {
				logError(fmt.Sprintf("error getting people (%s): %v", *next, err))

				return nil, err
			}

			logInfo("getPeople response done")

			if response.StatusCode != http.StatusOK {
				logError(fmt.Sprintf("error getting people (%s): %s", *next, response.Status))

				return nil, err
			}

			logInfo("getPeople response status OK")

			dto := new(peopleDto)

			if err = json.NewDecoder(response.Body).Decode(dto); err != nil {
				logError(fmt.Sprintf("error decode people (%s): %v", *next, err))

				return nil, err
			}

			logInfo("getPeople decode done")

			for index, item := range dto.Results {
				var crctr *character.Character
				crctr, err = convertToCharacter(&item)

				if err != nil {
					logError(fmt.Sprintf("error toCharacter (%s)[%d]: %v", *next, index, err))

					return nil, err
				}

				logInfo("getPeople toCharacter done")

				people = append(people, crctr)
			}

			if err = response.Body.Close(); err != nil {
				logError(fmt.Sprintf("error close people (%s): %v", *next, err))

				return nil, err
			}

			next = dto.Next
		}

		return people, nil
	}
}
