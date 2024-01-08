package port

import (
	"errors"
	"fmt"
	"testing"
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
		port       uint16
		errWrapped error
		errMessage string
	}{
		"option_error": {
			options:    []Option{badOption},
			errWrapped: ErrOption,
			errMessage: "option error: test error",
		},
		"bad_string": {
			value:      "a",
			errWrapped: ErrPortNotAnInteger,
			errMessage: "port value is not an integer: a",
		},
		"negative_port": {
			value:      "-1",
			errWrapped: ErrPortNegative,
			errMessage: "port cannot be negative: -1",
		},
		"port_too_high": {
			value:      "70000",
			errWrapped: ErrPortTooHigh,
			errMessage: "port cannot be higher than 65535: 70000",
		},
		"not_listening_zero_port": {
			value:      "0",
			errWrapped: ErrPortZero,
			errMessage: "port cannot be zero",
		},
		"not_listening_valid": {
			value: "80",
			port:  80,
		},
		"listening_non_privileged_port": {
			value:   "2000",
			options: []Option{OptionPortListening(1000)},
			port:    2000,
		},
		"listening_zero_port": {
			value:   "0",
			options: []Option{OptionPortListening(1000)},
		},
		"listening_zero_port_disallowed": {
			value:      "0",
			options:    []Option{OptionPortListening(1000, OptionListeningPortZero(true))},
			errWrapped: ErrListenPortZero,
			errMessage: "listening port cannot be zero",
		},
		"listening_privileged_port_as_root": {
			value:   "100",
			options: []Option{OptionPortListening(0)},
			port:    100,
		},
		"listening_privileged_port_as_windows": {
			value:   "100",
			options: []Option{OptionPortListening(-1)},
			port:    100,
		},
		"listening_privileged_port_as_uid_1000": {
			value:      "100",
			options:    []Option{OptionPortListening(1000)},
			errWrapped: ErrListenPrivilegedPort,
			errMessage: "listening port cannot be privileged port: " +
				"100 when running with uid 1000",
		},
		"listening_privileged_port_as_uid_1000_allowed": {
			value:   "100",
			options: []Option{OptionPortListening(1000, OptionListeningPortPrivilegedAllowed(100))},
			port:    100,
		},
		"listening_privileged_port_as_uid_1000_not_allowed": {
			value:      "200",
			options:    []Option{OptionPortListening(1000, OptionListeningPortPrivilegedAllowed(100))},
			errWrapped: ErrListenPrivilegedPort,
			errMessage: "listening port cannot be privileged port: " +
				"port 200 is not part of allowed ports 100",
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			port, err := Validate(testCase.value, testCase.options...)

			if !errors.Is(err, testCase.errWrapped) {
				t.Errorf("expected error %q to wrap error %q", err, testCase.errWrapped)
			}

			if testCase.errWrapped != nil {
				if err.Error() != testCase.errMessage {
					t.Errorf("expected error %q but got %q", testCase.errMessage, err)
				}
			}

			if testCase.port != port {
				t.Errorf("expected port %d but got %d", testCase.port, port)
			}
		})
	}
}
