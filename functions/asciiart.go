package functions

import (
	"fmt"
	"os"
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

func alignLine(line string, alignment Alignment, termWidth int) string {
	lineWidth := len(line)
	if lineWidth >= termWidth {
		return line[:termWidth]
	}

	switch alignment {
	case Left:
		return line + strings.Repeat(" ", termWidth-lineWidth)
	case Right:
		return strings.Repeat(" ", termWidth-lineWidth) + line
	case Center:
		leftPad := (termWidth - lineWidth) / 2
		rightPad := termWidth - lineWidth - leftPad
		return strings.Repeat(" ", leftPad) + line + strings.Repeat(" ", rightPad)
	case Justify:
		if lineWidth == termWidth {
			return line
		}
		words := strings.Fields(line)
		if len(words) == 1 {
			return line + strings.Repeat(" ", termWidth-lineWidth)
		}
		spaces := termWidth - lineWidth + len(words) - 1
		spacePerGap := spaces / (len(words) - 1)
		extraSpaces := spaces % (len(words) - 1)
		var justifiedLine strings.Builder
		for i, word := range words {
			justifiedLine.WriteString(word)
			if i < len(words)-1 {
				justifiedLine.WriteString(strings.Repeat(" ", spacePerGap))
				if i < extraSpaces {
					justifiedLine.WriteString(" ")
				}
			}
		}
		return justifiedLine.String()
	default:
		return line
	}
}

func PrintUsage() {
	fmt.Println("Usage go run . [OPTION] [STRING] [BANNER]")
	fmt.Println()
	fmt.Println("Example: go run . --align=right something standard")
}

func AsciiValue() string {
	var alignment Alignment
	var input, banner string

	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--align=") {
			switch strings.TrimPrefix(arg, "--align=") {
			case "left":
				alignment = Left
			case "right":
				alignment = Right
			case "center":
				alignment = Center
			case "justify":
				alignment = Justify
			default:
				PrintUsage()
				return ""
			}

		}
	}
	input = os.Args[2]
	banner = "banner/" + os.Args[3] + ".txt"

	bannerSlice, err := ReadAscii(banner)
	if err != nil {
		// fmt.Println("Error loading banner:", err)
		return "Error loading banner: " + err.Error()
	}

	termWidth, _, err := getTerminalSize()
	if err != nil {
		return err.Error()
	}
	result, err := AsciiArt(bannerSlice, input, alignment, termWidth)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	return result
}
