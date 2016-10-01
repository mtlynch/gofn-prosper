package types

import (
	"testing"
)

func TestNewFloat64Range(t *testing.T) {
	r := NewFloat64Range(-1.0, 5.6)
	if *r.Min != -1.0 {
		t.Errorf("NewFloat64Range failed to set Min properly. got %v, want %v", r.Min, -1.0)
	}
	if *r.Max != 5.6 {
		t.Errorf("NewFloat64Range failed to set Max properly. got %v, want %v", r.Max, 5.6)
	}
}

func TestNewInt32Range(t *testing.T) {
	r := NewInt32Range(42, -9)
	if *r.Min != 42 {
		t.Errorf("NewInt32Range failed to set Min properly. got %v, want %v", r.Min, 42)
	}
	if *r.Max != -9 {
		t.Errorf("NewInt32Range failed to set Max properly. got %v, want %v", r.Max, -9)
	}
}

func TestFloat64RangeEqual(t *testing.T) {
	var tests = []struct {
		a    Float64Range
		b    Float64Range
		want bool
	}{
		{
			a:    Float64Range{Min: nil, Max: nil},
			b:    Float64Range{Min: nil, Max: nil},
			want: true,
		},
		{
			a:    Float64Range{Min: CreateFloat64(5.5), Max: nil},
			b:    Float64Range{Min: nil, Max: nil},
			want: false,
		},
		{
			a:    Float64Range{Min: CreateFloat64(5.5), Max: CreateFloat64(7.7)},
			b:    Float64Range{Min: nil, Max: nil},
			want: false,
		},
		{
			a:    Float64Range{Min: CreateFloat64(5.5), Max: nil},
			b:    Float64Range{Min: CreateFloat64(5.5), Max: nil},
			want: true,
		},
		{
			a:    Float64Range{Min: CreateFloat64(5.5), Max: CreateFloat64(7.7)},
			b:    Float64Range{Min: CreateFloat64(5.5), Max: nil},
			want: false,
		},
		{
			a:    Float64Range{Min: CreateFloat64(5.5), Max: CreateFloat64(7.7)},
			b:    Float64Range{Min: CreateFloat64(5.5), Max: CreateFloat64(7.7)},
			want: true,
		},
		{
			a:    Float64Range{Min: CreateFloat64(5.5), Max: CreateFloat64(7.7)},
			b:    Float64Range{Min: CreateFloat64(58.58), Max: CreateFloat64(7.7)},
			want: false,
		},
		{
			a:    Float64Range{Min: CreateFloat64(5.5), Max: CreateFloat64(7.7)},
			b:    Float64Range{Min: CreateFloat64(5.5), Max: CreateFloat64(58.58)},
			want: false,
		},
	}
	for _, tt := range tests {
		got := Float64RangeEqual(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("Comparing equality of %s and %s: got %v, want: %v", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestInt32RangeEqual(t *testing.T) {
	var tests = []struct {
		a    Int32Range
		b    Int32Range
		want bool
	}{
		{
			a:    Int32Range{Min: nil, Max: nil},
			b:    Int32Range{Min: nil, Max: nil},
			want: true,
		},
		{
			a:    Int32Range{Min: CreateInt32(5), Max: nil},
			b:    Int32Range{Min: nil, Max: nil},
			want: false,
		},
		{
			a:    Int32Range{Min: CreateInt32(5), Max: CreateInt32(7)},
			b:    Int32Range{Min: nil, Max: nil},
			want: false,
		},
		{
			a:    Int32Range{Min: CreateInt32(5), Max: nil},
			b:    Int32Range{Min: CreateInt32(5), Max: nil},
			want: true,
		},
		{
			a:    Int32Range{Min: CreateInt32(5), Max: CreateInt32(7)},
			b:    Int32Range{Min: CreateInt32(5), Max: nil},
			want: false,
		},
		{
			a:    Int32Range{Min: CreateInt32(5), Max: CreateInt32(7)},
			b:    Int32Range{Min: CreateInt32(5), Max: CreateInt32(7)},
			want: true,
		},
		{
			a:    Int32Range{Min: CreateInt32(5), Max: CreateInt32(7)},
			b:    Int32Range{Min: CreateInt32(58), Max: CreateInt32(7)},
			want: false,
		},
		{
			a:    Int32Range{Min: CreateInt32(5), Max: CreateInt32(7)},
			b:    Int32Range{Min: CreateInt32(5), Max: CreateInt32(58)},
			want: false,
		},
	}
	for _, tt := range tests {
		got := Int32RangeEqual(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("Comparing equality of %s and %s: got %v, want: %v", tt.a, tt.b, got, tt.want)
		}
	}
}
