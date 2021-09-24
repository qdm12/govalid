package rooturl

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
		rootURL string
		err     error
	}{
		"option error": {
			options: []Option{badOption},
			err:     errors.New("option error: some error"),
		},
		"bad string": {
			value: "a",
			err:   errors.New("value is not valid: a"),
		},
		"dirty root url": {
			value:   " /some//path ",
			rootURL: "/some/path",
		},
		"trailing slash": {
			value:   "/some/path/",
			rootURL: "/some/path",
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			rootURL, err := Validate(testCase.value, testCase.options...)

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

			if testCase.rootURL != rootURL {
				t.Errorf("expected port %s but got %s", testCase.rootURL, rootURL)
			}
		})
	}
}
