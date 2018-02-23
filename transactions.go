package sbankenSDK

import (
	"strconv"
	"time"
)

type TransactionService service

type Transaction entity

type TransactionRequest struct {
	AccountNumber string
	StartIndex    int
	Lenght        int
	StartDate     time.Time
	EndDate       time.Time
}

// Gets transactions for a specified account
func (ts *TransactionService) GetTransactions(customerId string, request TransactionRequest) ([]Transaction, error) {
	queryParams := map[string]string{
		"index":     strconv.Itoa(request.StartIndex),
		"length":    strconv.Itoa(request.Lenght),
		"startDate": request.StartDate.Format(time.RFC3339),
		"endDate":   request.EndDate.Format(time.RFC3339),
	}

	var transactionsRsp response
	_, err := ts.client.get(ts.client.config.TransactionsEndpoint+customerId+"/"+request.AccountNumber, queryParams, &transactionsRsp.properties)
	if err != nil {
		return []Transaction{}, err
	}

	if err := transactionsRsp.getError(); err != nil {
		return []Transaction{}, err
	}

	items, err := getRequiredProperty(ITEMS, transactionsRsp.properties)
	if err != nil {
		return nil, err
	}

	var val []Transaction
	for _, element := range items.([]interface{}) {
		val = append(val, Transaction{properties: element.(map[string]interface{})})
	}

	return val, err
}

func (t *Transaction) GetTransactionId() (val string, isSet bool, isNull bool) {
	return getString(TRANSACTION_ID, t.properties)
}

func (t *Transaction) GetCustomerId() (val string, isSet bool, isNull bool) {
	return getString(CUSTOMER_ID, t.properties)
}

func (t *Transaction) GetAccountNumber() (val string, isSet bool, isNull bool) {
	return getString(ACCOUNT_NUMBER, t.properties)
}

func (t *Transaction) GetOtherAccountNumber() (val string, isSet bool, isNull bool) {
	return getString(OTHER_ACCOUNT_NUMBER, t.properties)
}

func (t *Transaction) GetAmount() (val float64, isSet bool, isNull bool) {
	return getFloat64(AMOUNT, t.properties)
}

func (t *Transaction) GetText() (val string, isSet bool, isNull bool) {
	return getString(TEXT, t.properties)
}

func (t *Transaction) GetTransactionType() (val string, isSet bool, isNull bool) {
	return getString(TRANSACTION_ID, t.properties)
}

func (t *Transaction) GetRegistrationDate() (val time.Time, isSet bool, isNull bool) {
	return getTime(REGISTRATION_DATE, t.properties)
}

func (t *Transaction) GetAccountingDate() (val time.Time, isSet bool, isNull bool) {
	return getTime(ACCOUNTING_DATE, t.properties)
}

func (t *Transaction) GetInterestDate() (val time.Time, isSet bool, isNull bool) {
	return getTime(INTEREST_DATE, t.properties)
}