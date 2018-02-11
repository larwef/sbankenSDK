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

	request := TransferRequest{
		FromAccount: "account1",
		ToAccount:   "account2",
		Amount:      100,
		Message:     "testMessage",
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
