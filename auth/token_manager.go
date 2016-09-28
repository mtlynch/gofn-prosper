package auth

import (
	"time"

	"github.com/mtlynch/gofn-prosper/types"
)

type oauthToken struct {
	AccessToken  string
	TokenType    string
	RefreshToken string
	Expiration   time.Time
}

// TokenManager manages the OAuth tokens for the Prosper API, providing a valid
// token to the caller and refreshing it when needed.
type TokenManager interface {
	Token() (oauthToken, error)
}

type defaultTokenManager struct {
	token         oauthToken
	authenticator ProsperAuthenticator
	clock         types.Clock
}

// NewTokenManager creates a new TokenManager instance that authenticates to
// Propser with the given authenticator.
func NewTokenManager(authenticator ProsperAuthenticator) TokenManager {
	return &defaultTokenManager{
		token:         oauthToken{},
		authenticator: authenticator,
		clock:         types.DefaultClock{},
	}
}

// Token returns a valid OAuth token, retrieving a new one from the Propser
// server if necessary.
func (m *defaultTokenManager) Token() (oauthToken, error) {
	if m.clock.Now().Before(m.token.Expiration) {
		return m.token, nil
	}
	token, err := m.tokenFromAuthenticator()
	if err != nil {
		return oauthToken{}, err
	}
	m.token = token
	return m.token, nil
}

func (m defaultTokenManager) tokenFromAuthenticator() (token oauthToken, err error) {
	response, err := m.authenticator.Authenticate()
	if err != nil {
		return oauthToken{}, err
	}
	expiration := m.clock.Now().Add((time.Duration(response.ExpiresIn) * time.Second))
	return oauthToken{
		AccessToken:  response.AccessToken,
		TokenType:    response.TokenType,
		RefreshToken: response.RefreshToken,
		Expiration:   expiration,
	}, nil
}
