package functions

import (
	"fmt"
	"os"
	"strings"
)

type Arguments struct {
	input  string
	flag   []string
	banner string
}

func ParseArguments() (Arguments, error) {
	args := Arguments{}

	cmdArgs := os.Args[1:]

	if len(cmdArgs) == 1 {
		args.input = cmdArgs[0]
		args.flag = append(args.flag, "left")
		args.banner = "standard"
		return args, nil
	} else if len(cmdArgs) == 2 {
		args.input = cmdArgs[0]
		args.flag = append(args.flag, "left")
		args.banner = cmdArgs[1]
		return args, nil
	}

	for i, arg := range cmdArgs {
		if strings.HasPrefix(arg, "--align=") {
			alignment := strings.TrimPrefix(arg, "--align=")
			switch alignment {
			case "left", "right", "center", "justify":
				args.flag = append(args.flag, alignment)
			default:
				return args, fmt.Errorf("invalid alignment: %s", alignment)
			}
		} else if i == 1 {
			args.input = arg
		} else if i == 2 {
			args.banner = arg
		}
	}

	if args.input == "" || args.banner == "" {
		return args, fmt.Errorf("missing input or banner")
	}

	return args, nil
}

func AsciiValue() string {
	args, err := ParseArguments()
	if err != nil {
		PrintUsage()
		return err.Error()
	}

	var alignment string
	if len(args.flag) > 0 {
		alignment = args.flag[0]
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

	printAsciiArt(args.input, bannerMap, alignment, termWidth)
	return ""
}

func PrintUsage() {
	fmt.Println("Usage go run . [OPTION] [STRING] [BANNER]")
	fmt.Println()
	fmt.Println("Example: go run . --align=right something standard")
}
