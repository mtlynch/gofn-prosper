package prosper

import (
	"github.com/mtlynch/gofn-prosper/prosper/thin"
	"github.com/mtlynch/gofn-prosper/types"
)

type notesResponseParser interface {
	Parse(thin.NotesResponse) (types.NotesResponse, error)
}

type defaultNotesResponseParser struct {
	np noteParser
}

func NewNotesResponseParser() notesResponseParser {
	return defaultNotesResponseParser{
		np: defaultNoteParser{},
	}
}

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
