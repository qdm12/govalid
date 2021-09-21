package rooturl

import (
	"errors"
	"fmt"
	"path"
	"regexp"
	"strings"
)

var rootURLRegex = regexp.MustCompile(`^\/[a-zA-Z0-9\-_/\+]*$`)

var (
	ErrOption        = errors.New("option error")
	ErrValueNotValid = errors.New("value is not valid")
)

// Validate verify the root url path matches an expected regular
// expression and removes any trailing slash at the end of the value.
func Validate(value string, options ...Option) (rootURL string, err error) {
	s := newSettings()
	for _, option := range options {
		err := option(s)
		if err != nil {
			return "", fmt.Errorf("%w: %s", ErrOption, err)
		}
	}

	rootURL = path.Clean(value)

	if !rootURLRegex.MatchString(rootURL) {
		return "", fmt.Errorf("%w: %s", ErrValueNotValid, rootURL)
	}

	for strings.HasSuffix(rootURL, "/") {
		// already have / from paths of router
		rootURL = strings.TrimSuffix(rootURL, "/")
	}

	return rootURL, nil
}
