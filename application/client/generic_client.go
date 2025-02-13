package client

import (
	"net/http"
)

type GenericClient struct{}

func NewGenericClient() *GenericClient {
	return &GenericClient{}
}

type GenericClientInterface interface {
	CallClient() *http.Response
}

func (c GenericClient) CallClient(url string) (*http.Response, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return http.DefaultClient.Do(req)
}
