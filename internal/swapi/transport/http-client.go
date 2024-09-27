package transport

import (
	"goOnGo/internal/swapi/config"
	"goOnGo/internal/swapi/model"
	"net/http"
	"net/url"
)

type HTTPClient struct {
	url    string
	client *http.Client
	logger model.Logger
}

func NewSwapiClient(config *config.Config, logger model.Logger) *HTTPClient {
	return &HTTPClient{
		url:    config.SwapiURL,
		client: http.DefaultClient,
		logger: logger,
	}
}

func (client *HTTPClient) Do(request *http.Request) (*http.Response, error) {
	// Если хост не указан, то URL считается относительным
	if request.URL.Host == "" {
		client.logger.Infof("INFO: URL is relative, adding base URL: %s", client.url)

		var err error
		request.URL, err = url.Parse(client.url + request.URL.String())

		if err != nil {
			client.logger.Errorf("ERROR: failed to parse URL: %v", err)

			return nil, err
		}
	}

	client.logger.Infof("INFO: sending request to %s", request.URL.String())

	return client.client.Do(request)
}
