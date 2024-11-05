package swapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"goOnGo/internal/swapi-func/model/domain/planet"
	"goOnGo/internal/swapi-func/model/logging"
	"net/http"
	"strings"
)

type getPlanet func(url string) (*planet.Planet, error)
type toPlanet func(dto *planetDto) (*planet.Planet, error)
type getPlanets func(url string) ([]*planet.Planet, error)

func newToPlanet(logError logging.Log) toPlanet {
	return func(dto *planetDto) (*planet.Planet, error) {
		rotationPeriod, err := parseInt(dto.RotationPeriod)

		if err != nil {
			logError(fmt.Sprintf("error rotationPeriod: %v", err))

			return nil, err
		}

		orbitalPeriod, err := parseInt(dto.OrbitalPeriod)

		if err != nil {
			logError(fmt.Sprintf("error orbitalPeriod: %v", err))

			return nil, err
		}

		diameter, err := parseInt(dto.Diameter)

		if err != nil {
			logError(fmt.Sprintf("error diameter: %v", err))

			return nil, err
		}

		terrains := strings.Split(dto.Terrain, ",")
		bioms := make([]planet.Biom, len(terrains))

		for i, terrain := range terrains {
			bioms[i] = planet.Biom(strings.TrimSpace(terrain))
		}

		surfaceWater, err := parseFloat(dto.SurfaceWater)

		if err != nil {
			logError(fmt.Sprintf("error surfaceWater: %v", err))

			return nil, err
		}

		population, err := parseInt(dto.Population)

		if err != nil {
			logError(fmt.Sprintf("error population: %v", err))

			return nil, err
		}

		return planet.New(
			dto.Name,
			rotationPeriod,
			orbitalPeriod,
			diameter,
			planet.Climate(dto.Climate),
			planet.Gravity(dto.Gravity),
			bioms,
			surfaceWater,
			population,
		), nil
	}
}

func newGetPlanet(doGetRequest DoGetRequest, logger *logging.Logger) getPlanet {
	convertToPlanet := newToPlanet(logger.Error)

	return func(url string) (*planet.Planet, error) {
		response, err := doGetRequest(url)
		if err != nil {
			logger.Error(fmt.Sprintf("error getting planet (%s): %v", url, err))

			return nil, err
		}

		logger.Info("getPlanet response done")

		defer func() {
			err = errors.Join(err, response.Body.Close())
		}()

		if response.StatusCode != http.StatusOK {
			logger.Error(fmt.Sprintf("error getting planet (%s): %s", url, response.Status))

			return nil, err
		}

		logger.Info("getPlanet response status OK")

		dto := new(planetDto)
		if err = json.NewDecoder(response.Body).Decode(dto); err != nil {
			logger.Error(fmt.Sprintf("error decode planet (%s): %v", url, err))

			return nil, err
		}

		logger.Info("getPlanet decode done")

		plnt, err := convertToPlanet(dto)

		if err != nil {
			logger.Error(fmt.Sprintf("error toPlanet (%s): %v", url, err))

			return nil, err
		}

		logger.Info("getPlanet toPlanet done")

		return plnt, nil
	}
}

func newGetPlanets(doGetRequest DoGetRequest, logger *logging.Logger) getPlanets {
	convertToPlanet := newToPlanet(logger.Error)

	return func(url string) ([]*planet.Planet, error) {
		planets := make([]*planet.Planet, 0)

		var next = &url
		for next != nil {
			response, err := doGetRequest(*next)
			if err != nil {
				logger.Error(fmt.Sprintf("error getting planets (%s): %v", *next, err))

				return nil, err
			}

			logger.Info("getPlanets response done")

			if response.StatusCode != http.StatusOK {
				logger.Error(fmt.Sprintf("error getting planets (%s): %s", *next, response.Status))

				return nil, err
			}

			logger.Info("getPlanets response status OK")

			dto := new(planetsDto)

			if err = json.NewDecoder(response.Body).Decode(dto); err != nil {
				logger.Error(fmt.Sprintf("error decode planets (%s): %v", *next, err))

				return nil, err
			}

			logger.Info("getPlanets decode done")

			for index, item := range dto.Results {
				var plnt *planet.Planet
				plnt, err = convertToPlanet(&item)

				if err != nil {
					logger.Error(fmt.Sprintf("error toPlanet (%s)[%d]: %v", *next, index, err))

					return nil, err
				}

				logger.Info("getPlanets toPlanet done")

				planets = append(planets, plnt)
			}

			if err = response.Body.Close(); err != nil {
				logger.Error(fmt.Sprintf("error close planets (%s): %v", *next, err))

				return nil, err
			}

			next = dto.Next
		}

		return planets, nil
	}
}
