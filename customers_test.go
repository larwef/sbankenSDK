package sbankenSDK

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestCustomersService_GetCustomer(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "customerId", "customerId")
		fmt.Fprint(w, getTestFileAsString(t, "testdata/customer_response.json"))
	})

	customer, err := client.Customers.GetCustomer()
	assertNotError(t, err)

	assertEqual(t, *customer.CustomerId, "customerId")
	expectedDate, _ := time.Parse(time.RFC3339, "1992-02-13T00:00:00Z")
	assertEqual(t, *customer.DateOfBirth.Time, expectedDate)
	assertEqual(t, *customer.EmailAddress, "donduc@gmail.com")
	assertEqual(t, *customer.FirstName, "DONALD")
	assertEqual(t, *customer.LastName, "DUCK")
	assertEqual(t, *customer.PhoneNumbers[0].Number, "1213141516")
	assertEqual(t, *customer.PostalAddress.AddressLine1, "Andebyveien 101")
	assertEqual(t, *customer.StreetAddress.City, "Andeby")
}

func TestCustomersService_GetCustomer_WithError(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, getTestFileAsString(t, "testdata/error_response.json"))
	})

	_, err := client.Customers.GetCustomer()
	assertEqual(t, err.Error(), "SomeErrorMessage")
}
