package sbanken

import (
	"encoding/json"
	"github.com/larwef/sbanken/common"
	"github.com/larwef/sbanken/client"
	"github.com/larwef/sbanken/authentication"
)

type transactionsRepository struct {
	url    string
	client *client.SbankenClient
}

type transactionsResponse struct {
	AvailableItems int           `json:"availableItems"`
	Items          []Transaction `json:"items"`
	common.Error
}

type Transaction struct {
	TransactionId      string  `json:"transactionId"`
	CustomerId         string  `json:"customerId"`
	AccountNumber      string  `json:"accountNumber"`
	OtherAccountNumber string  `json:"otherAccountNumber"`
	Amount             float64 `json:"amount"`
	Text               string  `json:"text"`
	TransactionType    string  `json:"transactionType"`
	RegistrationDate   string  `json:"registrationDate"`
	AccountingDate     string  `json:"accountingDate"`
	InterestDate       string  `json:"interestDate"`
}

func NewTransactionRepository(config Config) (*transactionsRepository) {
	token := authentication.NewSbankenToken(config.IdentityServer, config.ClientId, config.ClientSecret)
	return &transactionsRepository{url: config.TransactionsEndpoint, client: client.NewSbankenClient(&token)}
}

func (tr transactionsRepository) GetTransactions(customerId string, accountNumber string) ([]Transaction, error) {
	var transactionsRsp transactionsResponse
	response, err := tr.client.Get(tr.url + customerId + "/" + accountNumber)
	if err != nil {
		return transactionsRsp.Items, err
	}

	json.Unmarshal(response, &transactionsRsp)

	return transactionsRsp.Items, err
}
