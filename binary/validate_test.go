package binary

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
		binary  bool
		err     error
	}{
		"option error": {
			options: []Option{badOption},
			err:     errors.New("option error: some error"),
		},
		"default enabled": {
			value:  "yes",
			binary: true,
		},
		"default disabled": {
			value: "off",
		},
		"invalid value": {
			value: "value",
			err:   errors.New(`value is not valid: value "value" can only be one of enabled, yes, on, disabled, no, off`),
		},
		"enabled with option": {
			value:   "Custom",
			options: []Option{OptionEnabled("custom")},
			binary:  true,
		},
		"disabled with option": {
			value:   "Custom",
			options: []Option{OptionDisabled("custom")},
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			binary, err := Validate(testCase.value, testCase.options...)

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

			if testCase.binary != binary {
				t.Errorf("expected binary %t but got %t", testCase.binary, binary)
			}
		})
	}
}
