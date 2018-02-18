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

	assertEqual(t, *transactions[0].TransactionId, "transaction1")
	assertEqual(t, *transactions[0].CustomerId, "customerId")
	assertEqual(t, *transactions[0].AccountNumber, "account1")
	assertEqual(t, *transactions[0].Amount, -5.06)
	assertEqual(t, *transactions[0].Text, "USD 0.64 Amazon web services Kurs: 7.9063")

	assertEqual(t, *transactions[1].TransactionId, "transaction2")
	assertEqual(t, *transactions[1].CustomerId, "customerId")
	assertEqual(t, *transactions[1].AccountNumber, "account1")
	assertEqual(t, *transactions[1].Amount, -100.0)
	assertEqual(t, *transactions[1].Text, "SomeTest")

	assertEqual(t, *transactions[3].TransactionId, "transaction4")
	assertEqual(t, *transactions[3].CustomerId, "customerId")
	assertEqual(t, *transactions[3].AccountNumber, "account1")
	assertEqual(t, *transactions[3].Amount, 100.0)
	assertEqual(t, *transactions[3].Text, "Til Bruk")
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
