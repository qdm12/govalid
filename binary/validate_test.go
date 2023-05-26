package binary

import (
	"errors"
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
		binary  *bool
		err     error
	}{
		"option error": {
			options: []Option{badOption},
			err:     errors.New("option error: some error"),
		},
		"default enabled": {
			value:  "yes",
			binary: ptrToBool(true),
		},
		"default disabled": {
			value:  "off",
			binary: ptrToBool(false),
		},
		"invalid value": {
			value: "value",
			err:   errors.New(`value is not valid: value "value" can only be one of enabled, yes, on, disabled, no, off`),
		},
		"enabled with option": {
			value:   "Custom",
			options: []Option{OptionEnabled("custom")},
			binary:  ptrToBool(true),
		},
		"disabled with option": {
			value:   "Custom",
			options: []Option{OptionDisabled("custom")},
			binary:  ptrToBool(false),
		},
		"empty string": {},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			binary, err := Validate(testCase.value, testCase.options...)

			helpers.AssertError(t, testCase.err, err)

			switch {
			case testCase.binary == nil && binary == nil:
			case testCase.binary == nil && binary != nil:
				t.Errorf("expected binary to be nil but got %v", binary)
			case testCase.binary != nil && binary == nil:
				t.Errorf("expected binary to be %v but got nil", *testCase.binary)
			case *testCase.binary != *binary:
				t.Errorf("expected binary %t but got %t", *testCase.binary, *binary)
			}
		})
	}
}
