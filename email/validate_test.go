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
			err:     errors.New("option error: some error"),
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
			err:   errors.New("email format is not valid: aa.com"),
		},
		"gmail MX lookup": {
			value:   "a@gmail.com",
			options: []Option{OptionMXLookup()},
			email:   "a@gmail.com",
		},
		"localhost.com MX lookup": {
			value:   "a@localhost.example",
			options: []Option{OptionMXLookup()},
			err:     errors.New("email host is not reachable: lookup localhost.example on 127.0.0.11:53: no such host"),
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			email, err := Validate(testCase.value, testCase.options...)

			if testCase.err != nil {
				if err == nil {
					t.Fatalf("expected an error but got nil instead")
				} else if testCase.err.Error() != err.Error() {
					t.Errorf("expected error %q but got %q", testCase.err, err)
				}
			} else {
				if err != nil {
					t.Errorf("received an unexpected error %q", err)
				}
			}

			if testCase.email != email {
				t.Errorf("expected email %q but got %q", testCase.email, email)
			}
		})
	}
}
