package swapi

import (
	"fmt"
	"goOnGo/internal/swapi/model/logging"
	"net/http"
)

type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	doer   Doer
	logger logging.Logger
}

func NewClient(doer Doer) *Client {
	return &Client{doer: doer}
}

func (clnt *Client) get(url string) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		clnt.logger.Error(fmt.Sprintf("error creating request for %s: %v", url, err))

		return nil, err
	}

	return clnt.doer.Do(request)
}
