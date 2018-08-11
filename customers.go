package sbankensdk

import (
	"errors"
	"strings"
	"time"
)

// CustomersService handles communication with the Customers part of the API.
type CustomersService service

type customerResponse struct {
	Item *Customer `json:"item,omitempty"`
	sbankenError
}

// Customer represents a customer resource.
type Customer struct {
	CustomerID    *string        `json:"customerId,omitempty"`
	FirstName     *string        `json:"firstName,omitempty"`
	LastName      *string        `json:"lastName,omitempty"`
	EmailAddress  *string        `json:"emailAddress,omitempty"`
	DateOfBirth   *customTime    `json:"dateOfBirth,omitempty"`
	PostalAddress *postalAddress `json:"postalAddress,omitempty"`
	StreetAddress *streetAddress `json:"streetAddress,omitempty"`
	PhoneNumbers  []*phoneNumber `json:"phoneNumbers,omitempty"`
}

type postalAddress struct {
	AddressLine1 *string `json:"addressLine1,omitempty"`
	AddressLine2 *string `json:"addressLine2,omitempty"`
	AddressLine3 *string `json:"addressLine3,omitempty"`
	AddressLine4 *string `json:"addressLine4,omitempty"`
	Country      *string `json:"country,omitempty"`
	ZipCode      *string `json:"zipCode,omitempty"`
	City         *string `json:"city,omitempty"`
}

type streetAddress struct {
	AddressLine1 *string `json:"addressLine1,omitempty"`
	AddressLine2 *string `json:"addressLine2,omitempty"`
	AddressLine3 *string `json:"addressLine3,omitempty"`
	AddressLine4 *string `json:"addressLine4,omitempty"`
	Country      *string `json:"country,omitempty"`
	ZipCode      *string `json:"zipCode,omitempty"`
	City         *string `json:"city,omitempty"`
}

type phoneNumber struct {
	CountryCode *string `json:"countryCode,omitempty"`
	Number      *string `json:"number,omitempty"`
}

// GetCustomer gets information about a customer.
func (as *CustomersService) GetCustomer() (Customer, error) {
	var response customerResponse
	_, err := as.client.get(as.client.config.CustomersEndpoint, nil, &response)

	if response.IsError != nil && *response.IsError == true {
		return Customer{}, errors.New(*response.ErrorMessage)
	}

	return *response.Item, err
}

// Need this because the customer response contain an unsupported time format for birthDate.
type customTime struct {
	*time.Time
}

func (ct *customTime) UnmarshalJSON(b []byte) (err error) {
	t, err := time.Parse(time.RFC3339, strings.Trim(string(b), "\"")+"Z")
	*ct = customTime{&t}
	return err
}
