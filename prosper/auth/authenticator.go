// Package auth provides APIs for authenticating to Prosper. This includes
// authenticating to Prosper using the user's credentials as well as managing
// and refreshing the user's OAuth token.
package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/mtlynch/gofn-prosper/types"
)

const baseProsperUrl = "https://api.prosper.com/v1"

// ProsperAuthenticator authenticates to Prosper to retrieve a raw OAuth
// response.
type ProsperAuthenticator interface {
	Authenticate() (oauthResponse, error)
}

type authenticator struct {
	baseUrl string
	creds   types.ClientCredentials
}

// NewAuthenticator creates a new, unauthenticated Prosper API client with the
// given Prosper credentials.
func NewAuthenticator(creds types.ClientCredentials) ProsperAuthenticator {
	return &authenticator{
		baseUrl: baseProsperUrl,
		creds:   creds,
	}
}

type oauthResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

// Authenticate authenticates to the Prosper API server and retrieves a raw
// OAuth response.
func (c authenticator) Authenticate() (response oauthResponse, err error) {
	resp, err := http.PostForm(c.baseUrl+"/security/oauth/token",
		url.Values{
			"grant_type":    {"password"},
			"client_id":     {c.creds.ClientId},
			"client_secret": {c.creds.ClientSecret},
			"username":      {c.creds.Username},
			"password":      {c.creds.Password},
		})
	if err != nil {
		return oauthResponse{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return oauthResponse{}, errors.New("Prosper server error: " + resp.Status)
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return oauthResponse{}, err
	}
	return response, nil
}
