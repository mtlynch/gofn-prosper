package prosper

import (
	"github.com/mtlynch/gofn-prosper/prosper/auth"
	"github.com/mtlynch/gofn-prosper/prosper/thin"
	"github.com/mtlynch/gofn-prosper/types"
)

type Client struct {
	rawClient     thin.RawApiHandler
	ap            accountsParser
	nrp           notesResponseParser
	listingParser listingParser
	orderParser   orderParser
}

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
