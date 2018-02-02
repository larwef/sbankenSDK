package common

import "github.com/larwef/sbankenSDK/client"

type Repository struct {
	Url    string
	Client *client.SbankenClient
}
