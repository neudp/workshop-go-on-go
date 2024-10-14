package swapi

import (
	"strconv"
	"strings"
)

type characterDto struct {
	Name      string `json:"name"`
	Height    string `json:"height"`
	Mass      string `json:"mass"`
	HairColor string `json:"hair_color"`
	SkinColor string `json:"skin_color"`
	EyeColor  string `json:"eye_color"`
	BirthYear string `json:"birth_year"`
	Gender    string `json:"gender"`
	Homeworld string `json:"homeworld"`
}

type peopleDto struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []characterDto `json:"results"`
}

type planetDto struct {
	Name           string `json:"name"`
	RotationPeriod string `json:"rotation_period"`
	OrbitalPeriod  string `json:"orbital_period"`
	Diameter       string `json:"diameter"`
	Climate        string `json:"climate"`
	Gravity        string `json:"gravity"`
	Terrain        string `json:"terrain"`
	SurfaceWater   string `json:"surface_water"`
	Population     string `json:"population"`
}

type planetsDto struct {
	Count    int         `json:"count"`
	Next     *string     `json:"next"`
	Previous *string     `json:"previous"`
	Results  []planetDto `json:"results"`
}

func parseInt(value string) (*int, error) {
	if value == "unknown" {
		return nil, nil
	}

	value = strings.Replace(value, ",", "", -1)

	parsed, err := strconv.Atoi(value)

	return &parsed, err
}

func parseFloat(value string) (*float64, error) {
	if value == "unknown" {
		return nil, nil
	}

	value = strings.Replace(value, ",", "", -1)

	parsed, err := strconv.ParseFloat(value, 64)

	return &parsed, err
}
