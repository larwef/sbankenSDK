package sbankenSDK

import (
	"testing"
)

const (
	customerId           = "customer1"
	identityServer       = "https://api.sbanken.no/identityserver/connect/token"
	clientId             = "client1"
	clientSecret         = "secret1"
	accountsEndpoint     = "https://api.sbanken.no/bank/api/v1/Accounts/"
	transactionsEndpoint = "https://api.sbanken.no/bank/api/v1/Transactions/"
	transfersEndpoint    = "https://api.sbanken.no/bank/api/v1/Transfers/"
)

func TestConfigFromFile(t *testing.T) {
	config := ConfigFromFile("testdata/config.json")

	assertEqual(t, config.CustomerId, customerId)
	assertEqual(t, config.IdentityServer, identityServer)
	assertEqual(t, config.ClientId, clientId)
	assertEqual(t, config.ClientSecret, clientSecret)
	assertEqual(t, config.AccountsEndpoint, accountsEndpoint)
	assertEqual(t, config.TransactionsEndpoint, transactionsEndpoint)
	assertEqual(t, config.TransfersEndpoint, transfersEndpoint)
}
