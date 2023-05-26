package rooturl

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var rootURLRegex = regexp.MustCompile(`^(|\/[a-zA-Z0-9\-_/\+]*)$`)

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

	rootURL = strings.TrimSpace(value)

	// Clean path and remove trailing slash(es)
	// we already have / from paths of router
	for strings.HasSuffix(rootURL, "/") {
		rootURL = strings.TrimSuffix(rootURL, "/")
	}
	for strings.Contains(rootURL, "//") {
		rootURL = strings.ReplaceAll(rootURL, "//", "/")
	}

	if !rootURLRegex.MatchString(rootURL) {
		return "", fmt.Errorf("%w: %s", ErrValueNotValid, rootURL)
	}

	return rootURL, nil
}
