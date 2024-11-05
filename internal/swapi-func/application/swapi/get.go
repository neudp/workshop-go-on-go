package swapi

import (
	"fmt"
	"goOnGo/internal/swapi-func/model/logging"
	"net/http"
)

type doGet func(url string) (*http.Response, error)
type DoRequest func(request *http.Request) (*http.Response, error)

func newGet(logError logging.Log, doRequest DoRequest) doGet {
	return func(url string) (*http.Response, error) {
		request, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			// f(g(x)) - композиция функций, обычная практика в функциональном программировании
			logError(fmt.Sprintf("error creating request for %s: %v", url, err))

			return nil, err
		}

		return doRequest(request)
	}
}
