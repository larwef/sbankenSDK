package sbankenSDK

import (
	"errors"
)

type TransferService service

type transferResponse struct {
	sbankenError
}

type TransferRequest struct {
	FromAccountId *string  `json:"fromAccountId,omitempty"`
	ToAccountId   *string  `json:"toAccountId,omitempty"`
	Amount        *float64 `json:"amount,omitempty"`
	Message       *string  `json:"message,omitempty"`
}

// Transfer funds from one account to another
func (ts *TransferService) Transfer(transferRequest TransferRequest) error {
	var transferRsp transferResponse
	_, err := ts.client.post(ts.client.config.TransfersEndpoint, nil, transferRequest, &transferRsp)

	if transferRsp.IsError != nil && *transferRsp.IsError == true {
		return errors.New(*transferRsp.ErrorMessage)
	}

	return err
}
