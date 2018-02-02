package sbankenSDK

import (
	"encoding/json"
	"errors"
	"github.com/larwef/sbankenSDK/authentication"
	"github.com/larwef/sbankenSDK/client"
	"github.com/larwef/sbankenSDK/common"
	"strconv"
	"time"
)

type TransactionsRepository struct {
	common.Repository
}

type transactionsResponse struct {
	AvailableItems int           `json:"availableItems"`
	Items          []Transaction `json:"items"`
	common.Error
}

type TransactionRequest struct {
	AccountNumber string
	StartIndex    int
	Lenght        int
	StartDate     time.Time
	EndDate       time.Time
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

// Constructor for TransactionRepository
func NewTransactionRepository(config Config) *TransactionsRepository {
	token := authentication.NewSbankenToken(config.IdentityServer, config.ClientId, config.ClientSecret)
	return &TransactionsRepository{common.Repository{Url: config.TransactionsEndpoint, Client: client.NewSbankenClient(&token)}}
}

// Gets transactions for a specified account
func (tr *TransactionsRepository) GetTransactions(customerId string, request TransactionRequest) ([]Transaction, error) {
	queryParams := map[string]string{
		"index":     strconv.Itoa(request.StartIndex),
		"length":    strconv.Itoa(request.Lenght),
		"startDate": request.StartDate.Format(time.RFC3339),
		"endDate":   request.EndDate.Format(time.RFC3339),
	}

	response, err := tr.Client.Get(tr.Url+customerId+"/"+request.AccountNumber, queryParams)
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
