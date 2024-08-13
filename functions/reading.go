package functions

import (
    "bufio"
    "errors"
    "io"
    "os"
    "strings"
)

const (
    hashStandard   = "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
    hashShadow     = "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
    hashThinkerToy = "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3"
    asciiArtHeight = 8
)

func ReadAsciiArt(bannerFilePath string) (map[rune]string, error) {
    bannerContent, err := readFile(bannerFilePath)
    if err != nil {
        return nil, err
    }

    err = validateBannerFile(bannerFilePath, bannerContent)
    if err != nil {
        return nil, err
    }

    asciiArtMap, err := parseAsciiArt(bannerContent)
    if err != nil {
        return nil, err
    }

    return asciiArtMap, nil
}

func readFile(filePath string) ([]byte, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    return io.ReadAll(file)
}

func validateBannerFile(filePath string, content []byte) error {
    fileHash := checkBanner(content)
    fileName := filePath[strings.LastIndex(filePath, "/")+1:]

    switch fileName {
    case "standard.txt":
        if fileHash != hashStandard {
            return errors.New("the banner file \"standard.txt\" is corrupted")
        }
    case "shadow.txt":
        if fileHash != hashShadow {
            return errors.New("the banner file \"shadow.txt\" is corrupted")
        }
    case "thinkertoy.txt":
        if fileHash != hashThinkerToy {
            return errors.New("the banner file \"thinkertoy.txt\" is corrupted")
        }
    default:
        return errors.New("unknown banner file")
    }

    return nil
}

func parseAsciiArt(content []byte) (map[rune]string, error) {
    scanner := bufio.NewScanner(strings.NewReader(string(content)))
    asciiArtMap := make(map[rune]string)
    currentChar := ' ' 
    var currentArtBuilder strings.Builder
    lineCount := 0

    for scanner.Scan() {
        line := scanner.Text()
        
        if lineCount == asciiArtHeight {
            // We've read 8 lines, so save this ASCII art and move to the next character
            asciiArtMap[currentChar] = currentArtBuilder.String()
            currentArtBuilder.Reset()
            currentChar++
            lineCount = 0
        }

        if line == "" {
            continue
        }

        currentArtBuilder.WriteString(line + "\n")
        lineCount++
    }

   
    if lineCount == asciiArtHeight {
        asciiArtMap[currentChar] = currentArtBuilder.String()
    }

    if currentChar != '~' {
        return nil, errors.New("unexpected end of file")
    }

    return asciiArtMap, nil
}
