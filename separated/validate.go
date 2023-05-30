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
	s, err := newSettings(options...)
	if err != nil {
		return nil, err
	}

	values = strings.Split(value, s.separator)
	if s.ignoreEmpty {
		values = removeEmpty(values)
	}

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

	var invalidValues []valuePosition

	for i, value := range values {
		if _, ok := acceptedSet[value]; !ok {
			invalidValues = append(invalidValues, valuePosition{i + 1, value})
		}
	}

	if len(invalidValues) == 0 {
		return values, nil
	}

	return nil, makeInvalidError(s.accepted, invalidValues)
}

func removeEmpty(values []string) (nonEmptyValues []string) {
	i := 0
	for _, value := range values {
		if value == "" {
			continue
		}
		values[i] = value
		i++
	}

	if i == 0 {
		return nil
	}

	values = values[:i]
	return values
}

type valuePosition struct {
	position int
	value    string
}

func makeInvalidError(accepted []string, invalidValues []valuePosition) (err error) {
	acceptedValues := strings.Join(accepted, ", ")

	invalidMessages := make([]string, len(invalidValues))
	for i := range invalidValues {
		invalidMessages[i] = fmt.Sprintf("value %q at position %d",
			invalidValues[i].value, invalidValues[i].position)
	}
	invalidMessage := strings.Join(invalidMessages, ", ")

	return fmt.Errorf("%w: %s; accepted values are: %s",
		ErrValueNotValid, invalidMessage, acceptedValues)
}
