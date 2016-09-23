package prosper

import (
	"github.com/mtlynch/gofn-prosper/prosper/thin"
)

type Client struct {
	rawClient     thin.RawApiHandler
	ap            accountsParser
	nrp           notesResponseParser
	listingParser listingParser
	orderParser   orderParser
}
