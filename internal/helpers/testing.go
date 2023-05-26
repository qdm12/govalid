package helpers

import "testing"

func AssertError(t *testing.T, expected, actual error) {
	t.Helper()

	if expected != nil {
		if actual == nil {
			t.Errorf("expected an error but got nil instead")
		} else if expected.Error() != actual.Error() {
			t.Errorf("expected error %q but got %q", expected, actual)
		}
	} else {
		if actual != nil {
			t.Errorf("received an unexpected error %q", actual)
		}
	}
}
