package functions

import (
	"fmt"
	"strings"
)

// containsSpecialCharacters checks for special characters in the text and prints an error if found
func containsSpecialCharacters(text string) bool {
	specialChars := map[string]string{
		"\\t": "Tab",
		"\\b": "Backspace",
		"\\v": "Vertical Tab",
		"\\0": "Null",
		"\\f": "Form Feed",
		"\\r": "Carriage Return",
	}

	for spChar, description := range specialChars {
		if strings.Contains(text, spChar) {
			fmt.Printf("Print Error: Special ASCII character (%s) or (%s) detected \n", spChar, description)
			return true
		}
	}

	return false
}
