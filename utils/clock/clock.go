// Package clock provides an abstract layer over the standard time package
package clock

import (
	"time"
)

//TODO: use mutex to avoid race

// Clock is an interface to the standard library time.
// It is used to implement a real or a mock clock. The latter is used in tests.
type Clock interface {
	Now() time.Time
}

type clock struct{}

func (c *clock) Now() time.Time {
	return time.Now()
}

// Mock is a mock instance of clock
type Mock struct {
	currentTime time.Time
}

// SetNow sets the current time for the mock clock
func (c *Mock) SetNow(t time.Time) {
	c.currentTime = t
}

// Now returns the current time
func (c *Mock) Now() time.Time {
	return c.currentTime
}

// New returns an instance of a real clock
func New() Clock {
	return &clock{}
}

// NewMock returns an instance of a mock clock
func NewMock() *Mock {
	return &Mock{
		currentTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
	}
}
