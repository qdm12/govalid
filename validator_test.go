package govalid

import (
	"net/url"
	"testing"
	"time"

	"github.com/qdm12/govalid/digest"
)

func Test_Validator(t *testing.T) {
	t.Parallel()

	t.Run("address", func(t *testing.T) {
		t.Parallel()
		const value = ":8000"
		err := ValidateAddress(value)
		assertNoError(t, err)
	})

	t.Run("binary", func(t *testing.T) {
		t.Parallel()
		const value = "on"
		output, err := ValidateBinary(value)
		assertNoError(t, err)
		assertBool(t, true, output)
	})

	t.Run("digest", func(t *testing.T) {
		t.Parallel()
		const value = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
		err := ValidateDigest(value, digest.MD5Hex)
		assertNoError(t, err)
	})

	t.Run("duration", func(t *testing.T) {
		t.Parallel()
		const value = "1s"
		output, err := ValidateDuration(value)
		assertNoError(t, err)
		assertDuration(t, time.Second, output)
	})

	t.Run("email", func(t *testing.T) {
		t.Parallel()
		const value = "a@a.com"
		output, err := ValidateEmail(value)
		assertNoError(t, err)
		assertString(t, "a@a.com", output)
	})

	t.Run("integer", func(t *testing.T) {
		t.Parallel()
		const value = "5"
		output, err := ValidateInteger(value)
		assertNoError(t, err)
		assertInteger(t, 5, output)
	})

	t.Run("port", func(t *testing.T) {
		t.Parallel()
		const value = "5000"
		output, err := ValidatePort(value)
		assertNoError(t, err)
		assertInteger(t, 5000, int(output))
	})

	t.Run("root url", func(t *testing.T) {
		t.Parallel()
		const value = "/rooturl"
		output, err := ValidateRootURL(value)
		assertNoError(t, err)
		assertString(t, "/rooturl", output)
	})

	t.Run("separated", func(t *testing.T) {
		t.Parallel()
		const value = "a,b,c"
		output, err := ValidateSeparated(value)
		assertNoError(t, err)
		assertStringsSlice(t, []string{"a", "b", "c"}, output)
	})

	t.Run("url", func(t *testing.T) {
		t.Parallel()
		const value = "http://a.com"
		output, err := ValidateURL(value)
		assertNoError(t, err)
		assertURL(t, &url.URL{Scheme: "http", Host: "a.com"}, output)
	})
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("expected no error but got %s", err)
	}
}

func assertString(t *testing.T, expected, actual string) {
	t.Helper()
	if expected != actual {
		t.Errorf("expected %q got %q", expected, actual)
	}
}

func assertBool(t *testing.T, expected, actual bool) {
	t.Helper()
	if expected != actual {
		t.Errorf("expected %t got %t", expected, actual)
	}
}

func assertDuration(t *testing.T, expected, actual time.Duration) {
	t.Helper()
	if expected != actual {
		t.Errorf("expected %s got %s", expected, actual)
	}
}

func assertInteger(t *testing.T, expected, actual int) {
	t.Helper()
	if expected != actual {
		t.Errorf("expected %d got %d", expected, actual)
	}
}

func assertStringsSlice(t *testing.T, expected, actual []string) {
	t.Helper()
	if len(actual) != len(expected) {
		t.Fatalf("expected %d values but got %d values", len(expected), len(actual))
	}

	for i := range actual {
		assertString(t, expected[i], actual[i])
	}
}

func assertURL(t *testing.T, expected, actual *url.URL) {
	switch {
	case expected == nil && actual == nil: // success
	case expected == nil && actual != nil:
		t.Fatalf("expected a nil url but got %s", actual.String())
	case expected != nil && actual == nil:
		t.Fatalf("expected a nil url but got %s", actual.String())
	case expected.String() != actual.String():
		t.Errorf("expected %s but got %s", expected.String(), actual.String())
	}
}
