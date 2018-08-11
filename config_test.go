package sbankensdk

import (
	"testing"
)

const (
	customerID           = "customer1"
	identityServer       = "https://api.sbanken.no/identityserver/connect/token"
	clientID             = "client1"
	clientSecret         = "secret1"
	accountsEndpoint     = "https://api.sbanken.no/bank/api/v1/Accounts/"
	transactionsEndpoint = "https://api.sbanken.no/bank/api/v1/Transactions/"
	transfersEndpoint    = "https://api.sbanken.no/bank/api/v1/Transfers/"
	customersEndpoint    = "https://api.sbanken.no/customers/api/v1/Customers"
)

func TestConfigFromFile(t *testing.T) {
	config := ConfigFromFile("testdata/config.json")

	assertEqual(t, config.CustomerID, customerID)
	assertEqual(t, config.IdentityServer, identityServer)
	assertEqual(t, config.ClientID, clientID)
	assertEqual(t, config.ClientSecret, clientSecret)
	assertEqual(t, config.AccountsEndpoint, accountsEndpoint)
	assertEqual(t, config.TransactionsEndpoint, transactionsEndpoint)
	assertEqual(t, config.TransfersEndpoint, transfersEndpoint)
	assertEqual(t, config.CustomersEndpoint, customersEndpoint)
}
