package sbankenSDK

import (
	"fmt"
	"errors"
	"time"
)

// Json fields
const (
	AVAILABLE_ITEMS      = "availableItems"
	ITEMS                = "items"
	ITEM                 = "item"
	OTHER_ACCOUNT_NUMBER = "otherAccountNumber"
	AMOUNT               = "amount"
	TEXT                 = "text"
	TRANSACTION_TYPE     = "transactionType"
	REGISTRATION_DATE    = "registrationDate"
	ACCOUNTING_DATE      = "accountingDate"
	INTEREST_DATE        = "interestDate"
	ACCOUNT_NUMBER       = "accountNumber"
	TRANSACTION_ID       = "transactionId"
	CUSTOMER_ID          = "customerId"
	OWNER_CUSTOMER_ID    = "ownerCustomerId"
	NAME                 = "name"
	ACCOUNT_TYPE         = "accountType"
	AVAILABLE            = "available"
	BALANCE              = "balance"
	CREDIT_LIMIT         = "creditLimit"
	DEFAULT_ACCOUNT      = "defaultAccount"
	ERROR_TYPE           = "errorType"
	IS_ERROR             = "isError"
	ERROR_MESSAGE        = "errorMessage"
	TRACE_ID             = "traceId"
	FROM_ACCOUNT         = "fromAccount"
	TO_ACCOUNT           = "toAccount"
	MESSAGE              = "message"
)

type entity struct {
	properties map[string]interface{}
}

type response entity

func (ar *response) getError() (error) {
	isError, err := getRequiredProperty(IS_ERROR, ar.properties)
	if err != nil {
		return err
	}

	if isError.(bool) {
		errorMessage, err := getRequiredProperty(ERROR_MESSAGE, ar.properties)
		if err != nil {
			return err
		}
		return errors.New(errorMessage.(string))
	}
	return nil
}

func getString(key string, m map[string]interface{}) (val string, isSet bool, isNull bool) {
	var v interface{}
	if v, isSet, isNull = getInterface(key, m); isSet && !isNull {
		val = v.(string)
	}
	return
}

func getInt(key string, m map[string]interface{}) (val int, isSet bool, isNull bool) {
	var v interface{}
	if v, isSet, isNull = getInterface(key, m); isSet && !isNull {
		val = int(v.(float64))
	}
	return
}

func getFloat64(key string, m map[string]interface{}) (val float64, isSet bool, isNull bool) {
	var v interface{}
	if v, isSet, isNull = getInterface(key, m); isSet && !isNull {
		val = v.(float64)
	}
	return
}

func getBool(key string, m map[string]interface{}) (val bool, isSet bool, isNull bool) {
	var v interface{}
	if v, isSet, isNull = getInterface(key, m); isSet && !isNull {
		val = v.(bool)
	}
	return
}

func getTime(key string, m map[string]interface{}) (time.Time, bool, bool) {
	v, isSet, isNull := getInterface(key, m)
	if isSet && !isNull {
		if val, err := time.Parse(time.RFC3339, v.(string)); err != nil {
			return val, isSet, isNull
		}
	}
	return time.Time{}, isSet, isNull
}

func getInterfaceArray(key string, m map[string]interface{}) (val []interface{}, isSet bool, isNull bool) {
	var v interface{}
	if v, isSet, isNull = getInterface(key, m); isSet && !isNull {
		val = v.([]interface{})
	}
	return
}

func getInterface(key string, m map[string]interface{}) (val interface{}, isSet bool, isNull bool) {
	val, isSet = m[key]
	// If not set, do not assume null
	isNull = val == nil && isSet
	return
}

func getRequiredProperty(key string, m map[string]interface{}) (interface{}, error) {
	val, isSet, isNull := getInterface(key, m)
	if !isSet || isNull {
		return nil, fmt.Errorf("[%s] is either missing or null. isSet: %t, isNull: %t", key, isSet, isNull)
	}
	return val, nil
}
