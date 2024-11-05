package transport

import (
	"fmt"
	"goOnGo/internal/swapi/model/logging"
	"net/http"
	"net/url"
	"time"
)

type HTTPClient struct {
	baseUrl string
	client  *http.Client
	logger  logging.Logger
}

func NewHttpClient(baseUrl string, logger logging.Logger) *HTTPClient {
	return &HTTPClient{
		baseUrl: baseUrl,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		logger: logger,
	}
}

func (client *HTTPClient) Do(request *http.Request) (*http.Response, error) {
	// Если хост не указан, то URL считается относительным
	if request.URL.Host == "" {
		client.logger.Info(fmt.Sprintf("INFO: URL is relative, adding base URL: %s", client.baseUrl))

		var err error
		request.URL, err = url.Parse(client.baseUrl + request.URL.String())

		if err != nil {
			client.logger.Error(fmt.Sprintf("ERROR: failed to parse URL: %v", err))

			return nil, err
		}
	}

	client.logger.Info(fmt.Sprintf("INFO: sending request to %s", request.URL.String()))

	return client.client.Do(request)
}
