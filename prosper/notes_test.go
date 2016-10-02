package prosper

import (
	"errors"
	"reflect"
	"testing"

	"github.com/mtlynch/gofn-prosper/prosper/thin"
	"github.com/mtlynch/gofn-prosper/types"
)

var (
	gotOffset             int
	gotLimit              int
	mockRawNotesResponseA = thin.NotesResponse{TotalCount: 25}
	mockRawNotesResponseB = thin.NotesResponse{TotalCount: 50}
	mockNotesResponseA    = types.NotesResponse{TotalCount: 25}
	mockNotesResponseB    = types.NotesResponse{TotalCount: 50}
	errMockParserFail     = errors.New("mock parser error")
)

func (c *mockRawClient) Notes(offset, limit int) (thin.NotesResponse, error) {
	gotOffset = offset
	gotLimit = limit
	return c.notesResponse, c.err
}

type mockNotesResponseParser struct {
	gotNotesResponse thin.NotesResponse
	notesResponse    types.NotesResponse
	err              error
}

func (p *mockNotesResponseParser) Parse(r thin.NotesResponse) (types.NotesResponse, error) {
	p.gotNotesResponse = r
	return p.notesResponse, p.err
}

func TestNotes(t *testing.T) {
	tests := []struct {
		offset       int
		limit        int
		rawResponse  thin.NotesResponse
		clientErr    error
		parserResult types.NotesResponse
		parserErr    error
		want         types.NotesResponse
		wantErr      error
	}{
		{
			offset:    0,
			limit:     25,
			clientErr: errMockRawClientFail,
			wantErr:   errMockRawClientFail,
		},
		{
			offset:    0,
			limit:     25,
			parserErr: errMockParserFail,
			wantErr:   errMockParserFail,
		},
		{
			offset:       0,
			limit:        25,
			rawResponse:  mockRawNotesResponseA,
			parserResult: mockNotesResponseA,
			want:         mockNotesResponseA,
		},
		{
			offset:       25,
			limit:        75,
			rawResponse:  mockRawNotesResponseB,
			parserResult: mockNotesResponseB,
			want:         mockNotesResponseB,
		},
	}
	for _, tt := range tests {
		parser := mockNotesResponseParser{
			notesResponse: tt.parserResult,
			err:           tt.parserErr,
		}
		client := Client{
			rawClient: &mockRawClient{
				notesResponse: tt.rawResponse,
				err:           tt.clientErr,
			},
			nrp: &parser,
		}
		got, err := client.Notes(tt.offset, tt.limit)
		if err != tt.wantErr {
			t.Errorf("unexpected failure from client.Notes. got: %v, want: %v", err, tt.wantErr)
		}
		if gotOffset != tt.offset {
			t.Errorf("unexpected offset passed to raw client. got: %v, want: %v", gotOffset, tt.offset)
		}
		if gotLimit != tt.limit {
			t.Errorf("unexpected limit passed to raw client. got: %v, want: %v", gotLimit, tt.limit)
		}
		if tt.wantErr != nil {
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unexpected result from client.Notes. got %#v, want %#v", got, tt.want)
			}
			if !reflect.DeepEqual(parser.gotNotesResponse, tt.rawResponse) {
				t.Errorf("parser got: %v, want %v", parser.gotNotesResponse, tt.rawResponse)
			}
		}
	}
}
