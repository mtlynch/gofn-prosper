package interval

import "time"

// CreateInt32 creates a pointer to an int32 type.
func CreateInt32(x int32) *int32 {
	return &x
}

// CreateFloat64 creates a pointer to a float64 type.
func CreateFloat64(x float64) *float64 {
	return &x
}

// CreateTime creates a pointer to a time.Time type.
func CreateTime(x time.Time) *time.Time {
	return &x
}
