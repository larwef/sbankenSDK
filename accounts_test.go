package sbankenSDK

import (
	"fmt"
	"net/http"
	"testing"
)

func TestAccountService_GetAccounts(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "customerId", "customerId")
		fmt.Fprint(w, getTestFileAsString(t, "testdata/accounts_response.json"))
	})

	accounts, err := client.Accounts.GetAccounts()
	assertNotError(t, err)

	assertEqual(t, len(accounts), 4)

	assertEqual(t, *accounts[0].OwnerCustomerId, "customerId")
	assertEqual(t, *accounts[0].AccountNumber, "account1")
	assertEqual(t, *accounts[0].AccountId, "accountId1")

	assertEqual(t, *accounts[1].OwnerCustomerId, "customerId")
	assertEqual(t, *accounts[1].AccountNumber, "account2")
	assertEqual(t, *accounts[1].AccountId, "accountId2")

	assertEqual(t, *accounts[2].OwnerCustomerId, "customerId")
	assertEqual(t, *accounts[2].AccountNumber, "account3")
	assertEqual(t, *accounts[2].AccountId, "accountId3")

	assertEqual(t, *accounts[3].OwnerCustomerId, "customerId")
	assertEqual(t, *accounts[3].AccountNumber, "account4")
	assertEqual(t, *accounts[3].AccountId, "accountId4")
}

func TestAccountService_GetAccounts_WithError(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "customerId", "customerId")
		fmt.Fprint(w, getTestFileAsString(t, "testdata/error_response.json"))
	})

	accounts, err := client.Accounts.GetAccounts()
	assertEqual(t, err.Error(), "SomeErrorMessage")
	assertEqual(t, len(accounts), 0)
}

func TestAccountService_GetAccount(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/accountId1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "customerId", "customerId")
		fmt.Fprint(w, getTestFileAsString(t, "testdata/account_response.json"))
	})

	account, err := client.Accounts.GetAccount("accountId1")
	assertNotError(t, err)
	assertEqual(t, *account.OwnerCustomerId, "customerId")
	assertEqual(t, *account.AccountNumber, "account1")
	assertEqual(t, *account.Available, 10389.51)
}

func TestAccountService_GetAccount_WithError(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/accountId1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, getTestFileAsString(t, "testdata/error_response.json"))
	})

	_, err := client.Accounts.GetAccount("accountId1")
	assertEqual(t, err.Error(), "SomeErrorMessage")
}
