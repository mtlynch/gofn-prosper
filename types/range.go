package types

import (
	"fmt"
	"time"
)

type (
	Float64Range struct {
		Min *float64
		Max *float64
	}

	Int32Range struct {
		Min *int32
		Max *int32
	}

	TimeRange struct {
		Min *time.Time
		Max *time.Time
	}
)

func NewFloat64Range(min, max float64) Float64Range {
	return Float64Range{&min, &max}
}

func NewInt32Range(min, max int32) Int32Range {
	return Int32Range{&min, &max}
}

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
