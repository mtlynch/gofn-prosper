package prosper

import (
	"testing"
	"time"
)

func TestParseProsperOldTime(t *testing.T) {
	var tests = []struct {
		input         string
		want          time.Time
		expectSuccess bool
		msg           string
	}{
		{
			input:         "03221991",
			want:          time.Date(1991, 3, 22, 0, 0, 0, 0, time.UTC),
			expectSuccess: true,
			msg:           "normal time should parse normally",
		},
		{
			input:         "27051991",
			expectSuccess: false,
			msg:           "date with invalid month should fail",
		},
		{
			input:         "1",
			want:          time.Time{},
			expectSuccess: true,
			msg:           "'1' value should be treated as nil time",
		},
		{
			input:         "2",
			want:          time.Time{},
			expectSuccess: true,
			msg:           "'2' value should be treated as nil time",
		},
	}
	for _, tt := range tests {
		got, err := parseProsperOldTime(tt.input)
		if tt.expectSuccess && err != nil {
			t.Errorf("%s - expected successful parsing of %+v, got error: %v", tt.msg, tt.input, err)
		} else if !tt.expectSuccess && err == nil {
			t.Errorf("%s - expected failure for %+v, got nil", tt.msg, tt.input)
		}
		if tt.expectSuccess && !got.Equal(tt.want) {
			t.Errorf("%s - parseProsperOldTime got: %#v, want: %#v", tt.msg, got, tt.want)
		}
	}
}
