package duration

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrOption            = errors.New("option error")
	ErrDurationMalformed = errors.New("duration is malformed")
	ErrDurationNegative  = errors.New("duration cannot be negative")
	ErrDurationZero      = errors.New("duration cannot be zero")
)

// Validate parses the duration from the value and verify it matches
// the options given.
func Validate(value string, options ...Option) (duration time.Duration, err error) {
	s := newSettings()
	for _, option := range options {
		err := option(s)
		if err != nil {
			return 0, fmt.Errorf("%w: %w", ErrOption, err)
		}
	}

	duration, err = time.ParseDuration(value)
	if err != nil {
		return 0, fmt.Errorf("%w: %s: %w", ErrDurationMalformed, value, err)
	}

	if !s.allowNegative && duration < 0 {
		return 0, fmt.Errorf("%w: %s", ErrDurationNegative, duration.String())
	}

	if !s.allowZero && duration == 0 {
		return 0, fmt.Errorf("%w: %s", ErrDurationZero, value)
	}

	return duration, nil
}
