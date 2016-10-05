package prosper

import (
	"github.com/mtlynch/gofn-prosper/prosper/thin"
	"github.com/mtlynch/gofn-prosper/types"
)

// NoteFetcher supports the Notes API for retrieving the user's notes.
type NoteFetcher interface {
	Notes(offset, limit int) (types.NotesResponse, error)
}

// Notes returns a subset of the notes that the user owns. Notes partially
// implements the REST API described at:
// https://developers.prosper.com/docs/investor/notes-api/
func (c Client) Notes(offset, limit int) (types.NotesResponse, error) {
	notesResponseRaw, err := c.rawClient.Notes(thin.NotesParams{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return types.NotesResponse{}, err
	}
	return c.nrp.Parse(notesResponseRaw)
}
