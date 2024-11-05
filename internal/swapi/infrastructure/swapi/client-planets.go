package swapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"goOnGo/internal/swapi/model/domain/planet"
	"goOnGo/internal/swapi/model/logging"
	"net/http"
	"strings"
)

type PlanetsClient struct {
	client *Client
	logger logging.Logger
}

func NewPlanetsClient(client *Client, logger logging.Logger) *PlanetsClient {
	return &PlanetsClient{
		client: client,
		logger: logger,
	}
}

func (clnt *PlanetsClient) GetPlanet(id int) (*planet.Planet, error) {
	return clnt.getPlanet(fmt.Sprintf("/api/planets/%d/", id))
}

func (clnt *PlanetsClient) GetPlanets() ([]*planet.Planet, error) {
	return clnt.getPlanets("/api/planets/")
}

func (clnt *PlanetsClient) toPlanet(dto *planetDto) (*planet.Planet, error) {
	rotationPeriod, err := parseInt(dto.RotationPeriod)

	if err != nil {
		clnt.logger.Error(fmt.Sprintf("error rotationPeriod: %v", err))

		return nil, err
	}

	orbitalPeriod, err := parseInt(dto.OrbitalPeriod)

	if err != nil {
		clnt.logger.Error(fmt.Sprintf("error orbitalPeriod: %v", err))

		return nil, err
	}

	diameter, err := parseInt(dto.Diameter)

	if err != nil {
		clnt.logger.Error(fmt.Sprintf("error diameter: %v", err))

		return nil, err
	}

	terrains := strings.Split(dto.Terrain, ",")
	bioms := make([]planet.Biom, len(terrains))

	for i, terrain := range terrains {
		bioms[i] = planet.Biom(strings.TrimSpace(terrain))
	}

	surfaceWater, err := parseFloat(dto.SurfaceWater)

	if err != nil {
		clnt.logger.Error(fmt.Sprintf("error surfaceWater: %v", err))

		return nil, err
	}

	population, err := parseInt(dto.Population)

	if err != nil {
		clnt.logger.Error(fmt.Sprintf("error population: %v", err))

		return nil, err
	}

	return planet.NewPlanet(
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

func (clnt *PlanetsClient) getPlanet(url string) (*planet.Planet, error) {
	response, err := clnt.client.get(url)
	if err != nil {
		clnt.logger.Error(fmt.Sprintf("error getting planet (%s): %v", url, err))

		return nil, err
	}

	clnt.logger.Info("getPlanet response done")

	defer func() {
		err = errors.Join(err, response.Body.Close())
	}()

	if response.StatusCode != http.StatusOK {
		clnt.logger.Error(fmt.Sprintf("error getting planet (%s): %s", url, response.Status))

		return nil, err
	}

	clnt.logger.Info("getPlanet response status OK")

	dto := new(planetDto)
	if err = json.NewDecoder(response.Body).Decode(dto); err != nil {
		clnt.logger.Error(fmt.Sprintf("error decode planet (%s): %v", url, err))

		return nil, err
	}

	clnt.logger.Info("getPlanet decode done")

	plnt, err := clnt.toPlanet(dto)

	if err != nil {
		clnt.logger.Error(fmt.Sprintf("error toPlanet (%s): %v", url, err))

		return nil, err
	}

	clnt.logger.Info("getPlanet toPlanet done")

	return plnt, nil
}

func (clnt *PlanetsClient) getPlanets(url string) ([]*planet.Planet, error) {
	planets := make([]*planet.Planet, 0)

	var next = &url
	for next != nil {
		response, err := clnt.client.get(*next)
		if err != nil {
			clnt.logger.Error(fmt.Sprintf("error getting planets (%s): %v", *next, err))

			return nil, err
		}

		clnt.logger.Info("getPlanets response done")

		if response.StatusCode != http.StatusOK {
			clnt.logger.Error(fmt.Sprintf("error getting planets (%s): %s", *next, response.Status))

			return nil, err
		}

		clnt.logger.Info("getPlanets response status OK")

		dto := new(planetsDto)

		if err = json.NewDecoder(response.Body).Decode(dto); err != nil {
			clnt.logger.Error(fmt.Sprintf("error decode planets (%s): %v", *next, err))

			return nil, err
		}

		clnt.logger.Info("getPlanets decode done")

		for index, item := range dto.Results {
			var plnt *planet.Planet
			plnt, err = clnt.toPlanet(&item)

			if err != nil {
				clnt.logger.Error(fmt.Sprintf("error toPlanet (%s)[%d]: %v", *next, index, err))

				return nil, err
			}

			clnt.logger.Info("getPlanets toPlanet done")

			planets = append(planets, plnt)
		}

		if err = response.Body.Close(); err != nil {
			clnt.logger.Error(fmt.Sprintf("error close planets (%s): %v", *next, err))

			return nil, err
		}

		next = dto.Next
	}

	return planets, nil
}
