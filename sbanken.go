package sbankenSDK

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Client struct {
	client *http.Client
	config Config

	common service

	Accounts     *AccountService
	Transactions *TransactionService
	Transfers    *TransferService
}

func NewClient(httpClient *http.Client, config Config) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{client: httpClient, config: config}

	c.common.client = c
	c.Accounts = (*AccountService)(&c.common)
	c.Transactions = (*TransactionService)(&c.common)
	c.Transfers = (*TransferService)(&c.common)

	return c
}

type service struct {
	client *Client
}

func (c *Client) get(url string, queryParams map[string]string, responseObj interface{}) (*http.Response, error) {
	request, err := c.getRequest(url, http.MethodGet, queryParams, nil, responseObj)
	if err != nil {
		return nil, err
	}

	return c.do(request, responseObj)
}

func (c *Client) post(url string, queryParams map[string]string, requestObj interface{}, responseObj interface{}) (*http.Response, error) {
	request, err := c.getRequest(url, http.MethodPost, queryParams, requestObj, responseObj)
	request.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	return c.do(request, responseObj)
}

func (c *Client) getRequest(url string, method string, queryParams map[string]string, requestObj interface{}, responseObj interface{}) (*http.Request, error) {
	payload, err := json.Marshal(requestObj)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return request, err
	}

	request.Header.Add("Accept", "application/json")

	query := request.URL.Query()
	for key, value := range queryParams {
		query.Add(key, value)
	}
	request.URL.RawQuery = query.Encode()

	return request, err
}

func (c *Client) do(req *http.Request, responseObj interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if responseObj != nil {
		if w, ok := responseObj.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(responseObj)
			if err == io.EOF {
				err = nil // ignore EOF errors caused by empty response body
			}
		}
	}

	return resp, err
}
