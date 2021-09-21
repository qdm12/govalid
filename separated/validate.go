package separated

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrOption        = errors.New("option error")
	ErrValueNotValid = errors.New("value is not valid")
)

// Validate returns a slice of strings from the
// value given using the separator set, which defaults to ",".
// Accepted values can be set with the AcceptedValues option
// and it defaults to accept all values.
// The string comparisons do not consider case sensitivity.
func Validate(value string, options ...Option) (
	values []string, err error) {
	s := newSettings()
	for _, option := range options {
		err := option(s)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrOption, err)
		}
	}

	values = strings.Split(value, s.separator)

	if s.lowercase {
		for i := range values {
			values[i] = strings.ToLower(values[i])
		}
	}

	if len(s.accepted) == 0 {
		return values, nil
	}

	// Check values are all accepted
	acceptedSet := make(map[string]struct{}, len(s.accepted))
	for _, accepted := range s.accepted {
		acceptedSet[accepted] = struct{}{}
	}

	type valuePosition struct {
		position int
		value    string
	}
	var invalidValues []valuePosition

	for i, value := range values {
		if !isValueAccepted(value, acceptedSet) {
			invalidValues = append(invalidValues, valuePosition{i + 1, value})
		}
	}

	if len(invalidValues) == 0 {
		return values, nil
	}

	acceptedValues := strings.Join(s.accepted, ", ")

	invalidMessages := make([]string, len(invalidValues))
	for i := range invalidValues {
		invalidMessages[i] = fmt.Sprintf("value %q at position %d",
			invalidValues[i].value, invalidValues[i].position)
	}
	invalidMessage := strings.Join(invalidMessages, ", ")

	return nil, fmt.Errorf("%w: %s; accepted values are: %s",
		ErrValueNotValid, invalidMessage, acceptedValues)
}

func isValueAccepted(value string, acceptedSet map[string]struct{}) (ok bool) {
	if len(acceptedSet) == 0 {
		return true
	}
	_, ok = acceptedSet[value]
	return ok
}
