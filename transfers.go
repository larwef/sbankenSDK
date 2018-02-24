package sbankenSDK

type TransferService service

type Transfer entity

type TransferRequest entity

// Transfer funds from one account to another
func (ts *TransferService) Transfer(customerId string, transferRequest TransferRequest) error {
	var transferRsp response
	_, err := ts.client.post(ts.client.config.TransfersEndpoint+customerId, nil, transferRequest.properties, &transferRsp.properties)

	if err := transferRsp.getError(); err != nil {
		return err
	}

	return err
}

func NewTransferRequest() (*TransferRequest) {
	return &TransferRequest{properties: make(map[string]interface{})}
}

func (t *TransferRequest) WithFromAccount(fromAccount string) (*TransferRequest) {
	t.properties[FROM_ACCOUNT] = fromAccount
	return t
}

func (t *TransferRequest) WithToAccount(toAccount string) (*TransferRequest) {
	t.properties[TO_ACCOUNT] = toAccount
	return t
}

func (t *TransferRequest) WithAmount(amount float64) (*TransferRequest) {
	t.properties[AMOUNT] = amount
	return t
}

func (t *TransferRequest) WithMessage(message string) (*TransferRequest) {
	t.properties[MESSAGE] = message
	return t
}
