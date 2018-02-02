package sbankenSDK

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/larwef/sbankenSDK/authentication"
	"github.com/larwef/sbankenSDK/client"
	"github.com/larwef/sbankenSDK/common"
)

type TransfersRepository struct {
	common.Repository
}

type transferResponse struct {
	common.Error
}

type TranferRequest struct {
	FromAccount string  `json:"fromAccount"`
	ToAccount   string  `json:"toAccount"`
	Amount      float64 `json:"amount"`
	Message     string  `json:"message"`
}

// Constructor for TransferRepository
func NewTranfsersRepository(config Config) *TransfersRepository {
	token := authentication.NewSbankenToken(config.IdentityServer, config.ClientId, config.ClientSecret)
	return &TransfersRepository{common.Repository{Url: config.TransfersEndpoint, Client: client.NewSbankenClient(&token)}}
}

// Transfer funds from one account to another
func (tr *TransfersRepository) Transfer(customerId string, transferRequest TranferRequest) error {
	payload, err := json.Marshal(transferRequest)
	if err != nil {
		return err
	}

	response, err := tr.Client.Post(tr.Url+customerId, nil, bytes.NewBuffer(payload))
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
