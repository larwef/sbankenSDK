# Sbanken Go SDK
Implements Accounts, Transaction and Transfer.

API-documentation: https://api.sbanken.no/Bank/swagger/index.html

Code examples in other languages: https://github.com/Sbanken/api-examples

### Configuration
Example configuration:
```
{
    "customerId": "<social security number>",
    "identityServer": "https://api.sbanken.no/identityserver/connect/token",
    "clientId": "<application client id>",
    "clientSecret": "<application client secret>",
    "accountsEndpoint": "https://api.sbanken.no/bank/api/v1/Accounts/",
    "transactionsEndpoint": "https://api.sbanken.no/bank/api/v1/Transactions/",
    "transfersEndpoint": "https://api.sbanken.no/bank/api/v1/Transfers/"
}
```

### How to create a http.Client with Oauth
```
import (
    "golang.org/x/oauth2/clientcredentials"
    "golang.org/x/net/context"
    "github.com/larwef/sbankenSDK"
)

func main() {
    config := sbankenSDK.ConfigFromFile("./config.json")
    
    oauthConfig := clientcredentials.Config{
        ClientID:     config.ClientId,
        ClientSecret: config.ClientSecret,
        TokenURL:     config.IdentityServer,
        Scopes:       []string{},
    }
    
    oauthClient := oauthConfig.Client(context.Background())
    
    sbankenClient := sbankenSDK.NewClient(oauthClient, config)
}
```