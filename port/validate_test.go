package port

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
		port    uint16
		err     error
	}{
		"option error": {
			options: []Option{badOption},
			err:     errors.New("option error: some error"),
		},
		"bad string": {
			value: "a",
			err:   errors.New("port value is not an integer: a"),
		},
		"zero port not allowed": {
			value: "0",
			err:   errors.New("port cannot be lower than 1: 0"),
		},
		"zero port for listening": {
			value:   "0",
			options: []Option{OptionPortListening(1000)},
		},
		"negative port for listening": {
			value:   "-1",
			options: []Option{OptionPortListening(1000)},
			err:     errors.New("listening port cannot be lower than 0: -1"),
		},
		"port too high": {
			value: "70000",
			err:   errors.New("port cannot be higher than 65535: 70000"),
		},
		"privileged port as root": {
			value:   "100",
			options: []Option{OptionPortListening(0)},
			port:    100,
		},
		"privileged port as windows": {
			value:   "100",
			options: []Option{OptionPortListening(-1)},
			port:    100,
		},
		"privileged port without root": {
			value:   "100",
			options: []Option{OptionPortListening(1000)},
			err:     errors.New("invalid listening port: cannot use privileged ports (1 to 1023) when running without root: 100"),
		},
		"privileged port without root allowed": {
			value:   "100",
			options: []Option{OptionPortListening(1000, OptionListeningPortPrivilegedAllowed(100))},
			port:    100,
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			port, err := Validate(testCase.value, testCase.options...)

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

			if testCase.port != port {
				t.Errorf("expected port %d but got %d", testCase.port, port)
			}
		})
	}
}
