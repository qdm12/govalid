package binary

import (
	"errors"
	"fmt"
	"strings"

	"github.com/qdm12/govalid/internal/helpers"
)

var (
	ErrOption        = errors.New("option error")
	ErrValueNotValid = errors.New("value is not valid")
)

// Validate returns true if the value is one of the enabled values,
// false if it is one of the disabled values, `nil` if value is empty,
// and an error otherwise.
// The enabled values default to 'enabled', 'yes' and 'on'.
// The disabled values default to 'disabled', 'no' and 'off'.
// The string comparison does not consider case sensitivity.
func Validate(value string, options ...Option) (enabled *bool, err error) {
	s := newSettings()
	for _, option := range options {
		err := option(s)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrOption, err)
		}
	}

	for _, enabledString := range s.enabled {
		if strings.EqualFold(value, enabledString) {
			return ptrToBool(true), nil
		}
	}

	for _, disabledString := range s.disabled {
		if strings.EqualFold(value, disabledString) {
			return ptrToBool(false), nil
		}
	}

	if value == "" {
		return nil, nil
	}

	choices := append(s.enabled, s.disabled...)
	return nil, fmt.Errorf("%w: value %q can only be one of %s",
		ErrValueNotValid, value, helpers.CommaJoin(choices))
}

func ptrToBool(b bool) *bool { return &b }
