package shared

import "strings"

func CapitalizeFirstLetter(input string) string {
	input = strings.TrimSpace(input)

	if len(input) == 0 {
		return input
	}

	return strings.ToUpper(input[:1]) + strings.ToLower(input[1:])
}
