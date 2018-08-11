package sbankensdk

import "errors"

// AccountService handles communication with the account part of the API.
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

// Account represents an account resource.
type Account struct {
	AccountID       *string  `json:"accountId,omitempty"`
	AccountNumber   *string  `json:"accountNumber,omitempty"`
	OwnerCustomerID *string  `json:"ownerCustomerId,omitempty"`
	Name            *string  `json:"name,omitempty"`
	AccountType     *string  `json:"accountType,omitempty"`
	Available       *float64 `json:"available,omitempty"`
	Balance         *float64 `json:"balance,omitempty"`
	CreditLimit     *float64 `json:"creditLimit,omitempty"`
}

// GetAccounts gets all accounts for user.
func (as *AccountService) GetAccounts() ([]Account, error) {
	var accountsRsp accountsResponse
	_, err := as.client.get(as.client.config.AccountsEndpoint, nil, &accountsRsp)

	if accountsRsp.IsError != nil && *accountsRsp.IsError == true {
		return nil, errors.New(*accountsRsp.ErrorMessage)
	}

	return accountsRsp.Items, err
}

// GetAccount gets information about a specified account.
func (as *AccountService) GetAccount(accountID string) (*Account, error) {
	var response accountResponse
	_, err := as.client.get(as.client.config.AccountsEndpoint+accountID, nil, &response)

	if response.IsError != nil && *response.IsError == true {
		return &Account{}, errors.New(*response.ErrorMessage)
	}

	return response.Item, err
}
