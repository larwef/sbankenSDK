package sbankenSDK

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	CustomerId           string `json:"customerId,omitempty"`
	ClientId             string `json:"clientId,omitempty"`
	ClientSecret         string `json:"clientSecret,omitempty"`
	IdentityServer       string `json:"identityServer,omitempty"`
	AccountsEndpoint     string `json:"accountsEndpoint,omitempty"`
	TransactionsEndpoint string `json:"transactionsEndpoint,omitempty"`
	TransfersEndpoint    string `json:"transfersEndpoint,omitempty"`
	CustomersEndpoint    string `json:"customersEndpoint,omitempty"`
}

// Retrieves config from .json file
func ConfigFromFile(filepath string) Config {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	var config Config
	json.Unmarshal(file, &config)

	return config
}
