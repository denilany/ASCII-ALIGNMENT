package functions

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Arguments struct {
	input  string
	flag   string
	banner string
}

func ParseArguments() (Arguments, error) {
	args := Arguments{}
	cmdArgs := os.Args[1:]

	if len(cmdArgs) == 0 {
		return args, errors.New("no arguments provided")
	}

	// Default values
	args.flag = "left"
	args.banner = "standard"

	for _, arg := range cmdArgs {
		if strings.HasPrefix(arg, "--align=") {
			alignment := strings.TrimPrefix(arg, "--align=")
			switch alignment {
			case "left", "right", "center", "justify":
				args.flag = alignment
			default:
				return args, fmt.Errorf("invalid alignment: %s", alignment)
			}
		} else if strings.HasPrefix(arg, "--") {
			// Catch any unrecognized flags
			return args, errors.New("invalid option: " + arg)
		} else if args.input == "" {
			args.input = arg
		} else if args.banner == "standard" {
			args.banner = arg
		} else {
			return args, fmt.Errorf("unexpected argument: %s", arg)
		}
	}

	if args.input == "" {
		return args, errors.New("input string is required")
	}

	return args, nil
}

func AsciiValue() string {
	args, err := ParseArguments()
	if err != nil {
		PrintUsage()
		return err.Error()
	}

	bannerPath := "banner/" + args.banner + ".txt"
	bannerMap, err := ReadAsciiArt(bannerPath)
	if err != nil {
		return "Error loading banner: " + err.Error()
	}

	termWidth, _, err := getTerminalSize()
	if err != nil {
		return err.Error()
	}

	printAsciiArt(args.input, bannerMap, args.flag, termWidth)
	return ""
}

func PrintUsage() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
	fmt.Println()
	fmt.Println("Example: go run . --align=right 'Hello World' standard")
}
