package sbankenSDK

import (
	"fmt"
	"net/http"
	"testing"
)

func TestTransferService_Transfer(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/customerId", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, getTestFileAsString(t, "testdata/transactions_response.json"))
	})

	fromAccount := "accoutn1"
	toAccount := "account2"
	amount := 100.0
	message := "testMessage"

	request := TransferRequest{
		FromAccount: &fromAccount,
		ToAccount:   &toAccount,
		Amount:      &amount,
		Message:     &message,
	}

	err := client.Transfers.Transfer("customerId", request)
	assertNotError(t, err)
}

func TestTransferService_Transfer_WithError(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/customerId", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, getTestFileAsString(t, "testdata/error_response.json"))
	})

	request := TransferRequest{}
	err := client.Transfers.Transfer("customerId", request)
	assertEqual(t, err.Error(), "SomeErrorMessage")
}
