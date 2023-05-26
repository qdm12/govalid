package integer

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	ErrNotAnInteger    = errors.New("value is not an integer")
	ErrIntegerTooSmall = errors.New("value is too small")
	ErrIntegerTooBig   = errors.New("value is too big")
	ErrOption          = errors.New("option error")
)

// Validate returns the parsed integer from the value.
// It does additional checks depending on the options given.
func Validate(value string, options ...Option) (integer int, err error) {
	s := newSettings()
	for _, option := range options {
		err := option(s)
		if err != nil {
			return 0, fmt.Errorf("%w: %w", ErrOption, err)
		}
	}

	integer, err = strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%w: %s", ErrNotAnInteger, value)
	}

	if s.minimum != nil && integer < *s.minimum {
		return 0, fmt.Errorf("%w: %d is smaller than minimum of %d",
			ErrIntegerTooSmall, integer, *s.minimum)
	}

	if s.maximum != nil && integer > *s.maximum {
		return 0, fmt.Errorf("%w: %d is bigger than maximum of %d",
			ErrIntegerTooBig, integer, *s.maximum)
	}

	return integer, nil
}
