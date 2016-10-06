package prosper

import "github.com/mtlynch/gofn-prosper/prosper/thin"

// notesResponseParser parses Prosper notes into native typed Notes.
type notesResponseParser interface {
	Parse(thin.NotesResponse) (NotesResponse, error)
}

type defaultNotesResponseParser struct {
	np noteParser
}

// newNotesResponseParser creates a new parser for Notes API responses.
func newNotesResponseParser() notesResponseParser {
	return defaultNotesResponseParser{
		np: defaultNoteParser{},
	}
}

// Parse parses a thin.NotesResponse into the richer NotesResponse.
func (p defaultNotesResponseParser) Parse(r thin.NotesResponse) (NotesResponse, error) {
	var notes []Note
	for _, nRaw := range r.Result {
		note, err := p.np.Parse(nRaw)
		if err != nil {
			return NotesResponse{}, err
		}
		notes = append(notes, note)
	}
	return NotesResponse{
		Result:      notes,
		ResultCount: r.ResultCount,
		TotalCount:  r.TotalCount,
	}, nil
}
