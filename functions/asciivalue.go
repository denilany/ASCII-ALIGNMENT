package functions

import (
	"fmt"
	"os"
	"strings"
)

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

func PrintUsage() {
	fmt.Println("Usage go run . [OPTION] [STRING] [BANNER]")
	fmt.Println()
	fmt.Println("Example: go run . --align=right something standard")
}
