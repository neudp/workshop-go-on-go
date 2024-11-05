package swapi

import (
	"fmt"
	"goOnGo/internal/swapi-func/model/logging"
	"net/http"
)

type DoGetRequest func(url string) (*http.Response, error)
type DoRequest func(request *http.Request) (*http.Response, error)

func NewDoGetRequest(doRequest DoRequest, logger *logging.Logger) DoGetRequest {
	return func(url string) (*http.Response, error) {
		request, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			logger.Error(fmt.Sprintf("error creating request for %s: %v", url, err))

			return nil, err
		}

		return doRequest(request)
	}
}
