package thin

import (
	"time"
)

const redisKeyOAuthToken = "oauth"

type oauthToken struct {
	AccessToken  string
	TokenType    string
	RefreshToken string
	Expiration   time.Time
}

type (
	tokenManager interface {
		Token() (oauthToken, error)
	}
)
