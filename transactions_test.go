package sbankensdk

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

	mux.HandleFunc("/accountId1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "customerId", "customerId")
		queryValues := r.URL.Query()
		assertEqual(t, queryValues["index"][0], "0")
		assertEqual(t, queryValues["length"][0], "100")
		assertEqual(t, queryValues["startDate"][0], startDate.Format(time.RFC3339))
		assertEqual(t, queryValues["endDate"][0], endDate.Format(time.RFC3339))
		fmt.Fprint(w, getTestFileAsString(t, "testdata/transactions_response.json"))
	})
	request := TransactionRequest{
		AccountID:  "accountId1",
		StartIndex: 0,
		Lenght:     100,
		StartDate:  startDate,
		EndDate:    endDate,
	}

	transactions, err := client.Transactions.GetTransactions(request)
	assertNotError(t, err)

	assertEqual(t, len(transactions), 5)

	assertEqual(t, *transactions[0].CardDetailsSpecified, false)
	assertEqual(t, *transactions[0].TransactionType, "Visa")
	assertEqual(t, *transactions[0].Amount, -5.4)
	assertEqual(t, *transactions[0].Text, "Sometext1")

	assertEqual(t, *transactions[1].CardDetailsSpecified, false)
	assertEqual(t, *transactions[1].TransactionType, "KREDITRTE")
	assertEqual(t, *transactions[1].Amount, 0.31)
	assertEqual(t, *transactions[1].Text, "Sometext2")

	assertEqual(t, *transactions[3].CardDetailsSpecified, false)
	assertEqual(t, *transactions[3].TransactionType, "AVTGI")
	assertEqual(t, *transactions[3].Amount, -1740.0)
	assertEqual(t, *transactions[3].Text, "Sometext4")
}

func TestTransactionService_GetTransactions_WithError(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/accountId1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, getTestFileAsString(t, "testdata/error_response.json"))
	})

	request := TransactionRequest{
		AccountID:  "accountId1",
		StartIndex: 0,
		Lenght:     100,
		StartDate:  time.Now().Add(-24 * time.Hour),
		EndDate:    time.Now(),
	}

	transactions, err := client.Transactions.GetTransactions(request)
	assertEqual(t, len(transactions), 0)
	assertEqual(t, err.Error(), "SomeErrorMessage")
}
