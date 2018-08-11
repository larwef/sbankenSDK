package sbankensdk

import (
	"errors"
)

// TransferService handles communication with the Transfers part of the API.
type TransferService service

type transferResponse struct {
	sbankenError
}

// TransferRequest holds parameters for transfering funds from one account to another.
type TransferRequest struct {
	FromAccountID *string  `json:"fromAccountId,omitempty"`
	ToAccountID   *string  `json:"toAccountId,omitempty"`
	Amount        *float64 `json:"amount,omitempty"`
	Message       *string  `json:"message,omitempty"`
}

// Transfer funds from one account to another
func (ts *TransferService) Transfer(transferRequest TransferRequest) error {
	var response transferResponse
	_, err := ts.client.post(ts.client.config.TransfersEndpoint, nil, transferRequest, &response)

	if response.IsError != nil && *response.IsError == true {
		return errors.New(*response.ErrorMessage)
	}

	return err
}
