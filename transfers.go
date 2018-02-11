package sbankenSDK

import (
	"errors"
)

type TransferService service

type transferResponse struct {
	sbankenError
}

type TransferRequest struct {
	FromAccount string  `json:"fromAccount,omitempty"`
	ToAccount   string  `json:"toAccount,omitempty"`
	Amount      float64 `json:"amount,omitempty"`
	Message     string  `json:"message,omitempty"`
}

// Transfer funds from one account to another
func (ts *TransferService) Transfer(customerId string, transferRequest TransferRequest) error {
	var transferRsp transferResponse
	_, err := ts.client.post(ts.client.config.TransfersEndpoint+customerId, nil, transferRequest, &transferRsp)

	if transferRsp.IsError == true {
		return errors.New(transferRsp.ErrorMessage)
	}

	return err
}
