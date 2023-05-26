package digest

import (
	"errors"
	"fmt"
	"regexp"
)

var (
	ErrOption             = errors.New("option error")
	ErrDigestMalformed    = errors.New("digest is malformed")
	ErrDigestTypeNotValid = errors.New("digest type is not valid")
)

// Validate verifies the digest string matches the expected digest type format.
func Validate(value string, digestType Type, options ...Option) (err error) {
	s := newSettings()
	for _, option := range options {
		err := option(s)
		if err != nil {
			return fmt.Errorf("%w: %w", ErrOption, err)
		}
	}

	var ok bool
	switch digestType {
	case SHA256Hex:
		ok = match64BytesHex.MatchString(value)
	case MD5Hex:
		ok = match32BytesHex.MatchString(value)
	default:
		return fmt.Errorf("%w: %s", ErrDigestTypeNotValid, digestType)
	}

	if !ok {
		return fmt.Errorf("%w: for format %s: %s", ErrDigestMalformed, digestType, value)
	}

	return nil
}

const (
	regex32BytesHex = `[a-fA-F0-9]{32}`
	regex64BytesHex = `[a-fA-F0-9]{64}`
)

var (
	match32BytesHex = regexp.MustCompile("^" + regex32BytesHex + "$")
	match64BytesHex = regexp.MustCompile("^" + regex64BytesHex + "$")
)
