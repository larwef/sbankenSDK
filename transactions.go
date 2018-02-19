package sbankenSDK

import (
	"errors"
	"strconv"
	"time"
)

type TransactionService service

type transactionsResponse struct {
	AvailableItems *int          `json:"availableItems,omitempty"`
	Items          []Transaction `json:"items,omitempty"`
	sbankenError
}

type TransactionRequest struct {
	AccountNumber string
	StartIndex    int
	Lenght        int
	StartDate     time.Time
	EndDate       time.Time
}

type Transaction struct {
	TransactionId      *string    `json:"transactionId,omitempty"`
	CustomerId         *string    `json:"customerId,omitempty"`
	AccountNumber      *string    `json:"accountNumber,omitempty"`
	OtherAccountNumber *string    `json:"otherAccountNumber,omitempty"`
	Amount             *float64   `json:"amount,omitempty"`
	Text               *string    `json:"text,omitempty"`
	TransactionType    *string    `json:"transactionType,omitempty"`
	RegistrationDate   *time.Time `json:"registrationDate,omitempty"`
	AccountingDate     *time.Time `json:"accountingDate,omitempty"`
	InterestDate       *time.Time `json:"interestDate,omitempty"`
}

// Gets transactions for a specified account
func (ts *TransactionService) GetTransactions(customerId string, request TransactionRequest) ([]Transaction, error) {
	queryParams := map[string]string{
		"index":     strconv.Itoa(request.StartIndex),
		"length":    strconv.Itoa(request.Lenght),
		"startDate": request.StartDate.Format(time.RFC3339),
		"endDate":   request.EndDate.Format(time.RFC3339),
	}

	var transactionsRsp transactionsResponse
	response, err := ts.client.get(ts.client.config.TransactionsEndpoint+customerId+"/"+request.AccountNumber, queryParams, &transactionsRsp)
	defer response.Body.Close()
	if err != nil {
		return []Transaction{}, err
	}

	if *transactionsRsp.IsError == true {
		return []Transaction{}, errors.New(*transactionsRsp.ErrorMessage)
	}

	return transactionsRsp.Items, err
}
