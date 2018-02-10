package sbankenSDK

import (
	"testing"
	"net/http"
	"fmt"
)

func TestAccountService_GetAccounts(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/customerId", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, getTestFileAsString(t, "testdata/accounts_response.json"))
	})

	accounts, err := client.Accounts.GetAccounts("customerId")
	assertNotError(t, err)

	assertEqual(t, len(accounts), 4)

	assertEqual(t, accounts[0].CustomerId, "customerId")
	assertEqual(t, accounts[0].AccountNumber, "account1")

	assertEqual(t, accounts[1].CustomerId, "customerId")
	assertEqual(t, accounts[1].AccountNumber, "account2")

	assertEqual(t, accounts[2].CustomerId, "customerId")
	assertEqual(t, accounts[2].AccountNumber, "account3")

	assertEqual(t, accounts[3].CustomerId, "customerId")
	assertEqual(t, accounts[3].AccountNumber, "account4")
}

func TestAccountService_GetAccounts_WithError(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/customerId", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, getTestFileAsString(t, "testdata/error_response.json"))
	})

	accounts, err := client.Accounts.GetAccounts("customerId")
	assertEqual(t, err.Error(), "SomeErrorMessage")
	assertEqual(t, len(accounts), 0)
}

func TestAccountService_GetAccount(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/customerId/account1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, getTestFileAsString(t, "testdata/account_response.json"))
	})

	account, err := client.Accounts.GetAccount("customerId", "account1")
	assertNotError(t, err)
	assertEqual(t, account.CustomerId, "customerId")
	assertEqual(t, account.AccountNumber, "account1")
	assertEqual(t, account.Available, 10389.51)
}

func TestAccountService_GetAccount_WithError(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/customerId/account1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, getTestFileAsString(t, "testdata/error_response.json"))
	})

	_, err := client.Accounts.GetAccount("customerId", "account1")
	assertEqual(t, err.Error(), "SomeErrorMessage")
}
