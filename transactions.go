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
	AccountId  string
	StartIndex int
	Lenght     int
	StartDate  time.Time
	EndDate    time.Time
}

type Transaction struct {
	AccountingDate              *time.Time   `json:"accountingDate,omitempty"`
	InterestDate                *time.Time   `json:"interestDate,omitempty"`
	OtherAccountNumber          *string      `json:"otherAccountNumber,omitempty"`
	OtherAccountNumberSpecified *bool        `json:"otherAccountNumberSpecified,omitempty"`
	Amount                      *float64     `json:"amount,omitempty"`
	Text                        *string      `json:"text,omitempty"`
	TransactionType             *string      `json:"transactionType,omitempty"`
	TransactionTypeCode         *int64       `json:"transactionTypeCode,omitempty"`
	TransactionTypeText         *string      `json:"transactionTypeText,omitempty"`
	IsReservation               *bool        `json:"isReservation,omitempty"`
	ReservationType             *bool        `json:"reservationType,omitempty"`
	Source                      *int         `json:"source,omitempty"`
	CardDetails                 *cardDetails `json:"cardDetails,omitempty"`
	CardDetailsSpecified        *bool        `json:"cardDetailsSpecified,omitempty"`
}

type cardDetails struct {
	CardNumber                  *string    `json:"cardNumber,omitempty"`
	CurrencyAmount              *float64   `json:"currencyAmount,omitempty"`
	CurrencyRate                *float64   `json:"currencyRate"`
	MerchantCategoryCode        *string    `json:"merchantCategoryCode,omitempty"`
	MerchantCategoryDescription *string    `json:"merchantCategoryDescription,omitempty"`
	MerchantCity                *string    `json:"merchantCity,omitempty"`
	MerchantName                *string    `json:"merchantName,omitempty"`
	OriginalCurrencyCode        *string    `json:"originalCurrencyCode,omitempty"`
	PurchaseDate                *time.Time `json:"purchaseDate,omitempty"`
	TransactionId               *string    `json:"transactionId"`
}

// Gets transactions for a specified account
func (ts *TransactionService) GetTransactions(request TransactionRequest) ([]Transaction, error) {
	queryParams := map[string]string{
		"index":     strconv.Itoa(request.StartIndex),
		"length":    strconv.Itoa(request.Lenght),
		"startDate": request.StartDate.Format(time.RFC3339),
		"endDate":   request.EndDate.Format(time.RFC3339),
	}

	var response transactionsResponse
	rsp, err := ts.client.get(ts.client.config.TransactionsEndpoint+request.AccountId, queryParams, &response)
	defer rsp.Body.Close()
	if err != nil {
		return []Transaction{}, err
	}

	if response.IsError != nil && *response.IsError == true {
		return []Transaction{}, errors.New(*response.ErrorMessage)
	}

	return response.Items, err
}
