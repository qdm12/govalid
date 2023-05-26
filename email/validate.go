package email

import (
	"errors"
	"fmt"
	"net"
	"regexp"
	"strings"
)

var (
	ErrOption               = errors.New("option error")
	ErrEmailFormatNotValid  = errors.New("email format is not valid")
	ErrEmailHostUnreachable = errors.New("email host is not reachable")
)

const regexEmail = `[a-zA-Z0-9-_.+]+@[a-zA-Z0-9-_.]+\.[a-zA-Z]{2,10}`

var emailMatcher = regexp.MustCompile("^" + regexEmail + "$")

// Validate verifies the value is an email address and does
// additional checks for any option given.
func Validate(value string, options ...Option) (
	email string, err error) {
	s := settings{}
	for _, option := range options {
		err := option(&s)
		if err != nil {
			return "", fmt.Errorf("%w: %w", ErrOption, err)
		}
	}

	email = strings.TrimSpace(value)

	if !emailMatcher.MatchString(email) {
		return "", fmt.Errorf("%w: %s", ErrEmailFormatNotValid, email)
	}

	if s.mxLookup {
		err = emailMxLookup(email)
		if err != nil {
			return "", fmt.Errorf("%w: %w", ErrEmailHostUnreachable, err)
		}
	}

	return email, nil
}

func emailMxLookup(email string) (err error) {
	i := strings.LastIndexByte(email, '@')
	host := email[i+1:]
	_, err = net.LookupMX(host)
	return err
}
