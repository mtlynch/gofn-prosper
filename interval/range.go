// Package interval offers types and functions for representing ranges of data
// with a minimum and/or maximum value.
package interval

import (
	"fmt"
	"time"
)

type (
	// Float64Range represents a range of numbers of type float64.
	Float64Range struct {
		Min *float64
		Max *float64
	}

	// Int32Range represents a range of numbers of type int32.
	Int32Range struct {
		Min *int32
		Max *int32
	}

	// TimeRange represents a range of time.
	TimeRange struct {
		Min *time.Time
		Max *time.Time
	}
)

// NewFloat64Range creates a new Float64Range object.
func NewFloat64Range(min, max float64) Float64Range {
	return Float64Range{&min, &max}
}

// NewInt32Range creates a new Int32Range object.
func NewInt32Range(min, max int32) Int32Range {
	return Int32Range{&min, &max}
}

// Float64RangeEqual returns true if two ranges are equal.
func Float64RangeEqual(a, b Float64Range) bool {
	if (a.Min == nil && b.Min != nil) ||
		(a.Min != nil && b.Min == nil) ||
		(a.Max == nil && b.Max != nil) ||
		(a.Max != nil && b.Max == nil) {
		return false
	}
	if (a.Min != nil) && (b.Min != nil) &&
		(*a.Min != *b.Min) {
		return false
	}
	if (a.Max != nil) && (b.Max != nil) &&
		(*a.Max != *b.Max) {
		return false
	}
	return true
}

// String returns a string representation of a Float64Range.
func (r Float64Range) String() string {
	var min, max string
	if r.Min != nil {
		min = fmt.Sprintf("%f", *r.Min)
	} else {
		min = "nil"
	}
	if r.Max != nil {
		max = fmt.Sprintf("%f", *r.Max)
	} else {
		max = "nil"
	}
	return fmt.Sprintf("{Min: %s, Max: %s}", min, max)
}

// Int32RangeEqual returns true if two ranges are equal.
func Int32RangeEqual(a, b Int32Range) bool {
	if (a.Min == nil && b.Min != nil) ||
		(a.Min != nil && b.Min == nil) ||
		(a.Max == nil && b.Max != nil) ||
		(a.Max != nil && b.Max == nil) {
		return false
	}
	if (a.Min != nil) && (b.Min != nil) &&
		(*a.Min != *b.Min) {
		return false
	}
	if (a.Max != nil) && (b.Max != nil) &&
		(*a.Max != *b.Max) {
		return false
	}
	return true
}

// String returns a string representation of an Int32Range.
func (r Int32Range) String() string {
	var min, max string
	if r.Min != nil {
		min = fmt.Sprintf("%d", *r.Min)
	} else {
		min = "nil"
	}
	if r.Max != nil {
		max = fmt.Sprintf("%d", *r.Max)
	} else {
		max = "nil"
	}
	return fmt.Sprintf("{Min: %s, Max: %s}", min, max)
}
