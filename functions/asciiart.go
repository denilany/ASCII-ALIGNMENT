package functions

import (
	"fmt"
	"strings"
)

const (
	asciiHeight = 8
)

type Alignment int

const (
	Left Alignment = iota
	Center
	Right
	Justify
)

func AsciiArt(bannerSlice []string, input string, alignment Alignment, termWidth int) (string, error) {
	var result strings.Builder

	input = replaceUnprint(input)
	input = replace(input)
	for _, ch := range input {
		if (ch < 32 || ch > 126) && string(ch) != "\r" && string(ch) != "\n" {
			return "", fmt.Errorf("contains a non-printable character")
		}
	}

	arguments := strings.Split(input, "\n")

	for _, word := range arguments {
		if word == "" {
			result.WriteString("\n")
		} else {
			lines := make([]string, asciiHeight)
			for j := 0; j < asciiHeight; j++ {
				for _, ch := range word {
					index := int(ch-32)*9 + 1
					// result.WriteString(bannerSlice[index+j])
					lines[j] += bannerSlice[index+j]
				}
				// result.WriteString("\n")
			}

			for _, line := range lines {
				alignedLine := alignLine(line, alignment, termWidth)
				result.WriteString(alignedLine + "\n")
			}
		}
	}
	return result.String(), nil
}
