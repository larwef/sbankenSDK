// Implements Accounts, Transaction and Transfer.
//
// API-documentation: https://api.sbanken.no/Bank/swagger/index.html
package sbankenSDK

import "errors"

type AccountService service

type accountsResponse struct {
	AvailableItems *int      `json:"availableItems,omitempty"`
	Items          []Account `json:"items,omitempty"`
	sbankenError
}

type accountResponse struct {
	Item *Account `json:"item,omitempty"`
	sbankenError
}

type Account struct {
	AccountId       *string  `json:"accountId,omitempty"`
	AccountNumber   *string  `json:"accountNumber,omitempty"`
	OwnerCustomerId *string  `json:"ownerCustomerId,omitempty"`
	Name            *string  `json:"name,omitempty"`
	AccountType     *string  `json:"accountType,omitempty"`
	Available       *float64 `json:"available,omitempty"`
	Balance         *float64 `json:"balance,omitempty"`
	CreditLimit     *float64 `json:"creditLimit,omitempty"`
}

// Gets all accounts for user.
func (as *AccountService) GetAccounts() ([]Account, error) {
	var accountsRsp accountsResponse
	_, err := as.client.get(as.client.config.AccountsEndpoint, nil, &accountsRsp)

	if accountsRsp.IsError != nil && *accountsRsp.IsError == true {
		return nil, errors.New(*accountsRsp.ErrorMessage)
	}

	return accountsRsp.Items, err
}

// Gets information about a specified account.
func (as *AccountService) GetAccount(accountId string) (Account, error) {
	var accountRsp accountResponse
	_, err := as.client.get(as.client.config.AccountsEndpoint+accountId, nil, &accountRsp)

	if *accountRsp.IsError == true {
		return Account{}, errors.New(*accountRsp.ErrorMessage)
	}

	return *accountRsp.Item, err
}
