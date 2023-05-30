package separated

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
		values  []string
		err     error
	}{
		"option error": {
			options: []Option{badOption},
			err:     errors.New("option error: some error"),
		},
		"empty_string": {
			value:  "a",
			values: []string{"a"},
		},
		"single value": {
			value:  "a",
			values: []string{"a"},
		},
		"multiple values": {
			value:  "a,b",
			values: []string{"a", "b"},
		},
		"lowercase values": {
			value:   "A,B",
			options: []Option{OptionLowercase()},
			values:  []string{"a", "b"},
		},
		"custom separator": {
			value:   "a;b",
			options: []Option{OptionSeparator(";")},
			values:  []string{"a", "b"},
		},
		"bad accepted value": {
			value:   "a,b",
			options: []Option{OptionAccepted("a")},
			err:     errors.New("value is not valid: value \"b\" at position 2; accepted values are: a"),
		},
		"all accepted values": {
			value:   "a,b",
			options: []Option{OptionAccepted("b", "a")},
			values:  []string{"a", "b"},
		},
		"ignore empty": {
			value:   "",
			options: []Option{OptionIgnoreEmpty()},
			values:  []string{},
		},
		"ignore empty multiple": {
			value:   "a,,,b",
			options: []Option{OptionIgnoreEmpty()},
			values:  []string{"a", "b"},
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			values, err := Validate(testCase.value, testCase.options...)

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

			if len(values) != len(testCase.values) {
				t.Fatalf("expected %d values but got %d values", len(testCase.values), len(values))
			}

			for i := range values {
				if testCase.values[i] != values[i] {
					t.Errorf("expected %s but got %s", testCase.values[i], values[i])
				}
			}
		})
	}
}
