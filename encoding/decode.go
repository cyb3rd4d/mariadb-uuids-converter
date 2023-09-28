package encoding

import (
	"strings"
)

func Decode(input string) (string, error) {
	partLengths := []int{8, 4, 4, 5}
	parts := []string{}
	cursor := 0

	for _, l := range partLengths {
		if cursor == 0 {
			parts = append(parts, input[:l])
			cursor = l
			continue
		}

		next := cursor + l
		parts = append(parts, input[cursor:next])
		cursor = next
	}

	parts = append(parts, input[cursor:])
	return strings.Join(parts, "-"), nil
}
