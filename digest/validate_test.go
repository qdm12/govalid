package digest

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
		value      string
		digestType Type
		options    []Option
		err        error
	}{
		"option error": {
			options: []Option{badOption},
			err:     errors.New("option error: some error"),
		},
		"invalid digest type": {
			digestType: Type("invalid"),
			err:        errors.New("digest type is not valid: invalid"),
		},
		"valid SHA256 hex": {
			value:      "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			digestType: SHA256Hex,
		},
		"invalid SHA256 hex": {
			value:      "abcde",
			digestType: SHA256Hex,
			err:        errors.New("digest is malformed: for format sha256 hex: abcde"),
		},
		"valid MD5 hex": {
			value:      "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			digestType: MD5Hex,
		},
		"invalid MD5 hex": {
			value:      "abcde",
			digestType: MD5Hex,
			err:        errors.New("digest is malformed: for format md5 hex: abcde"),
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := Validate(testCase.value, testCase.digestType, testCase.options...)

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
