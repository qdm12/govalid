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

// Validate returns true if the value is one of the enabled values
// and false if it is one of the disabled values.
// It returns an error if it is not recognized.
// The enabled values default to 'enabled', 'yes' and 'on'.
// The disabled values default to 'disabled', 'no' and 'off'.
// The string comparison does not consider case sensitivity.
func Validate(value string, options ...Option) (enabled bool, err error) {
	s := newSettings()
	for _, option := range options {
		err := option(s)
		if err != nil {
			return false, fmt.Errorf("%w: %s", ErrOption, err)
		}
	}

	for _, enabledString := range s.enabled {
		if strings.EqualFold(value, enabledString) {
			return true, nil
		}
	}

	for _, disabledString := range s.disabled {
		if strings.EqualFold(value, disabledString) {
			return false, nil
		}
	}

	choices := append(s.enabled, s.disabled...)
	return false, fmt.Errorf("%w: value %q can only be one of %s",
		ErrValueNotValid, value, helpers.CommaJoin(choices))
}
