package client

import (
	"net/http"
	"io/ioutil"
	"github.com/larwef/sbankenSDK/authentication"
	"bytes"
)

type SbankenClient struct {
	token  authentication.Token
	client http.Client
}

func NewSbankenClient(token authentication.Token) (*SbankenClient) {
	return &SbankenClient{token: token, client: http.Client{}}
}

func (sbt SbankenClient) Get(url string, queryParams map[string]string) ([]byte, error) {
	request, err := sbt.getRequest(url, http.MethodGet, queryParams, nil)
	if err != nil {
		return []byte{}, err
	}

	return sbt.doRequest(request)
}

func (sbt SbankenClient) Post(url string, queryParams map[string]string, payload []byte) ([]byte, error) {
	request, err := sbt.getRequest(url, http.MethodPost, queryParams, payload)
	request.Header.Add("Content-Type", "application/json")
	if err != nil {
		return []byte{}, err
	}

	return sbt.doRequest(request)
}

func (sbt SbankenClient) doRequest(request *http.Request) ([]byte, error) {
	response, err := sbt.client.Do(request)
	if err != nil {
		return []byte{}, err
	}

	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

func (sbt SbankenClient) getRequest(url string, method string, queryParams map[string]string, payload []byte) (*http.Request, error){
	request, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return request, err
	}
	// TODO: refresh token when close to expiration or expired
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Authorization", sbt.token.GetTokenType()+" "+sbt.token.GetTokenString())

	query := request.URL.Query()
	for key, value := range queryParams {
		query.Add(key, value)
	}
	request.URL.RawQuery = query.Encode()

	return request, err
}
