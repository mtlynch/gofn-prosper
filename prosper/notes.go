package prosper

import "github.com/mtlynch/gofn-prosper/types"

type NoteFetcher interface {
	Notes(offset, limit int) (types.NotesResponse, error)
}

func (c Client) Notes(offset, limit int) (types.NotesResponse, error) {
	notesResponseRaw, err := c.rawClient.Notes(offset, limit)
	if err != nil {
		return types.NotesResponse{}, err
	}
	return c.nrp.Parse(notesResponseRaw)
}
