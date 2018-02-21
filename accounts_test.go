package sbankenSDK

import (
	"fmt"
	"net/http"
	"testing"
)

func TestAccountService_GetAccounts(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/customerId", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, getTestFileAsString(t, "testdata/accounts_response.json"))
	})

	accounts, err := client.Accounts.GetAccounts("customerId")
	assertNotError(t, err)

	assertEqual(t, len(accounts), 4)

	var val interface{}
	var isSet, isNull bool
	val, isSet, isNull = accounts[0].GetCustomerId()
	assertEqual(t, val, "customerId")
	assertIsSetAndNotNull(t, isSet, isNull)
	val, isSet, isNull = accounts[0].GetAccountNumber()
	assertEqual(t, val, "account1")
	assertIsSetAndNotNull(t, isSet, isNull)

	val, isSet, isNull = accounts[1].GetCustomerId()
	assertEqual(t, val, "customerId")
	assertIsSetAndNotNull(t, isSet, isNull)
	val, isSet, isNull = accounts[1].GetAccountNumber()
	assertEqual(t, val, "account2")
	assertIsSetAndNotNull(t, isSet, isNull)

	val, isSet, isNull = accounts[2].GetCustomerId()
	assertEqual(t, val, "customerId")
	assertIsSetAndNotNull(t, isSet, isNull)
	val, isSet, isNull = accounts[2].GetAccountNumber()
	assertEqual(t, val, "account3")
	assertIsSetAndNotNull(t, isSet, isNull)

	val, isSet, isNull = accounts[3].GetCustomerId()
	assertEqual(t, val, "customerId")
	assertIsSetAndNotNull(t, isSet, isNull)
	val, isSet, isNull = accounts[3].GetAccountNumber()
	assertEqual(t, val, "account4")
	assertIsSetAndNotNull(t, isSet, isNull)
}

func TestAccountService_GetAccounts_WithError(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/customerId", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
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
		testMethod(t, r, http.MethodGet)
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, getTestFileAsString(t, "testdata/account_response.json"))
	})

	account, err := client.Accounts.GetAccount("customerId", "account1")
	assertNotError(t, err)
	var val interface{}
	var isSet, isNull bool
	val, isSet, isNull = account.GetCustomerId()
	assertEqual(t, val, "customerId")
	assertIsSetAndNotNull(t, isSet, isNull)
	val, isSet, isNull = account.GetAccountNumber()
	assertEqual(t, val, "account1")
	assertIsSetAndNotNull(t, isSet, isNull)
	val, isSet, isNull = account.GetAvailable()
	assertEqual(t, val, 10389.51)
	assertIsSetAndNotNull(t, isSet, isNull)
}

func TestAccountService_GetAccount_WithError(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/customerId/account1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, getTestFileAsString(t, "testdata/error_response.json"))
	})

	_, err := client.Accounts.GetAccount("customerId", "account1")
	assertEqual(t, err.Error(), "SomeErrorMessage")
}
