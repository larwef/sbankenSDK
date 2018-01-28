package sbankenSDK

import (
	"github.com/larwef/sbankenSDK/common"
	"github.com/larwef/sbankenSDK/client"
	"github.com/larwef/sbankenSDK/authentication"
	"encoding/json"
	"errors"
)

type transfersRepository struct {
	url    string
	client *client.SbankenClient
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

func NewTranfsersRepository(config Config) (*transfersRepository) {
	token := authentication.NewSbankenToken(config.IdentityServer, config.ClientId, config.ClientSecret)
	return &transfersRepository{url: config.TransfersEndpoint, client: client.NewSbankenClient(&token)}
}

func (tr transfersRepository) Transfer(customerId string, transferRequest TranferRequest) (error) {
	var transferRsp transferResponse

	payload, err := json.Marshal(transferRequest)
	if err != nil {
		return err
	}

	response, err := tr.client.Post(tr.url+customerId, nil, payload)
	if err != nil {
		return err
	}

	json.Unmarshal(response, &transferRsp)
	if transferRsp.IsError == true {
		return errors.New(transferRsp.ErrorMessage)
	}

	return err
}
