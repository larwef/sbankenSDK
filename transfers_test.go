package sbankenSDK

import (
	"fmt"
	"net/http"
	"testing"
	"io/ioutil"
	"strings"
)

func TestTransferService_Transfer(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/customerId", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testHeader(t, r, "Accept", "application/json")
		body, _ := ioutil.ReadAll(r.Body)
		payload := string(body)
		strings.Contains(payload, "\"fromAccount\":\"account1\"")
		strings.Contains(payload, "\"toAccount\":\"account2\"")
		strings.Contains(payload, "\"amount\":100")
		strings.Contains(payload, "\"message\":\"testMessage\"")
		fmt.Fprint(w, getTestFileAsString(t, "testdata/transactions_response.json"))
	})

	request := NewTransferRequest().
		WithFromAccount("account1").
		WithToAccount("account2").
		WithAmount(100).
		WithMessage("testMessage")

	err := client.Transfers.Transfer("customerId", *request)
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
