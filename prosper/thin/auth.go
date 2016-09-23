package thin

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/mtlynch/gofn-prosper/types"
)

type oauthResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

func (p unauthenticatedClient) Authenticate(credentials types.ClientCredentials) (response oauthResponse, err error) {
	resp, err := http.PostForm(p.baseUrl+"/security/oauth/token",
		url.Values{
			"grant_type":    {"password"},
			"client_id":     {credentials.ClientId},
			"client_secret": {credentials.ClientSecret},
			"username":      {credentials.Username},
			"password":      {credentials.Password},
		})
	if err != nil {
		return oauthResponse{}, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return oauthResponse{}, err
	}
	return response, nil
}

func (c Client) Token() (string, error) {
	token, err := c.tokenManager.Token()
	if err != nil {
		return "", err
	}
	return token.AccessToken, nil
}
