package types

import "time"

func CreateInt32(x int32) *int32 {
	return &x
}

func CreateFloat64(x float64) *float64 {
	return &x
}

func CreateTime(x time.Time) *time.Time {
	return &x
}
