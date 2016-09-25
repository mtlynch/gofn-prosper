package thin

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/mtlynch/gofn-prosper/types"
)

func TestAuthenticateSuccessfulResponse(t *testing.T) {
	setUp()
	defer tearDown()

	unauthenticatedClient := &unauthenticatedClient{
		baseUrl: server.URL,
	}
	credentials := types.ClientCredentials{
		ClientId:     "mock client id",
		ClientSecret: "mock client secret",
		Username:     "mock username",
		Password:     "mock password",
	}

	mux.HandleFunc("/security/oauth/token",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
			testFormValues(t, r, values{
				"grant_type":    "password",
				"client_id":     credentials.ClientId,
				"client_secret": credentials.ClientSecret,
				"username":      credentials.Username,
				"password":      credentials.Password,
			})
			fmt.Fprint(w, `{
				"access_token":"mock access token",
				"token_type":"bearer",
				"refresh_token":"mock refresh token",
				"expires_in":3599
			}`)
		},
	)

	got, err := unauthenticatedClient.Authenticate(credentials)
	if err != nil {
		t.Errorf("client.Authenticate failed: %v", err)
	}

	want := oauthResponse{
		AccessToken:  "mock access token",
		TokenType:    "bearer",
		RefreshToken: "mock refresh token",
		ExpiresIn:    3599,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("client.Authenticate returned %+v, want %+v",
			got, want)
	}
}

func TestAuthenticateHttpError(t *testing.T) {
	setUp()
	defer tearDown()

	unauthenticatedClient := &unauthenticatedClient{
		baseUrl: server.URL,
	}
	credentials := types.ClientCredentials{
		ClientId:     "mock client id",
		ClientSecret: "mock client secret",
		Username:     "mock username",
		Password:     "mock password",
	}

	_, err := unauthenticatedClient.Authenticate(credentials)
	if err == nil {
		t.Error("client.Authenticate should fail when server returns HTTP error")
	}
}

func TestAuthenticateFailedResponse(t *testing.T) {
	setUp()
	defer tearDown()

	unauthenticatedClient := &unauthenticatedClient{
		baseUrl: server.URL,
	}
	credentials := types.ClientCredentials{
		ClientId:     "mock client id",
		ClientSecret: "mock client secret",
		Username:     "mock username",
		Password:     "mock password",
	}

	mux.HandleFunc("/security/oauth/token",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "mock server error: request failed")
		},
	)

	_, err := unauthenticatedClient.Authenticate(credentials)
	if err == nil {
		t.Error("client.Authenticate should fail when server returns invalid response")
	}
}
