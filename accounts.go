package sbankenSDK

type AccountService service

type Account entity

// Gets all accounts for user.
func (as *AccountService) GetAccounts(customerId string) ([]Account, error) {
	var accountRsp response
	_, err := as.client.get(as.client.config.AccountsEndpoint+customerId, nil, &accountRsp.properties)
	if err != nil {
		return nil, err
	}

	if err := accountRsp.getError(); err != nil {
		return nil, err
	}

	items, err := getRequiredProperty(ITEMS, accountRsp.properties)
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
	var accountRsp response
	_, err := as.client.get(as.client.config.AccountsEndpoint+customerId+"/"+accountNumber, nil, &accountRsp.properties)
	if err != nil {
		return Account{}, err
	}

	if err := accountRsp.getError(); err != nil {
		return Account{}, err
	}

	val, err := getRequiredProperty(ITEM, accountRsp.properties)
	if err != nil {
		return Account{}, err
	}

	return Account{properties: val.(map[string]interface{})}, err
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
