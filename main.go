package main

import (
	"asciiweb/functions"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 4 {
		functions.PrintUsage()

		return
	}

	fmt.Println(functions.AsciiValue())
}
