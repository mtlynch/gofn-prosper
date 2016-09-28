package auth

import (
	"time"

	"github.com/mtlynch/gofn-prosper/types"
)

// OAuthToken is an authentication token from Prosper that is valid for a
// limited amount of time.
type OAuthToken struct {
	AccessToken  string
	TokenType    string
	RefreshToken string
	Expiration   time.Time
}

// TokenManager manages the OAuth tokens for the Prosper API, providing a valid
// token to the caller and refreshing it when needed.
type TokenManager interface {
	Token() (OAuthToken, error)
}

type defaultTokenManager struct {
	token         OAuthToken
	authenticator ProsperAuthenticator
	clock         types.Clock
}

// NewTokenManager creates a new TokenManager instance that authenticates to
// Propser with the given authenticator.
func NewTokenManager(authenticator ProsperAuthenticator) TokenManager {
	return &defaultTokenManager{
		token:         OAuthToken{},
		authenticator: authenticator,
		clock:         types.DefaultClock{},
	}
}

// Token returns a valid OAuth token, retrieving a new one from the Propser
// server if necessary.
func (m *defaultTokenManager) Token() (OAuthToken, error) {
	if m.clock.Now().Before(m.token.Expiration) {
		return m.token, nil
	}
	token, err := m.tokenFromAuthenticator()
	if err != nil {
		return OAuthToken{}, err
	}
	m.token = token
	return m.token, nil
}

func (m defaultTokenManager) tokenFromAuthenticator() (token OAuthToken, err error) {
	response, err := m.authenticator.Authenticate()
	if err != nil {
		return OAuthToken{}, err
	}
	expiration := m.clock.Now().Add((time.Duration(response.ExpiresIn) * time.Second))
	return OAuthToken{
		AccessToken:  response.AccessToken,
		TokenType:    response.TokenType,
		RefreshToken: response.RefreshToken,
		Expiration:   expiration,
	}, nil
}
