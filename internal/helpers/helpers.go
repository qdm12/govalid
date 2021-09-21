package helpers

import (
	"regexp"
	"strings"
)

func MatchRegex(regex string) *regexp.Regexp {
	return regexp.MustCompile("^" + regex + "$")
}

func CommaJoin(values []string) (joined string) {
	return strings.Join(values, ", ")
}
