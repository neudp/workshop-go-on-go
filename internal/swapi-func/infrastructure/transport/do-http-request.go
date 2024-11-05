package transport

import (
	"fmt"
	"goOnGo/internal/swapi-func/model/logging"
	"net/http"
	"net/url"
	"time"
)

type DoRequest func(*http.Request) (*http.Response, error)

var client = &http.Client{
	Timeout: 30 * time.Second,
}

func NewDoRequest(baseUrl string, logLevel logging.LogLevel) DoRequest {
	logInfo := logging.NewLog(logLevel, logging.Info)
	logError := logging.NewLog(logLevel, logging.Error)

	return func(request *http.Request) (*http.Response, error) {
		if request.URL.Host == "" {
			logInfo(fmt.Sprintf("URL is relative, adding base URL: %s", baseUrl))

			var err error
			request.URL, err = url.Parse(baseUrl + request.URL.String())

			if err != nil {
				logError(fmt.Sprintf("failed to parse URL: %v", err))

				return nil, err
			}
		}

		logInfo(fmt.Sprintf("sending request to %s", request.URL.String()))

		return client.Do(request)
	}
}
