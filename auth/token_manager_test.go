package auth

import (
	"errors"
	"reflect"
	"testing"
	"time"
)

type mockProsperAuthenticator struct {
	OAuthResponse     oauthResponse
	Err               error
	AuthenticateCalls int
}

func (m *mockProsperAuthenticator) Authenticate() (oauthResponse, error) {
	m.AuthenticateCalls++
	return m.OAuthResponse, m.Err
}

func TestMultipleCallsWithinExpirationPeriodOnlyAuthenticateOnce(t *testing.T) {
	a := &mockProsperAuthenticator{
		OAuthResponse: oauthResponse{
			AccessToken:  "mock oauth token",
			TokenType:    "mock token type",
			RefreshToken: "mock refresh token",
			ExpiresIn:    3599,
		},
	}
	now := time.Date(2015, 12, 24, 10, 0, 0, 0, time.UTC)
	m := defaultTokenManager{
		authenticator: a,
		clock:         mockClock{&now},
	}

	got, err := m.Token()

	if err != nil {
		t.Errorf("Token() failed: %v", err)
	}
	want := OAuthToken{
		"mock oauth token",
		"mock token type",
		"mock refresh token",
		time.Date(2015, 12, 24, 10, 59, 59, 0, time.UTC),
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Token() returned: %+v, want: %+v", got, want)
	}
	if a.AuthenticateCalls != 1 {
		t.Errorf("Called Authenticate() unexpected times: %+v, want: %+v", a.AuthenticateCalls, 1)
	}

	// Now try retrieving a token again. It should not increase the Authenticate()
	// calls.
	got, err = m.Token()
	if err != nil {
		t.Errorf("Token() failed: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Token() returned: %+v, want: %+v", got, want)
	}
	if a.AuthenticateCalls != 1 {
		t.Errorf("Called Authenticate() unexpected times: %+v, want: %+v", a.AuthenticateCalls, 1)
	}
}

func TestCallAfterExpirationPeriodRefreshesToken(t *testing.T) {
	a := &mockProsperAuthenticator{
		OAuthResponse: oauthResponse{
			AccessToken:  "mock oauth token",
			TokenType:    "mock token type",
			RefreshToken: "mock refresh token",
			ExpiresIn:    3599,
		},
	}
	now := time.Date(2015, 12, 24, 10, 0, 0, 0, time.UTC)
	m := defaultTokenManager{
		authenticator: a,
		clock:         mockClock{&now},
	}

	got, err := m.Token()
	if err != nil {
		t.Errorf("Token() failed: %v", err)
	}
	want := OAuthToken{
		"mock oauth token",
		"mock token type",
		"mock refresh token",
		time.Date(2015, 12, 24, 10, 59, 59, 0, time.UTC),
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Token() returned: %+v, want: %+v", got, want)
	}
	if a.AuthenticateCalls != 1 {
		t.Errorf("Called Authenticate() unexpected times: %+v, want: %+v", a.AuthenticateCalls, 1)
	}

	now = time.Date(2015, 12, 24, 11, 0, 0, 0, time.UTC)
	got, err = m.Token()
	if err != nil {
		t.Errorf("Token() failed: %v", err)
	}
	want = OAuthToken{
		"mock oauth token",
		"mock token type",
		"mock refresh token",
		time.Date(2015, 12, 24, 11, 59, 59, 0, time.UTC),
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Token() returned: %+v, want: %+v", got, want)
	}
	if a.AuthenticateCalls != 2 {
		t.Errorf("Called Authenticate() unexpected times: %+v, want: %+v", a.AuthenticateCalls, 2)
	}
}

func TestFailsWhenAuthenticatorFails(t *testing.T) {
	authErr := errors.New("mock auth error")
	a := &mockProsperAuthenticator{
		Err: authErr,
	}
	now := time.Date(2015, 12, 24, 10, 0, 0, 0, time.UTC)
	m := defaultTokenManager{
		authenticator: a,
		clock:         mockClock{&now},
	}

	_, err := m.Token()
	if err != authErr {
		t.Errorf("Token() failed with unexpected error, got: %v, want: %v", err, authErr)
	}
}
