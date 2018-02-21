package sbankenSDK

import (
	"fmt"
	"errors"
)

const (
	ERROR_TYPE    = "errorType"
	IS_ERROR      = "isError"
	ERROR_MESSAGE = "errorMessage"
	TRACE_ID      = "traceId"
)

type entity struct {
	properties map[string]interface{}
}

type response entity

func (ar response) getError() (error) {
	isError, err := getRequiredProperty(IS_ERROR, ar.properties)
	if err != nil {
		return err
	}

	if isError.(bool) {
		errorMessage, err:= getRequiredProperty(ERROR_MESSAGE, ar.properties)
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
		return nil, fmt.Errorf("%s is either missing or null. isSet: %t, isNull: %t", key, isSet, isNull)
	}
	return val, nil
}
