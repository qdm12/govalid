package helpers

import (
	"strings"
)

func CommaJoin(values []string) (joined string) {
	return strings.Join(values, ", ")
}
