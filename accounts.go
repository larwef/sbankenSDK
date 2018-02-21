package sbankenSDK

import (
	"encoding/json"
)

const (
	AVAILABLE_ITEMS = "availableItems"
	ITEMS           = "items"
	ITEM            = "item"

	ACCOUNT_NUMBER    = "accountNumber"
	CUSTOMER_ID       = "customerId"
	OWNER_CUSTOMER_ID = "ownerCustomerId"
	NAME              = "name"
	ACCOUNT_TYPE      = "accountType"
	AVAILABLE         = "available"
	BALANCE           = "balance"
	CREDIT_LIMIT      = "creditLimit"
	DEFAULT_ACCOUNT   = "defaultAccount"
)

type AccountService service

type Account entity

// Gets all accounts for user.
func (as *AccountService) GetAccounts(customerId string) ([]Account, error) {
	response, err := as.getResponse(as.client.config.AccountsEndpoint + customerId)
	if err != nil {
		return nil, err
	}

	items, err := getRequiredProperty(ITEMS, response.properties)
	if err != nil {
		return nil, err
	}

	var val []Account
	for _, element := range items.([]interface{}) {
		val = append(val, Account{properties: element.(map[string]interface{})})
	}

	return val, err
}

// Gets information about a specified account.
func (as *AccountService) GetAccount(customerId string, accountNumber string) (Account, error) {
	response, err := as.getResponse(as.client.config.AccountsEndpoint + customerId + "/" + accountNumber)
	if err != nil {
		return Account{}, err
	}

	val, err := getRequiredProperty(ITEM, response.properties)
	if err != nil {
		return Account{}, err
	}

	return Account{properties: val.(map[string] interface{})}, err
}

func (as* AccountService) getResponse(url string) (response, error) {
	var accountRsp response
	_, err := as.client.get(url, nil, &accountRsp.properties)
	if err != nil {
		return response{}, err
	}

	if err := accountRsp.getError(); err != nil {
		return response{}, err
	}

	return accountRsp, nil
}

func (a *Account) GetAccountNumber() (val string, isSet bool, isNull bool) {
	return getString(ACCOUNT_NUMBER, a.properties)
}

func (a *Account) GetCustomerId() (val string, isSet bool, isNull bool) {
	return getString(CUSTOMER_ID, a.properties)
}

func (a *Account) GetOwnerCustomerId() (val string, isSet bool, isNull bool) {
	return getString(OWNER_CUSTOMER_ID, a.properties)
}

func (a *Account) GetName() (val string, isSet bool, isNull bool) {
	return getString(NAME, a.properties)
}

func (a *Account) GetAccountType() (val string, isSet bool, isNull bool) {
	return getString(ACCOUNT_TYPE, a.properties)
}

func (a *Account) GetAvailable() (val float64, isSet bool, isNull bool) {
	return getFloat64(AVAILABLE, a.properties)
}

func (a *Account) GetBalance() (val float64, isSet bool, isNull bool) {
	return getFloat64(BALANCE, a.properties)
}

func (a *Account) GetCreditLimit() (val float64, isSet bool, isNull bool) {
	return getFloat64(CREDIT_LIMIT, a.properties)
}

func (a *Account) GetDefaultAccount() (val bool, isSet bool, isNull bool) {
	return getBool(DEFAULT_ACCOUNT, a.properties)
}

func (a *Account) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &a.properties)
}
