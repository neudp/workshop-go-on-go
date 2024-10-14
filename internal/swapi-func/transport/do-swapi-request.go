package transport

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type DoSwapiRequestContext interface {
	SwapiURL() string
	LogInfo(message string)
	LogError(message string)
}

var client = &http.Client{
	Timeout: 30 * time.Second,
}

func DoSwapiRequest(ctx DoSwapiRequestContext, request *http.Request) (*http.Response, error) {
	// Если хост не указан, то URL считается относительным
	if request.URL.Host == "" {
		ctx.LogInfo(fmt.Sprintf("INFO: URL is relative, adding base URL: %s", ctx.SwapiURL()))
		swapiURL := ctx.SwapiURL()

		var err error
		request.URL, err = url.Parse(swapiURL + request.URL.String())

		if err != nil {
			ctx.LogError(fmt.Sprintf("ERROR: failed to parse URL: %v", err))

			return nil, err
		}
	}

	ctx.LogInfo(fmt.Sprintf("INFO: sending request to %s", request.URL.String()))

	return client.Do(request)
}
