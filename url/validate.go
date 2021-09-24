package url

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

var (
	ErrOption            = errors.New("option error")
	ErrURLNotValid       = errors.New("url is not valid")
	ErrURLSchemeNotValid = errors.New("url scheme is not valid")
)

// Validate parses the URL from the value given and returns it.
// It verifies the scheme matches 'http' or 'https' by default.
func Validate(value string, options ...Option) (u *url.URL, err error) {
	s := newSettings()
	for _, option := range options {
		err := option(s)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrOption, err)
		}
	}

	u, err = url.Parse(value)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrURLNotValid, err)
	}

	schemeIsValid := false
	for _, accepted := range s.allowedSchemes {
		if strings.EqualFold(u.Scheme, accepted) {
			schemeIsValid = true
			break
		}
	}

	if !schemeIsValid {
		return nil, fmt.Errorf("%w: %s", ErrURLSchemeNotValid, u.String())
	}

	return u, nil
}
