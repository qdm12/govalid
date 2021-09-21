package integer

import (
	"errors"
	"fmt"
)

type Option func(s *settings) (err error)

var (
	errMinimumAlreadySet = errors.New("minimum is already set")
	errMaximumAlreadySet = errors.New("maximum is already set")
)

// OptionMinimum sets a minimum integer value.
func OptionMinimum(minimum int) Option {
	return func(s *settings) (err error) {
		if s.minimum != nil {
			return fmt.Errorf("%w: %d", errMinimumAlreadySet, *s.minimum)
		}
		s.minimum = &minimum
		return nil
	}
}

// OptionMaximum sets a maximum integer value.
func OptionMaximum(maximum int) Option {
	return func(s *settings) (err error) {
		if s.maximum != nil {
			return fmt.Errorf("%w: %d", errMaximumAlreadySet, *s.maximum)
		}
		s.maximum = &maximum
		return nil
	}
}

// OptionRange sets a range for the integer value.
func OptionRange(minimum, maximum int) Option {
	return func(s *settings) (err error) {
		if s.minimum != nil {
			return fmt.Errorf("%w: %d", errMinimumAlreadySet, *s.minimum)
		} else if s.maximum != nil {
			return fmt.Errorf("%w: %d", errMaximumAlreadySet, *s.maximum)
		}
		s.minimum = &minimum
		s.maximum = &maximum
		return nil
	}
}
