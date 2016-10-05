package prosper

import (
	"github.com/mtlynch/gofn-prosper/prosper/thin"
	"github.com/mtlynch/gofn-prosper/types"
)

// NotesParams contains the parameters to the Notes API.
type NotesParams struct {
	Offset int
	Limit  int
	// TODO(mtlynch): Implement support for the sort_by parameter.
}

// NoteFetcher supports the Notes API for retrieving the user's notes.
type NoteFetcher interface {
	Notes(NotesParams) (types.NotesResponse, error)
}

// Notes returns a subset of the notes that the user owns. Notes partially
// implements the REST API described at:
// https://developers.prosper.com/docs/investor/notes-api/
func (c Client) Notes(p NotesParams) (types.NotesResponse, error) {
	notesResponseRaw, err := c.rawClient.Notes(notesParamsToThinType(p))
	if err != nil {
		return types.NotesResponse{}, err
	}
	return c.nrp.Parse(notesResponseRaw)
}

func notesParamsToThinType(p NotesParams) thin.NotesParams {
	return thin.NotesParams{
		Offset: p.Offset,
		Limit:  p.Limit,
	}
}
