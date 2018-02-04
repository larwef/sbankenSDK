package sbankenSDK

import (
	"net/http"
	"github.com/larwef/sbankenSDK/authentication"
	"io"
)

// TODO: Use a client with authentication instead of token manually
type Client struct {
	client *http.Client
	config Config
	token  authentication.SbankenToken

	common service

	Accounts     *AccountService
	Transactions *TransactionService
	Transfers    *TransferService
}

func NewClient(httpClient *http.Client, config Config) (*Client) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{client: httpClient, config: config}

	c.common.client = c
	c.token = authentication.NewSbankenToken(config.IdentityServer, config.ClientId, config.ClientSecret)
	c.Accounts = (*AccountService)(&c.common)
	c.Transactions = (*TransactionService)(&c.common)
	c.Transfers = (*TransferService)(&c.common)

	return c
}

type service struct {
	client *Client
}

func (c *Client) Get(url string, queryParams map[string]string) (*http.Response, error) {
	request, err := c.getRequest(url, http.MethodGet, queryParams, nil)
	if err != nil {
		return nil, err
	}

	return c.client.Do(request)
}

func (c *Client) Post(url string, queryParams map[string]string, payload io.Reader) (*http.Response, error) {
	request, err := c.getRequest(url, http.MethodPost, queryParams, payload)
	request.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	return c.client.Do(request)
}

func (c *Client) getRequest(url string, method string, queryParams map[string]string, payload io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(method, url, payload)
	if err != nil {
		return request, err
	}
	// TODO: refresh token when close to expiration or expired or see todo in top
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Authorization", c.token.GetTokenType()+" "+c.token.GetTokenString())

	query := request.URL.Query()
	for key, value := range queryParams {
		query.Add(key, value)
	}
	request.URL.RawQuery = query.Encode()

	return request, err
}
