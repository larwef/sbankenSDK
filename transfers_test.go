package sbankensdk

import (
	"fmt"
	"net/http"
	"testing"
)

func TestTransferService_Transfer(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "customerId", "customerId")
		fmt.Fprint(w, getTestFileAsString(t, "testdata/transactions_response.json"))
	})

	fromAccountID := "accountId1"
	toAccountID := "accountId2"
	amount := 100.0
	message := "testMessage"

	request := TransferRequest{
		FromAccountID: &fromAccountID,
		ToAccountID:   &toAccountID,
		Amount:        &amount,
		Message:       &message,
	}

	err := client.Transfers.Transfer(request)
	assertNotError(t, err)
}

func TestTransferService_Transfer_WithError(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "customerId", "customerId")
		fmt.Fprint(w, getTestFileAsString(t, "testdata/error_response.json"))
	})

	request := TransferRequest{}
	err := client.Transfers.Transfer(request)
	assertEqual(t, err.Error(), "SomeErrorMessage")
}
