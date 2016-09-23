package prosper

import (
	"errors"
	"reflect"
	"testing"

	"github.com/mtlynch/gofn-prosper/prosper/thin"
	"github.com/mtlynch/gofn-prosper/types"
)

type mockParseResult struct {
	parsed types.Note
	err    error
}

type mockNoteParser struct {
	got     []thin.NoteResult
	returns []mockParseResult
}

func (p *mockNoteParser) Parse(r thin.NoteResult) (types.Note, error) {
	p.got = append(p.got, r)
	var result mockParseResult
	result, p.returns = p.returns[0], p.returns[1:]
	return result.parsed, result.err
}

func TestNotesResponseParser(t *testing.T) {
	var tests = []struct {
		input         thin.NotesResponse
		parseResults  []mockParseResult
		want          types.NotesResponse
		expectSuccess bool
		msg           string
	}{
		{
			input: thin.NotesResponse{
				Result: []thin.NoteResult{
					{LoanNumber: 123},
				},
				ResultCount: 1,
				TotalCount:  1,
			},
			parseResults: []mockParseResult{
				{
					parsed: types.Note{LoanNumber: 123},
				},
			},
			want: types.NotesResponse{
				Result: []types.Note{
					{LoanNumber: 123},
				},
				ResultCount: 1,
				TotalCount:  1,
			},
			expectSuccess: true,
			msg:           "valid note response should parse successfully",
		},
		{
			input: thin.NotesResponse{
				Result: []thin.NoteResult{
					{LoanNumber: 123},
					{LoanNumber: 124},
					{LoanNumber: 125},
				},
				ResultCount: 3,
				TotalCount:  3,
			},
			parseResults: []mockParseResult{
				{
					parsed: types.Note{LoanNumber: 123},
				},
				{
					parsed: types.Note{LoanNumber: 124},
				},
				{
					parsed: types.Note{LoanNumber: 125},
				},
			},
			want: types.NotesResponse{
				Result: []types.Note{
					{LoanNumber: 123},
					{LoanNumber: 124},
					{LoanNumber: 125},
				},
				ResultCount: 3,
				TotalCount:  3,
			},
			expectSuccess: true,
			msg:           "valid note response should parse successfully",
		},
		{
			input: thin.NotesResponse{
				Result: []thin.NoteResult{
					{LoanNumber: 123},
					{LoanNumber: 124},
					{LoanNumber: 125},
				},
				ResultCount: 3,
				TotalCount:  3,
			},
			parseResults: []mockParseResult{
				{
					parsed: types.Note{LoanNumber: 123},
				},
				{
					err: errors.New("mock note parsing error"),
				},
				{
					parsed: types.Note{LoanNumber: 125},
				},
			},
			expectSuccess: false,
			msg:           "parsing should fail if a note can't be parsed",
		},
	}
	for _, tt := range tests {
		noteParser := mockNoteParser{returns: tt.parseResults}
		got, err := defaultNotesResponseParser{
			np: &noteParser,
		}.Parse(tt.input)
		if tt.expectSuccess && err != nil {
			t.Errorf("%s - expected successful parsing of %+v, got error: %v", tt.msg, tt.input, err)
		} else if !tt.expectSuccess && err == nil {
			t.Errorf("%s - expected failure for %+v, got nil", tt.msg, tt.input)
		}
		if tt.expectSuccess && !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%s - defaultNoteResponseParser.Parse returned %#v, want %#v", tt.msg, got, tt.want)
		}
		if tt.expectSuccess && !reflect.DeepEqual(noteParser.got, tt.input.Result) {
			t.Errorf("%s - noteParser got %+v, want %+v", tt.msg, noteParser.got, tt.input.Result)
		}
	}
}
