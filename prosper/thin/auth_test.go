package thin

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/mtlynch/gofn-prosper/types"
)

type values map[string]string

var (
	mux    *http.ServeMux
	server *httptest.Server
)

func setUp() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
}

func tearDown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func testContentType(t *testing.T, r *http.Request, want string) {
	contentType, ok := r.Header["Content-Type"]
	if !ok || len(contentType) < 1 {
		t.Errorf("Content-Type header not present, want: %v", want)
		return
	}
	if got := r.Header["Content-Type"][0]; got != want {
		t.Errorf("Content-type: %v, want %v", got, want)
	}
}

func testFormValues(t *testing.T, r *http.Request, values values) {
	want := url.Values{}
	for k, v := range values {
		want.Add(k, v)
	}

	r.ParseForm()
	if got := r.Form; !reflect.DeepEqual(got, want) {
		t.Errorf("Request parameters: %v, want %v", got, want)
	}
}

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
