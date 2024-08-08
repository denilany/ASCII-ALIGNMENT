package functions

import "strings"

func replace(input string) string {
	input = strings.ReplaceAll(input, "\\n", "\n")
	input = strings.ReplaceAll(input, "\\t", "    ")
	input = strings.ReplaceAll(input, "\\b", "\b")

	for {
		index := strings.Index(input, "\b")

		if index == -1 {
			break
		}
		if index > 0 {
			input = input[:index-1] + input[index+1:]
		} else {
			input = input[index+1:]
		}
	}

	return input
}
