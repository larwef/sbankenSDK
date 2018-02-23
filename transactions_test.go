package sbankenSDK

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestTransactionService_GetTransactions(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	startDate := time.Now().Add(-24 * time.Hour)
	endDate := time.Now()

	mux.HandleFunc("/customerId/account1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testHeader(t, r, "Accept", "application/json")
		queryValues := r.URL.Query()
		assertEqual(t, queryValues["index"][0], "0")
		assertEqual(t, queryValues["length"][0], "100")
		assertEqual(t, queryValues["startDate"][0], startDate.Format(time.RFC3339))
		assertEqual(t, queryValues["endDate"][0], endDate.Format(time.RFC3339))
		fmt.Fprint(w, getTestFileAsString(t, "testdata/transactions_response.json"))
	})
	request := TransactionRequest{
		AccountNumber: "account1",
		StartIndex:    0,
		Lenght:        100,
		StartDate:     startDate,
		EndDate:       endDate,
	}

	transactions, err := client.Transactions.GetTransactions("customerId", request)
	assertNotError(t, err)

	assertEqual(t, len(transactions), 5)

	var val interface{}
	var isSet, isNull bool
	val, isSet, isNull = transactions[0].GetTransactionId()
	assertEqual(t, val, "transaction1")
	assertIsSetAndNotNull(t, isSet, isNull)
	val, isSet, isNull = transactions[0].GetAccountNumber()
	assertEqual(t, val, "account1")
	assertIsSetAndNotNull(t, isSet, isNull)
	val, isSet, isNull = transactions[0].GetAmount()
	assertEqual(t, val, -5.06)
	assertIsSetAndNotNull(t, isSet, isNull)
	val, isSet, isNull = transactions[0].GetText()
	assertEqual(t, val, "USD 0.64 Amazon web services Kurs: 7.9063")
	assertIsSetAndNotNull(t, isSet, isNull)

	val, isSet, isNull = transactions[1].GetTransactionId()
	assertEqual(t, val, "transaction2")
	assertIsSetAndNotNull(t, isSet, isNull)
	val, isSet, isNull = transactions[1].GetAccountNumber()
	assertEqual(t, val, "account1")
	assertIsSetAndNotNull(t, isSet, isNull)
	val, isSet, isNull = transactions[1].GetAmount()
	assertEqual(t, val, -100.0)
	assertIsSetAndNotNull(t, isSet, isNull)
	val, isSet, isNull = transactions[1].GetText()
	assertEqual(t, val, "SomeTest")
	assertIsSetAndNotNull(t, isSet, isNull)

	val, isSet, isNull = transactions[3].GetTransactionId()
	assertEqual(t, val, "transaction4")
	assertIsSetAndNotNull(t, isSet, isNull)
	val, isSet, isNull = transactions[3].GetAccountNumber()
	assertEqual(t, val, "account1")
	assertIsSetAndNotNull(t, isSet, isNull)
	val, isSet, isNull = transactions[3].GetAmount()
	assertEqual(t, val, 100.0)
	assertIsSetAndNotNull(t, isSet, isNull)
	val, isSet, isNull = transactions[3].GetText()
	assertEqual(t, val, "Til Bruk")
	assertIsSetAndNotNull(t, isSet, isNull)
}

func TestTransactionService_GetTransactions_WithError(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/customerId/account1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, getTestFileAsString(t, "testdata/error_response.json"))
	})

	request := TransactionRequest{
		AccountNumber: "account1",
		StartIndex:    0,
		Lenght:        100,
		StartDate:     time.Now().Add(-24 * time.Hour),
		EndDate:       time.Now(),
	}

	transactions, err := client.Transactions.GetTransactions("customerId", request)
	assertEqual(t, len(transactions), 0)
	assertEqual(t, err.Error(), "SomeErrorMessage")
}
