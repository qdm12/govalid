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

var _ Interface = (*Validator)(nil)

type Interface interface {
	ValidateAddress(value string, options ...address.Option) (addr string, err error)
	ValidateBinary(value string, options ...binary.Option) (enabled bool, err error)
	ValidateDigest(value string, digestType digest.Type, options ...digest.Option) (err error)
	ValidateDuration(value string, options ...duration.Option) (duration time.Duration, err error)
	ValidateEmail(value string, options ...email.Option) (email string, err error)
	ValidateInteger(value string, options ...integer.Option) (integer int, err error)
	ValidatePort(value string, options ...port.Option) (port uint16, err error)
	ValidateRootURL(value string, options ...rooturl.Option) (rootURL string, err error)
	ValidateSeparated(value string, options ...separated.Option) (slice []string, err error)
	ValidateURL(value string, options ...govalidurl.Option) (url *url.URL, err error)
}

type Validator struct{}

func New() *Validator {
	return &Validator{}
}

func (v *Validator) ValidateAddress(value string, options ...address.Option) (addr string, err error) {
	return address.Validate(value, options...)
}

func (v *Validator) ValidateBinary(value string, options ...binary.Option) (enabled bool, err error) {
	return binary.Validate(value, options...)
}

func (v *Validator) ValidateDigest(value string, digestType digest.Type, options ...digest.Option) (err error) {
	return digest.Validate(value, digestType, options...)
}

func (v *Validator) ValidateDuration(value string, options ...duration.Option) (d time.Duration, err error) {
	return duration.Validate(value, options...)
}
func (v *Validator) ValidateEmail(value string, options ...email.Option) (emailStr string, err error) {
	return email.Validate(value, options...)
}
func (v *Validator) ValidateInteger(value string, options ...integer.Option) (n int, err error) {
	return integer.Validate(value, options...)
}
func (v *Validator) ValidatePort(value string, options ...port.Option) (p uint16, err error) {
	return port.Validate(value, options...)
}
func (v *Validator) ValidateRootURL(value string, options ...rooturl.Option) (rootURL string, err error) {
	return rooturl.Validate(value, options...)
}
func (v *Validator) ValidateSeparated(value string, options ...separated.Option) (slice []string, err error) {
	return separated.Validate(value, options...)
}
func (v *Validator) ValidateURL(value string, options ...govalidurl.Option) (u *url.URL, err error) {
	return govalidurl.Validate(value, options...)
}
