package address

import (
	"errors"
	"fmt"
	"net"

	"github.com/qdm12/govalid/port"
)

var (
	ErrOption        = errors.New("option error")
	ErrValueNotValid = errors.New("value is not valid")
	ErrInvalidPort   = errors.New("invalid port")
)

// Validate validates the value is a valid address.
// It does extra checks depending on the options given.
func Validate(value string, options ...Option) (err error) {
	s := settings{}
	for _, option := range options {
		err := option(&s)
		if err != nil {
			return fmt.Errorf("%w: %w", ErrOption, err)
		}
	}

	_, portStr, err := net.SplitHostPort(value)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrValueNotValid, err)
	}

	_, err = port.Validate(portStr, s.portOptions...)
	if err != nil {
		return err
	}

	return nil
}
