package functions

import (
	"fmt"
	"strings"
)

func printAsciiArt(text string, asciiArtMap map[rune]string, alignment string, termWidth int) {
	if text == "" {
		return
	}

	if containsSpecialCharacters(text) {
		return
	}

	asciiArtLines := [asciiArtHeight]string{}
	spacePositions := []int{}
	totalTextWidth := 0

	for _, char := range text {
		art, exists := asciiArtMap[char]
		if !exists {
			fmt.Println("Character not found in ASCII art map")
			return
		}

		lines := strings.Split(art, "\n")
		for i := 0; i < asciiArtHeight; i++ {
			if i < len(lines) {
				asciiArtLines[i] += lines[i]
			} else {
				asciiArtLines[i] += strings.Repeat(" ", len(lines[0]))
			}
		}

		if char == ' ' {
			spacePositions = append(spacePositions, totalTextWidth)
		}

		totalTextWidth += len(lines[0])
	}

	alignedString := applyTextAlignment(asciiArtLines[:], alignment, termWidth, totalTextWidth, spacePositions)

	fmt.Println(alignedString)
}

// applyTextAlignment aligns the ASCII art lines based on the specified alignment
// and returns a single aligned string
func applyTextAlignment(lines []string, alignment string, termWidth, totalTextWidth int, spacePositions []int) string {
	var alignedLines []string

	switch alignment {
	case "left":
		alignedLines = lines
	case "center":
		leftPadding := (termWidth - totalTextWidth) / 2
		for _, line := range lines {
			alignedLines = append(alignedLines, strings.Repeat(" ", leftPadding)+line)
		}
	case "right":
		leftPadding := termWidth - totalTextWidth
		for _, line := range lines {
			alignedLines = append(alignedLines, strings.Repeat(" ", leftPadding)+line)
		}
	case "justify":
		if totalTextWidth < termWidth && len(spacePositions) > 0 {
			extraSpaces := termWidth - totalTextWidth
			spacesToAdd := extraSpaces / len(spacePositions)
			remainder := extraSpaces % len(spacePositions)

			for _, line := range lines {
				newLine := []rune(line)
				offset := 0
				for j, pos := range spacePositions {
					additionalSpaces := spacesToAdd
					if j < remainder {
						additionalSpaces++
					}
					adjustedPos := pos + offset
					newLine = append(newLine[:adjustedPos], append([]rune(strings.Repeat(" ", additionalSpaces)), newLine[adjustedPos:]...)...)
					offset += additionalSpaces
				}
				alignedLines = append(alignedLines, string(newLine))
			}
		} else {
			alignedLines = lines
		}
	}

	return strings.Join(alignedLines, "\n")
}
