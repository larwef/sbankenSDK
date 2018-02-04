package sbankenSDK

import (
	"bytes"
	"encoding/json"
	"errors"
)

type TransferService service

type transferResponse struct {
	sbankenError
}

type TranferRequest struct {
	FromAccount string  `json:"fromAccount,omitempty"`
	ToAccount   string  `json:"toAccount,omitempty"`
	Amount      float64 `json:"amount,omitempty"`
	Message     string  `json:"message,omitempty"`
}

// Transfer funds from one account to another
func (ts *TransferService) Transfer(customerId string, transferRequest TranferRequest) error {
	payload, err := json.Marshal(transferRequest)
	if err != nil {
		return err
	}

	response, err := ts.client.Post(ts.client.config.TransfersEndpoint+customerId, nil, bytes.NewBuffer(payload))
	defer response.Body.Close()
	if err != nil {
		return err
	}

	var transferRsp transferResponse
	err = json.NewDecoder(response.Body).Decode(&transferRsp)
	if transferRsp.IsError == true {
		return errors.New(transferRsp.ErrorMessage)
	}

	return err
}
