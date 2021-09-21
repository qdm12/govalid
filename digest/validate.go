package digest

import (
	"errors"
	"fmt"

	"github.com/qdm12/govalid/internal/helpers"
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
			return fmt.Errorf("%w: %s", ErrOption, err)
		}
	}

	ok := false
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
	match32BytesHex = helpers.MatchRegex(regex32BytesHex)
	match64BytesHex = helpers.MatchRegex(regex64BytesHex)
)
