package functions

import (
	"fmt"
	"strings"
)

const asciiHeight = 8

// printAsciiArt prints ASCII art with specified alignment and width
func printAsciiArt(text string, asciiArtMap map[rune]string, alignment string, termWidth int) {
	if text == "" {
		return
	}

	if containsSpecialCharacters(text) {
		return
	}

	asciiArtLines := [asciiHeight]string{}
	spacePositions := []int{}
	totalTextWidth := 0

	for _, char := range text {
		art, exists := asciiArtMap[char]
		if !exists {
			fmt.Println("Character not found in ASCII art map")
			return
		}

		lines := strings.Split(art, "\n")
		for i := 0; i < asciiHeight; i++ {
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

	applyTextAlignment(asciiArtLines[:], alignment, termWidth, totalTextWidth, spacePositions)

	for _, line := range asciiArtLines {
		fmt.Println(line)
	}
}

// applyTextAlignment aligns the ASCII art lines based on the specified alignment
func applyTextAlignment(lines []string, alignment string, termWidth, totalTextWidth int, spacePositions []int) {
	switch alignment {
	case "center":
		padding := (termWidth - totalTextWidth) / 2
		for i := range lines {
			lines[i] = strings.Repeat(" ", padding) + lines[i]
		}
	case "right":
		padding := termWidth - totalTextWidth
		for i := range lines {
			lines[i] = strings.Repeat(" ", padding) + lines[i]
		}
	case "left":
		// No additional padding needed for left alignment
	case "justify":
		if totalTextWidth < termWidth && len(spacePositions) > 0 {
			extraSpaces := termWidth - totalTextWidth
			spacesToAdd := extraSpaces / len(spacePositions)
			remainder := extraSpaces % len(spacePositions)

			for i := range lines {
				newLine := []rune(lines[i])
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
				lines[i] = string(newLine)
			}
		}
	}
}
