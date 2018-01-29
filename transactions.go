package sbankenSDK

import (
	"github.com/larwef/sbankenSDK/common"
	"github.com/larwef/sbankenSDK/authentication"
	"github.com/larwef/sbankenSDK/client"
	"encoding/json"
	"time"
	"strconv"
	"errors"
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
	TransactionId      string    `json:"transactionId"`
	CustomerId         string    `json:"customerId"`
	AccountNumber      string    `json:"accountNumber"`
	OtherAccountNumber string    `json:"otherAccountNumber"`
	Amount             float64   `json:"amount"`
	Text               string    `json:"text"`
	TransactionType    string    `json:"transactionType"`
	RegistrationDate   time.Time `json:"registrationDate"`
	AccountingDate     time.Time `json:"accountingDate"`
	InterestDate       time.Time `json:"interestDate"`
}

func NewTransactionRepository(config Config) (*transactionsRepository) {
	token := authentication.NewSbankenToken(config.IdentityServer, config.ClientId, config.ClientSecret)
	return &transactionsRepository{url: config.TransactionsEndpoint, client: client.NewSbankenClient(&token)}
}

func (tr transactionsRepository) GetTransactions(customerId string, accountNumber string, index int, length int, startDate time.Time, endDate time.Time) ([]Transaction, error) {
	queryParams := map[string]string{
		"index":     strconv.Itoa(index),
		"length":    strconv.Itoa(length),
		"startDate": startDate.Format(time.RFC3339),
		"endDate":   endDate.Format(time.RFC3339),
	}

	response, err := tr.client.Get(tr.url+customerId+"/"+accountNumber, queryParams)
	defer response.Body.Close()
	if err != nil {
		return []Transaction{}, err
	}

	var transactionsRsp transactionsResponse
	err = json.NewDecoder(response.Body).Decode(&transactionsRsp)
	if transactionsRsp.IsError == true {
		return transactionsRsp.Items, errors.New(transactionsRsp.ErrorMessage)
	}

	return transactionsRsp.Items, err
}
