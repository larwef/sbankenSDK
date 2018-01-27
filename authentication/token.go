package authentication

import (
	"time"
)

type Token interface {
	GetTokenString() (string)
	GetExpirationTime() (time.Time)
	GetTokenType() (string)
	RefreshToken()
}
