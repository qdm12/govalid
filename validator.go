package govalid

import (
	"net/url"
	"time"

	"github.com/qdm12/govalid/address"
	"github.com/qdm12/govalid/binary"
	"github.com/qdm12/govalid/digest"
	"github.com/qdm12/govalid/duration"
	"github.com/qdm12/govalid/email"
	"github.com/qdm12/govalid/integer"
	"github.com/qdm12/govalid/port"
	"github.com/qdm12/govalid/rooturl"
	"github.com/qdm12/govalid/separated"
	govalidurl "github.com/qdm12/govalid/url"
)

func ValidateAddress(value string, options ...address.Option) (err error) {
	return address.Validate(value, options...)
}

func ValidateBinary(value string, options ...binary.Option) (enabled *bool, err error) {
	return binary.Validate(value, options...)
}

func ValidateDigest(value string, digestType digest.Type, options ...digest.Option) (err error) {
	return digest.Validate(value, digestType, options...)
}

func ValidateDuration(value string, options ...duration.Option) (d time.Duration, err error) {
	return duration.Validate(value, options...)
}
func ValidateEmail(value string, options ...email.Option) (emailStr string, err error) {
	return email.Validate(value, options...)
}
func ValidateInteger(value string, options ...integer.Option) (n int, err error) {
	return integer.Validate(value, options...)
}
func ValidatePort(value string, options ...port.Option) (p uint16, err error) {
	return port.Validate(value, options...)
}
func ValidateRootURL(value string, options ...rooturl.Option) (rootURL string, err error) {
	return rooturl.Validate(value, options...)
}
func ValidateSeparated(value string, options ...separated.Option) (slice []string, err error) {
	return separated.Validate(value, options...)
}
func ValidateURL(value string, options ...govalidurl.Option) (u *url.URL, err error) {
	return govalidurl.Validate(value, options...)
}
