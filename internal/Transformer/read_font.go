package transformer

import (
	"bufio"
	"os"
	"strings"

	"acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
)

func (t *Transformer) ReadFont(char rune, font string) ([]string, *models.Error) {
	path := ""

	switch strings.ToLower(font) {
	case "standard":
		path = "fonts/standard.txt"
	case "shadow":
		path = "fonts/shadow.txt"
	case "tinkertoy":
		path = "fonts/thinker_toy.txt"
	default:
		path = "fonts/standard.txt"
	}

	startLine := ((char - ' ') * 9) + 1
	endLine := startLine + 8
	currentLine := 0

	file, err := os.Open(path)
	if err != nil {
		return nil, &models.Error{
			Err:    "Read error",
			Detail: err.Error(),
		}
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var asciiChar []string

	for scanner.Scan() {
		if currentLine >= int(startLine) && currentLine <= int(endLine) {
			line := scanner.Text()

			asciiChar = append(asciiChar, line)
		} else if currentLine > int(endLine) {
			break
		}
		currentLine++
	}

	return asciiChar, nil
}
