package email

import (
	"errors"
	"testing"
)

func Test_Validate(t *testing.T) {
	t.Parallel()

	badOption := func(_ *settings) error {
		return errors.New("some error")
	}

	testCases := map[string]struct {
		value   string
		options []Option
		email   string
		err     error
	}{
		"option error": {
			options: []Option{badOption},
			err:     ErrOption,
		},
		"valid email": {
			value: "a@a.com",
			email: "a@a.com",
		},
		"email with spaces around": {
			value: " a@a.com ",
			email: "a@a.com",
		},
		"bad format": {
			value: " aa.com ",
			err:   ErrEmailFormatNotValid,
		},
		"gmail MX lookup": {
			value:   "a@gmail.com",
			options: []Option{OptionMXLookup()},
			email:   "a@gmail.com",
		},
		"localhost.com MX lookup": {
			value:   "a@localhost.example",
			options: []Option{OptionMXLookup()},
			err:     ErrEmailHostUnreachable,
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			email, err := Validate(testCase.value, testCase.options...)

			if !errors.Is(err, testCase.err) {
				t.Errorf("expected error %q but got %q", testCase.err, err)
			}

			if testCase.email != email {
				t.Errorf("expected email %q but got %q", testCase.email, email)
			}
		})
	}
}
