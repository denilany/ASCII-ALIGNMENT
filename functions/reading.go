package functions

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	hashStandard   = "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
	hashShadow     = "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
	hashThinkerToy = "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3"
)

func ReadAscii(banner string) ([]string, error) {
	file, err := os.Open(banner)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteSlice, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	fileHash := checkBanner(byteSlice)

	bannerFile := banner[strings.LastIndex(banner, "/")+1:]
	fmt.Println(bannerFile)

	switch bannerFile {
	case "standard.txt":
		if fileHash != hashStandard {
			return nil, errors.New("the banner file \"standard.txt\" is corrupted")
		}
	case "shadow.txt":
		if fileHash != hashShadow {
			return nil, errors.New("the banner file \"shadow.txt\" is corrupted")
		}
	case "thinkertoy.txt":
		if fileHash != hashThinkerToy {
			return nil, errors.New("the banner file \"thinkertoy.txt\" is corrupted")
		}
	default:
		return nil, errors.New("unknown banner file")
	}

	// Reset the file cursor to the beginning for the scanner
	file.Seek(0, io.SeekStart)

	newScan := bufio.NewScanner(file)
	var splitData []string
	for newScan.Scan() {
		line := newScan.Text()
		splitData = append(splitData, line)
	}
	return splitData, nil
}
