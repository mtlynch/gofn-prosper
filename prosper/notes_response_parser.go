package prosper

import (
	"github.com/mtlynch/gofn-prosper/prosper/thin"
	"github.com/mtlynch/gofn-prosper/types"
)

// NotesResponseParser parses Prosper notes into native typed Notes.
type NotesResponseParser interface {
	Parse(thin.NotesResponse) (types.NotesResponse, error)
}

type defaultNotesResponseParser struct {
	np noteParser
}

// NewNotesResponseParser creates a new parser for Notes API responses.
func NewNotesResponseParser() NotesResponseParser {
	return defaultNotesResponseParser{
		np: defaultNoteParser{},
	}
}

// Parse parses a thin.NotesResponse into the richer types.NotesResponse.
func (p defaultNotesResponseParser) Parse(r thin.NotesResponse) (types.NotesResponse, error) {
	var notes []types.Note
	for _, nRaw := range r.Result {
		note, err := p.np.Parse(nRaw)
		if err != nil {
			return types.NotesResponse{}, err
		}
		notes = append(notes, note)
	}
	return types.NotesResponse{
		Result:      notes,
		ResultCount: r.ResultCount,
		TotalCount:  r.TotalCount,
	}, nil
}
