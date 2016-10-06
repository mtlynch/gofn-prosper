package auth

// ClientCredentials represents the user's secret credentials used to
// authenticate to the Prosper API.
type ClientCredentials struct {
	ClientID     string
	ClientSecret string
	Username     string
	Password     string
}
