package authentication

import (
	"time"
)

type tokenBase struct {
	tokenString string
	validTo     time.Time
	tokenType   string
}
