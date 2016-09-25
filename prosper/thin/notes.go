package thin

import "fmt"

type (
	// NotesResponse contains the full response from the Notes API in minimally
	// parsed form.
	NotesResponse struct {
		Result      []NoteResult `json:"result"`
		ResultCount int          `json:"result_count"`
		TotalCount  int          `json:"total_count"`
	}
)

// Notes returns a subset of the notes that the user owns. Notes partially
// implements the REST API described at:
// https://developers.prosper.com/docs/investor/notes-api/
func (c Client) Notes(offset, limit int) (response NotesResponse, err error) {
	url := fmt.Sprintf("%s/notes/?offset=%d&limit=%d", c.baseUrl, offset, limit)
	err = c.DoRequest("GET", url, nil, &response)
	if err != nil {
		return NotesResponse{}, err
	}
	return response, nil
}
