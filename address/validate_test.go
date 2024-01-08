package address

import (
	"errors"
	"fmt"
	"testing"

	"github.com/qdm12/govalid/port"
)

func Test_Validate(t *testing.T) {
	t.Parallel()

	errTest := errors.New("test error")
	badOption := func(_ *settings) error {
		return fmt.Errorf("%w", errTest)
	}

	testCases := map[string]struct {
		value      string
		options    []Option
		errWrapped error
		errMessage string
	}{
		"option error": {
			options:    []Option{badOption},
			errWrapped: errTest,
			errMessage: "option error: test error",
		},
		"missing semicolon": {
			value:      "1.2.3.4",
			errWrapped: ErrValueNotValid,
			errMessage: "value is not valid: address 1.2.3.4: missing port in address",
		},
		"listening on zero port": {
			value:   "1.2.3.4:0",
			options: []Option{OptionListening(1000)},
		},
		"listening on privileged port without root": {
			value:      "1.2.3.4:100",
			options:    []Option{OptionListening(1000)},
			errWrapped: port.ErrListenPrivilegedPort,
			errMessage: "listening port cannot be privileged port: " +
				"100 when running with uid 1000",
		},
		"listening on privileged port with root": {
			value:   "1.2.3.4:100",
			options: []Option{OptionListening(0)},
		},
		"address with zero port": {
			value:      "1.2.3.4:0",
			errWrapped: port.ErrPortZero,
			errMessage: "port cannot be zero",
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

			if !errors.Is(err, testCase.errWrapped) {
				t.Fatalf("expected error %q to wrap error %q", testCase.errWrapped, err)
			}
			if testCase.errWrapped != nil {
				if err.Error() != testCase.errMessage {
					t.Errorf("expected error %q but got %q", testCase.errMessage, err)
				}
			}
		})
	}
}
