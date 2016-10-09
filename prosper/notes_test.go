package prosper

import (
	"errors"
	"reflect"
	"testing"

	"github.com/mtlynch/gofn-prosper/prosper/thin"
)

var (
	gotNotesParams        thin.NotesParams
	mockRawNotesResponseA = thin.NotesResponse{TotalCount: 25}
	mockRawNotesResponseB = thin.NotesResponse{TotalCount: 50}
	mockNotesResponseA    = NotesResponse{TotalCount: 25}
	mockNotesResponseB    = NotesResponse{TotalCount: 50}
	errMockParserFail     = errors.New("mock parser error")
)

func (c *mockRawClient) Notes(p thin.NotesParams) (thin.NotesResponse, error) {
	gotNotesParams = p
	return c.notesResponse, c.err
}

type mockNotesResponseParser struct {
	gotNotesResponse thin.NotesResponse
	notesResponse    NotesResponse
	err              error
}

func (p *mockNotesResponseParser) Parse(r thin.NotesResponse) (NotesResponse, error) {
	p.gotNotesResponse = r
	return p.notesResponse, p.err
}

func TestNotes(t *testing.T) {
	tests := []struct {
		params       NotesParams
		rawResponse  thin.NotesResponse
		clientErr    error
		parserResult NotesResponse
		parserErr    error
		want         NotesResponse
		wantParams   thin.NotesParams
		wantErr      error
	}{
		{
			params: NotesParams{
				Offset: 0,
				Limit:  25,
			},
			wantParams: thin.NotesParams{
				Offset: 0,
				Limit:  25,
			},
			clientErr: errMockRawClientFail,
			wantErr:   errMockRawClientFail,
		},
		{
			params: NotesParams{
				Offset: 0,
				Limit:  25,
			},
			wantParams: thin.NotesParams{
				Offset: 0,
				Limit:  25,
			},
			parserErr: errMockParserFail,
			wantErr:   errMockParserFail,
		},
		{
			params: NotesParams{
				Offset: 0,
				Limit:  25,
			},
			wantParams: thin.NotesParams{
				Offset: 0,
				Limit:  25,
			},
			rawResponse:  mockRawNotesResponseA,
			parserResult: mockNotesResponseA,
			want:         mockNotesResponseA,
		},
		{
			params: NotesParams{
				Offset: 25,
				Limit:  75,
			},
			wantParams: thin.NotesParams{
				Offset: 25,
				Limit:  75,
			},
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
		client := defaultClient{
			rawClient: &mockRawClient{
				notesResponse: tt.rawResponse,
				err:           tt.clientErr,
			},
			notesResponseParser: &parser,
		}
		got, err := client.Notes(tt.params)
		if err != tt.wantErr {
			t.Errorf("unexpected failure from client.Notes. got: %v, want: %v", err, tt.wantErr)
		}
		if !reflect.DeepEqual(gotNotesParams, tt.wantParams) {
			t.Errorf("unexpected params passed to raw client. got: %v, want: %v", gotNotesParams, tt.params)
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
