package thin

import "fmt"

type (
	NotesResponse struct {
		Result      []NoteResult `json:"result"`
		ResultCount int          `json:"result_count"`
		TotalCount  int          `json:"total_count"`
	}
)

func (c Client) Notes(offset, limit int) (response NotesResponse, err error) {
	url := fmt.Sprintf("%s/notes/?offset=%d&limit=%d", c.baseUrl, offset, limit)
	err = c.DoRequest("GET", url, nil, &response)
	if err != nil {
		return NotesResponse{}, err
	}
	return response, nil
}
