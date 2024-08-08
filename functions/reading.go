package functions

import (
	"bufio"
	"fmt"
	"os"
)

func ReadAscii(banner string) ([]string, error) {
	file, err := os.Open(banner)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	if fileInfo.Size() == 0 {
		fmt.Printf("file appears to be empty\n")
		os.Exit(1)
	}

	newScan := bufio.NewScanner(file)
	var splitData []string
	for newScan.Scan() {
		line := newScan.Text()
		splitData = append(splitData, line)
	}
	return splitData, nil
}
