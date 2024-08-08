package functions

import "strings"

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
