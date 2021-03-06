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
func Validate(value string, options ...Option) (
	address string, err error) {
	s := settings{}
	for _, option := range options {
		err := option(&s)
		if err != nil {
			return "", fmt.Errorf("%w: %s", ErrOption, err)
		}
	}

	host, portStr, err := net.SplitHostPort(value)
	if err != nil {
		return "", fmt.Errorf("%w: %s", ErrValueNotValid, err)
	}

	_, err = port.Validate(portStr, s.portOptions...)
	if err != nil {
		return "", err
	}

	address = host + ":" + portStr

	return address, nil
}
