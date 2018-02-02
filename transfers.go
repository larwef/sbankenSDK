package sbankenSDK

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/larwef/sbankenSDK/authentication"
	"github.com/larwef/sbankenSDK/client"
	"github.com/larwef/sbankenSDK/common"
)

type transfersRepository struct {
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

func NewTranfsersRepository(config Config) *transfersRepository {
	token := authentication.NewSbankenToken(config.IdentityServer, config.ClientId, config.ClientSecret)
	return &transfersRepository{common.Repository{Url: config.TransfersEndpoint, Client: client.NewSbankenClient(&token)}}
}

func (tr transfersRepository) Transfer(customerId string, transferRequest TranferRequest) error {
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
