package prosper

import (
	"errors"

	"github.com/mtlynch/gofn-prosper/prosper/thin"
)

type mockRawClient struct {
	accountsResponse thin.AccountsResponse
	notesResponse    thin.NotesResponse
	orderResponse    thin.OrderResponse
	searchParams     thin.SearchParams
	searchResponse   thin.SearchResponse
	err              error
}

var mockRawClientErr = errors.New("mock raw client error")
