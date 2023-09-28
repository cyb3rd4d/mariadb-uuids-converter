package encoding

import (
	"fmt"
	"strings"
)

func Encode(input string) (string, error) {
	parts := strings.Split(input, "-")
	if len(parts) != 5 {
		return "", fmt.Errorf("invalid input, should have 5 parts delimited by dashes, got %d", len(parts))
	}

	return strings.Join(parts, ""), nil
}
