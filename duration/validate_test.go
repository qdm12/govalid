package duration

import (
	"errors"
	"testing"
	"time"
)

func Test_Validate(t *testing.T) {
	t.Parallel()

	badOption := func(_ *settings) error {
		return errors.New("some error")
	}

	testCases := map[string]struct {
		value    string
		options  []Option
		duration time.Duration
		err      error
	}{
		"option error": {
			options: []Option{badOption},
			err:     errors.New("option error: some error"),
		},
		"invalid string": {
			value: "abc",
			err:   errors.New("duration is malformed: abc: time: invalid duration \"abc\""),
		},
		"some duration": {
			value:    "1s",
			duration: time.Second,
		},
		"negative duration not allowed": {
			value: "-1s",
			err:   errors.New("duration cannot be negative: -1s"),
		},
		"negative duration": {
			value:    "-1s",
			options:  []Option{OptionAllowNegative()},
			duration: -time.Second,
		},
		"zero duration not allowed": {
			value: "0",
			err:   errors.New("duration cannot be zero: 0"),
		},
		"zero duration": {
			value:   "0s",
			options: []Option{OptionAllowZero()},
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			duration, err := Validate(testCase.value, testCase.options...)

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

			if testCase.duration != duration {
				t.Errorf("expected duration %q but got %q", testCase.duration, duration)
			}
		})
	}
}
