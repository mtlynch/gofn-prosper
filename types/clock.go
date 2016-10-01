package types

import (
	"time"
)

// Clock is an object that gives the current time.
type Clock interface {
	Now() time.Time
}

// DefaultClock is a default clock implementation that returns the system time.
type DefaultClock struct {
}

// Now returns the current time.
func (c DefaultClock) Now() time.Time {
	return time.Now()
}
