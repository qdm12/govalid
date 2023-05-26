package address

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
		err     error
	}{
		"option error": {
			options: []Option{badOption},
			err:     errors.New("option error: some error"),
		},
		"missing semicolon": {
			value: "1.2.3.4",
			err:   errors.New("value is not valid: address 1.2.3.4: missing port in address"),
		},
		"listening on zero port": {
			value:   "1.2.3.4:0",
			options: []Option{OptionListening(1000)},
		},
		"listening on privileged port without root": {
			value:   "1.2.3.4:100",
			options: []Option{OptionListening(1000)},
			err: errors.New("invalid listening port: " +
				"cannot use privileged ports (1 to 1023) when running " +
				"without root: 100"),
		},
		"listening on privileged port with root": {
			value:   "1.2.3.4:100",
			options: []Option{OptionListening(0)},
		},
		"address with zero port": {
			value: "1.2.3.4:0",
			err:   errors.New("port cannot be lower than 1: 0"),
		},
		"valid address": {
			value: "1.2.3.4:8000",
		},
		"valid listening address without root": {
			value:   "1.2.3.4:8000",
			options: []Option{OptionListening(1000)},
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := Validate(testCase.value, testCase.options...)

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
		})
	}
}
