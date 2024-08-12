package functions

import (
	"fmt"
	"os"
	"strings"
)

func AsciiValue() string {
	var alignment string
	var input, bannerPath string

	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--align=") {
			switch strings.TrimPrefix(arg, "--align=") {
			case "left":
				alignment = "left"
			case "right":
				alignment = "right"
			case "center":
				alignment = "center"
			case "justify":
				alignment = "justify"
			default:
				PrintUsage()
				return ""
			}
		}
	}
	input = os.Args[2]
	bannerPath = "banner/" + os.Args[3] + ".txt"

	bannerMap, err := ReadAsciiArt(bannerPath)
	if err != nil {
		return "Error loading banner: " + err.Error()
	}

	termWidth, _, err := getTerminalSize()
	if err != nil {
		return err.Error()
	}

	printAsciiArt(input, bannerMap, alignment, termWidth)

	return ""
}

func PrintUsage() {
	fmt.Println("Usage go run . [OPTION] [STRING] [BANNER]")
	fmt.Println()
	fmt.Println("Example: go run . --align=right something standard")
}
