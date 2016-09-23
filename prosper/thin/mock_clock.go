package thin

import "time"

type mockClock struct {
	now time.Time
}

func (c mockClock) Now() time.Time {
	return c.now
}
