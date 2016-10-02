// Package prosper is a set of APIs that wrap the raw HTTP Prosper REST APIs.
package prosper

import (
	"github.com/mtlynch/gofn-prosper/prosper/auth"
	"github.com/mtlynch/gofn-prosper/prosper/thin"
	"github.com/mtlynch/gofn-prosper/types"
)

// Client is a Prosper client that communicates with the Prosper HTTP endpoints.
type Client struct {
	rawClient     thin.RawApiHandler
	ap            accountsParser
	nrp           notesResponseParser
	listingParser listingParser
	orderParser   orderParser
}

// NewClient creates a new Client with the given Prosper credentials.
func NewClient(creds types.ClientCredentials) Client {
	tokenMgr := auth.NewTokenManager(auth.NewAuthenticator(creds))
	return Client{
		rawClient:     thin.NewClient(tokenMgr),
		ap:            defaultAccountsParser{},
		nrp:           NewNotesResponseParser(),
		listingParser: defaultListingParser{},
		orderParser:   defaultOrderParser{},
	}
}
