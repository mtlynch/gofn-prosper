package types

import (
	"reflect"
	"sort"
	"testing"
)

func TestSortOrderIDs(t *testing.T) {
	var tests = []struct {
		unsorted OrderIDs
		want     OrderIDs
	}{
		{
			unsorted: OrderIDs{"id-a", "id-b", "id-c"},
			want:     OrderIDs{"id-a", "id-b", "id-c"},
		},
		{
			unsorted: OrderIDs{"id-c", "id-b", "id-a"},
			want:     OrderIDs{"id-a", "id-b", "id-c"},
		},
	}
	for _, tt := range tests {
		sort.Sort(tt.unsorted)
		if !reflect.DeepEqual(tt.unsorted, tt.want) {
			t.Errorf("sorting order IDs failed, got: %v, want: %v", tt.unsorted, tt.want)
		}
	}
}
