package auth

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

	creds := types.ClientCredentials{
		ClientId:     "mock client id",
		ClientSecret: "mock client secret",
		Username:     "mock username",
		Password:     "mock password",
	}
	a := &authenticator{
		baseUrl: server.URL,
		creds:   creds,
	}

	mux.HandleFunc("/security/oauth/token",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
			testFormValues(t, r, values{
				"grant_type":    "password",
				"client_id":     creds.ClientId,
				"client_secret": creds.ClientSecret,
				"username":      creds.Username,
				"password":      creds.Password,
			})
			fmt.Fprint(w, `{
				"access_token":"mock access token",
				"token_type":"bearer",
				"refresh_token":"mock refresh token",
				"expires_in":3599
			}`)
		},
	)

	got, err := a.Authenticate()
	if err != nil {
		t.Errorf("authenticator.Authenticate failed: %v", err)
	}

	want := oauthResponse{
		AccessToken:  "mock access token",
		TokenType:    "bearer",
		RefreshToken: "mock refresh token",
		ExpiresIn:    3599,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("authenticator.Authenticate returned %+v, want %+v",
			got, want)
	}
}

func TestAuthenticateHttpError(t *testing.T) {
	setUp()
	defer tearDown()

	creds := types.ClientCredentials{
		ClientId:     "mock client id",
		ClientSecret: "mock client secret",
		Username:     "mock username",
		Password:     "mock password",
	}
	a := &authenticator{
		baseUrl: server.URL,
		creds:   creds,
	}

	_, err := a.Authenticate()
	if err == nil {
		t.Error("authenticator.Authenticate should fail when server returns HTTP error")
	}
}

func TestAuthenticateFailedResponse(t *testing.T) {
	setUp()
	defer tearDown()

	creds := types.ClientCredentials{
		ClientId:     "mock client id",
		ClientSecret: "mock client secret",
		Username:     "mock username",
		Password:     "mock password",
	}
	a := &authenticator{
		baseUrl: server.URL,
		creds:   creds,
	}

	mux.HandleFunc("/security/oauth/token",
		func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "mock server error: request failed", 500)
		},
	)

	_, err := a.Authenticate()
	if err == nil {
		t.Error("authenticator.Authenticate should fail when server returns invalid response")
	}
}
