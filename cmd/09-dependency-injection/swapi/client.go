package swapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Transport interface {
	Do(request *http.Request) (*http.Response, error)
}

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

func (dto *characterDto) toCharacter() (*Character, error) {

}

type peopleDto struct {
	Count    int          `json:"count"`
	Next     *string      `json:"next"`
	Previous *string      `json:"previous"`
	Results  []*Character `json:"results"`
}

type Client struct {
	transport Transport
}

func NewSwapiClient(transport Transport) *Client {
	return &Client{transport: transport}
}

func (client *Client) GetPeople(limit, offset int) (*[]Character, error) {
	characters := make([]Character, 0)

	page := 1

	for {
		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/people/?limit=%d&offset=%d", limit, offset), nil)
		if err != nil {
			return nil, err
		}

		response, err := client.transport.Do(request)

		if err != nil {
			return nil, err
		}

		body, err := io.ReadAll(response.Body)

		if err != nil {
			return nil, err
		}

		var people peopleDto

		err = json.Unmarshal(body, &people)

		if err != nil {
			return nil, err
		}

		characters = append(characters, people.Results...)

		if people.Next == nil {
			break
		}

		page++
	}
}

func (client *Client) GetCharacter(id int) (*Character, error) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/people/%d/", id), nil)
	if err != nil {
		return nil, err
	}

	response, err := client.transport.Do(request)

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var character Character

	err = json.Unmarshal(body, &character)

	if err != nil {
		return nil, err
	}

	return &character, nil
}
