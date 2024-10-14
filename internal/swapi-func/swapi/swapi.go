package swapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"goOnGo/internal/swapi-func/model/domain/character"
	"goOnGo/internal/swapi-func/model/domain/color"
	"goOnGo/internal/swapi-func/model/domain/planet"
	"net/http"
	"strings"
)

type Context interface {
	DoRequest(request *http.Request) (*http.Response, error)
	LogInfo(message string)
	LogError(message string)
}

func GetPlanets(ctx Context) ([]*planet.Planet, error) {
	return getPlanets(ctx, "https://swapi.dev/api/planets/")
}

func GetPlanet(ctx Context, id int) (*planet.Planet, error) {
	return getPlanet(ctx, fmt.Sprintf("https://swapi.dev/api/planets/%d/", id))
}

func GetCharacter(ctx Context, id int) (*character.Character, error) {
	return getCharacter(ctx, fmt.Sprintf("https://swapi.dev/api/people/%d/", id))
}

func GetPeople(ctx Context) ([]*character.Character, error) {
	return getPeople(ctx, "https://swapi.dev/api/people/")
}

func get(ctx Context, url string) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		// f(g(x)) - композиция функций, обычная практика в функциональном программировании
		ctx.LogError(fmt.Sprintf("error creating request for %s: %v", url, err))

		return nil, err
	}

	return ctx.DoRequest(request)
}

func getPlanet(ctx Context, url string) (plnt *planet.Planet, err error) {
	response, err := get(ctx, url)
	if err != nil {
		ctx.LogError(fmt.Sprintf("error getting planet (%s): %v", url, err))

		return nil, err
	}

	ctx.LogInfo("getPlanet response done")

	defer func() {
		err = errors.Join(err, response.Body.Close())
	}()

	if response.StatusCode != http.StatusOK {
		ctx.LogError(fmt.Sprintf("error getting planet (%s): %s", url, response.Status))

		return nil, err
	}

	ctx.LogInfo("getPlanet response status OK")

	dto := new(planetDto)
	if err = json.NewDecoder(response.Body).Decode(dto); err != nil {
		ctx.LogError(fmt.Sprintf("error decode planet (%s): %v", url, err))

		return nil, err
	}

	ctx.LogInfo("getPlanet decode done")

	plnt, err = toPlanet(ctx, dto)

	if err != nil {
		ctx.LogError(fmt.Sprintf("error toPlanet (%s): %v", url, err))

		return nil, err
	}

	ctx.LogInfo("getPlanet toPlanet done")

	return plnt, nil
}

func toPlanet(ctx Context, dto *planetDto) (*planet.Planet, error) {
	rotationPeriod, err := parseInt(dto.RotationPeriod)

	if err != nil {
		ctx.LogError(fmt.Sprintf("error rotationPeriod: %v", err))

		return nil, err
	}

	orbitalPeriod, err := parseInt(dto.OrbitalPeriod)

	if err != nil {
		ctx.LogError(fmt.Sprintf("error orbitalPeriod: %v", err))

		return nil, err
	}

	diameter, err := parseInt(dto.Diameter)

	if err != nil {
		ctx.LogError(fmt.Sprintf("error diameter: %v", err))

		return nil, err
	}

	terrains := strings.Split(dto.Terrain, ",")
	bioms := make([]planet.Biom, len(terrains))

	for i, terrain := range terrains {
		bioms[i] = planet.Biom(strings.TrimSpace(terrain))
	}

	surfaceWater, err := parseFloat(dto.SurfaceWater)

	if err != nil {
		ctx.LogError(fmt.Sprintf("error surfaceWater: %v", err))

		return nil, err
	}

	population, err := parseInt(dto.Population)

	if err != nil {
		ctx.LogError(fmt.Sprintf("error population: %v", err))

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

func getPlanets(ctx Context, url string) ([]*planet.Planet, error) {
	planets := make([]*planet.Planet, 0)

	var next = &url
	for next != nil {
		response, err := get(ctx, *next)
		if err != nil {
			ctx.LogError(fmt.Sprintf("error getting planets (%s): %v", *next, err))

			return nil, err
		}

		ctx.LogInfo("getPlanets response done")

		if response.StatusCode != http.StatusOK {
			ctx.LogError(fmt.Sprintf("error getting planets (%s): %s", *next, response.Status))

			return nil, err
		}

		ctx.LogInfo("getPlanets response status OK")

		dto := new(planetsDto)

		if err = json.NewDecoder(response.Body).Decode(dto); err != nil {
			ctx.LogError(fmt.Sprintf("error decode planets (%s): %v", *next, err))

			return nil, err
		}

		ctx.LogInfo("getPlanets decode done")

		for index, item := range dto.Results {
			var plnt *planet.Planet
			plnt, err = toPlanet(ctx, &item)

			if err != nil {
				ctx.LogError(fmt.Sprintf("error toPlanet (%s)[%d]: %v", *next, index, err))

				return nil, err
			}

			ctx.LogInfo("getPlanets toPlanet done")

			planets = append(planets, plnt)
		}

		if err = response.Body.Close(); err != nil {
			ctx.LogError(fmt.Sprintf("error close planets (%s): %v", *next, err))

			return nil, err
		}

		next = dto.Next
	}

	return planets, nil
}

func getCharacter(ctx Context, url string) (chrctr *character.Character, err error) {
	response, err := get(ctx, url)
	if err != nil {
		ctx.LogError(fmt.Sprintf("error getting character (%s): %v", url, err))

		return nil, err
	}

	ctx.LogInfo("getCharacter response done")

	defer func() {
		err = errors.Join(err, response.Body.Close())
	}()

	if response.StatusCode != http.StatusOK {
		ctx.LogError(fmt.Sprintf("error getting character (%s): %s", url, response.Status))

		return nil, err
	}

	ctx.LogInfo("getCharacter response status OK")

	dto := new(characterDto)

	if err = json.NewDecoder(response.Body).Decode(dto); err != nil {
		ctx.LogError(fmt.Sprintf("error decode character (%s): %v", url, err))

		return nil, err
	}

	ctx.LogInfo("getCharacter decode done")

	chrctr, err = toCharacter(ctx, dto)

	if err != nil {
		ctx.LogError(fmt.Sprintf("error toCharacter (%s): %v", url, err))

		return nil, err
	}

	ctx.LogInfo("getCharacter toCharacter done")

	return chrctr, nil
}

func toCharacter(ctx Context, dto *characterDto) (*character.Character, error) {
	height, err := parseInt(dto.Height)
	if err != nil {
		ctx.LogError(fmt.Sprintf("error height: %v", err))

		return nil, err
	}

	mass, err := parseFloat(dto.Mass)
	if err != nil {
		ctx.LogError(fmt.Sprintf("error mass: %v", err))

		return nil, err
	}

	homeworld, err := getPlanet(ctx, dto.Homeworld)

	if err != nil {
		ctx.LogError(fmt.Sprintf("error getPlanet: %v", err))

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

func getPeople(ctx Context, url string) ([]*character.Character, error) {
	people := make([]*character.Character, 0)

	var next = &url
	for next != nil {
		response, err := get(ctx, *next)
		if err != nil {
			ctx.LogError(fmt.Sprintf("error getting people (%s): %v", *next, err))

			return nil, err
		}

		ctx.LogInfo("getPeople response done")

		if response.StatusCode != http.StatusOK {
			ctx.LogError(fmt.Sprintf("error getting people (%s): %s", *next, response.Status))

			return nil, err
		}

		ctx.LogInfo("getPeople response status OK")

		dto := new(peopleDto)

		if err = json.NewDecoder(response.Body).Decode(dto); err != nil {
			ctx.LogError(fmt.Sprintf("error decode people (%s): %v", *next, err))

			return nil, err
		}

		ctx.LogInfo("getPeople decode done")

		for index, item := range dto.Results {
			var crctr *character.Character
			crctr, err = toCharacter(ctx, &item)

			if err != nil {
				ctx.LogError(fmt.Sprintf("error toCharacter (%s)[%d]: %v", *next, index, err))

				return nil, err
			}

			ctx.LogInfo("getPeople toCharacter done")

			people = append(people, crctr)
		}

		if err = response.Body.Close(); err != nil {
			ctx.LogError(fmt.Sprintf("error close people (%s): %v", *next, err))

			return nil, err
		}

		next = dto.Next
	}

	return people, nil
}
