// Implements Accounts, Transaction and Transfer.
//
// API-documentation: https://api.sbanken.no/Bank/swagger/index.html
package sbankenSDK

import (
	"encoding/json"
	"errors"
)

type AccountService service

type accountsResponse struct {
	AvailableItems int       `json:"availableItems,omitempty"`
	Items          []Account `json:"items,omitempty"`
	sbankenError
}

type accountResponse struct {
	Item Account `json:"item,omitempty"`
	sbankenError
}

type Account struct {
	AccountNumber   string  `json:"accountNumber,omitempty"`
	CustomerId      string  `json:"customerId,omitempty"`
	OwnerCustomerId string  `json:"ownerCustomerId,omitempty"`
	Name            string  `json:"name,omitempty"`
	AccountType     string  `json:"accountType,omitempty"`
	Available       float64 `json:"available,omitempty"`
	Balance         float64 `json:"balance,omitempty"`
	CreditLimit     float64 `json:"creditLimit,omitempty"`
	DefaultAccount  bool    `json:"defaultAccount,omitempty"`
}

// Gets all accounts for user.
func (as *AccountService) GetAccounts(customerId string) ([]Account, error) {
	response, err := as.client.Get(as.client.config.AccountsEndpoint+customerId, nil)
	defer response.Body.Close()
	if err != nil {
		return []Account{}, err
	}

	var accountsRsp accountsResponse
	err = json.NewDecoder(response.Body).Decode(&accountsRsp)
	if accountsRsp.IsError == true {
		return accountsRsp.Items, errors.New(accountsRsp.ErrorMessage)
	}

	return accountsRsp.Items, err
}

// Gets information about a specified account.
func (as *AccountService) GetAccount(customerId string, accountNumber string) (Account, error) {
	response, err := as.client.Get(as.client.config.AccountsEndpoint+customerId+"/"+accountNumber, nil)
	defer response.Body.Close()
	if err != nil {
		return Account{}, err
	}

	var accountRsp accountResponse
	err = json.NewDecoder(response.Body).Decode(&accountRsp)
	if accountRsp.IsError == true {
		return accountRsp.Item, errors.New(accountRsp.ErrorMessage)
	}

	return accountRsp.Item, err
}
