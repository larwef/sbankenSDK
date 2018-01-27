package client

import (
	"net/http"
	"log"
	"os"
	"io/ioutil"
	"github.com/larwef/sbankenSDK/authentication"
)

type SbankenClient struct {
	token  authentication.Token
	client http.Client
}

func NewSbankenClient(token authentication.Token) (*SbankenClient) {
	return &SbankenClient{token: token, client: http.Client{}}
}

func (sbt SbankenClient) Get(url string) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	// TODO: refresh token when close to expiration or expired
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Authorization", sbt.token.GetTokenType() + " " +sbt.token.GetTokenString())

	response, err := sbt.client.Do(request)
	if err != nil {
		return []byte{}, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	return body, err
}
