package transport

import (
	"fmt"
	"goOnGo/internal/swapi-func/application/swapi"
	"goOnGo/internal/swapi-func/model/logging"
	"net/http"
	"net/url"
	"time"
)

var client = &http.Client{
	Timeout: 30 * time.Second,
}

func NewDoSwapiRequest(logLevel logging.LogLevel, swapiUrl string) swapi.DoRequest {
	logInfo := logging.NewLog(logLevel, logging.Info)
	logError := logging.NewLog(logLevel, logging.Error)

	return func(request *http.Request) (*http.Response, error) {
		if request.URL.Host == "" {
			logInfo(fmt.Sprintf("INFO: URL is relative, adding base URL: %s", swapiUrl))

			var err error
			request.URL, err = url.Parse(swapiUrl + request.URL.String())

			if err != nil {
				logError(fmt.Sprintf("ERROR: failed to parse URL: %v", err))

				return nil, err
			}
		}

		logInfo(fmt.Sprintf("INFO: sending request to %s", request.URL.String()))

		return client.Do(request)
	}
}
