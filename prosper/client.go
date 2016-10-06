// Package prosper is a set of APIs that wrap the raw HTTP Prosper REST APIs.
package prosper

import (
	"github.com/mtlynch/gofn-prosper/prosper/auth"
	"github.com/mtlynch/gofn-prosper/prosper/thin"
)

// Client is a Prosper client that communicates with the Prosper HTTP endpoints.
type Client struct {
	rawClient           thin.RawAPIHandler
	accountParser       accountParser
	notesResponseParser notesResponseParser
	listingParser       listingParser
	orderParser         orderParser
}

// NewClient creates a new Client with the given Prosper credentials.
func NewClient(creds auth.ClientCredentials) Client {
	tokenMgr := auth.NewTokenManager(auth.NewAuthenticator(creds))
	return Client{
		rawClient:           thin.NewClient(tokenMgr),
		accountParser:       defaultAccountParser{},
		notesResponseParser: newNotesResponseParser(),
		listingParser:       defaultListingParser{},
		orderParser:         defaultOrderParser{},
	}
}
