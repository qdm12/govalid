package integer

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
		integer int
		err     error
	}{
		"option error": {
			options: []Option{badOption},
			err:     errors.New("option error: some error"),
		},
		"bad string": {
			value: "a",
			err:   errors.New("value is not an integer: a"),
		},
		"valid integer": {
			value:   "1",
			integer: 1,
		},
		"within range": {
			options: []Option{OptionRange(1, 3)},
			value:   "2",
			integer: 2,
		},
		"out of range": {
			options: []Option{OptionRange(1, 3)},
			value:   "4",
			err:     errors.New("value is too big: 4 is bigger than maximum of 3"),
		},
		"above minimum": {
			options: []Option{OptionMinimum(1)},
			value:   "2",
			integer: 2,
		},
		"below minimum": {
			options: []Option{OptionMinimum(1)},
			value:   "0",
			err:     errors.New("value is too small: 0 is smaller than minimum of 1"),
		},
		"above maximum": {
			options: []Option{OptionMaximum(1)},
			value:   "2",
			err:     errors.New("value is too big: 2 is bigger than maximum of 1"),
		},
		"below maximum": {
			options: []Option{OptionMaximum(1)},
			value:   "-1",
			integer: -1,
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			integer, err := Validate(testCase.value, testCase.options...)

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

			if testCase.integer != integer {
				t.Errorf("expected integer %d but got %d", testCase.integer, integer)
			}
		})
	}
}
