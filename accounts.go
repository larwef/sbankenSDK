// Implements Accounts, Transaction and Transfer.
//
// API-documentation: https://api.sbanken.no/Bank/swagger/index.html
package sbankenSDK

import (
	"encoding/json"
	"errors"
	"github.com/larwef/sbankenSDK/authentication"
	"github.com/larwef/sbankenSDK/client"
	"github.com/larwef/sbankenSDK/common"
)

type AccountsRepository struct {
	common.Repository
}

type accountsResponse struct {
	AvailableItems int       `json:"availableItems"`
	Items          []Account `json:"items"`
	common.Error
}

type accountResponse struct {
	Item Account `json:"item"`
	common.Error
}

type Account struct {
	AccountNumber   string  `json:"accountNumber"`
	CustomerId      string  `json:"customerId"`
	OwnerCustomerId string  `json:"ownerCustomerId"`
	Name            string  `json:"name"`
	AccountType     string  `json:"accountType"`
	Available       float64 `json:"available"`
	Balance         float64 `json:"balance"`
	CreditLimit     float64 `json:"creditLimit"`
	DefaultAccount  bool    `json:"defaultAccount"`
}

// Constructor for AccountRepository
func NewAccountRepository(config Config) *AccountsRepository {
	token := authentication.NewSbankenToken(config.IdentityServer, config.ClientId, config.ClientSecret)
	return &AccountsRepository{common.Repository{Url: config.AccountsEndpoint, Client: client.NewSbankenClient(&token)}}
}

// Gets all accounts for user.
func (ar *AccountsRepository) GetAccounts(customerId string) ([]Account, error) {
	response, err := ar.Client.Get(ar.Url+customerId, nil)
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
func (ar *AccountsRepository) GetAccount(customerId string, accountNumber string) (Account, error) {
	response, err := ar.Client.Get(ar.Url+customerId+"/"+accountNumber, nil)
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
