package url

import (
	"errors"
	"net/url"
	"testing"

	"github.com/qdm12/govalid/internal/helpers"
)

func Test_Validate(t *testing.T) {
	t.Parallel()

	badOption := func(_ *settings) error {
		return errors.New("some error")
	}

	testCases := map[string]struct {
		value   string
		options []Option
		u       *url.URL
		err     error
	}{
		"option error": {
			options: []Option{badOption},
			err:     errors.New("option error: some error"),
		},
		"bad url": {
			value: "%^&",
			err:   errors.New("url is not valid: parse \"%^&\": invalid URL escape \"%^&\""),
		},
		"bad scheme": {
			value: "ftp://a.com",
			err:   errors.New("url scheme is not valid: ftp://a.com"),
		},
		"valid scheme": {
			value: "HTTPS://a.com",
			u:     &url.URL{Scheme: "https", Host: "a.com"},
		},
		"custom valid scheme": {
			value:   "ftp://a.com",
			options: []Option{OptionAllowSchemes("ftp")},
			u:       &url.URL{Scheme: "ftp", Host: "a.com"},
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			u, err := Validate(testCase.value, testCase.options...)

			helpers.AssertError(t, testCase.err, err)

			switch {
			case testCase.u == nil && u == nil: // success
			case testCase.u == nil && u != nil:
				t.Fatalf("expected a nil url but got %s", u.String())
			case testCase.u != nil && u == nil:
				t.Fatalf("expected a nil url but got %s", u.String())
			case testCase.u.String() != u.String():
				t.Errorf("expected %s but got %s", testCase.u.String(), u.String())
			}
		})
	}
}
