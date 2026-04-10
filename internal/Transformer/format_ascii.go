package transformer

import (
	"slices"
	"strings"
)

func (t *Transformer) AsciiFormatter(asciiWords [][][]string) string {
	var formattedAsciiWord strings.Builder

	for _, asciiWord := range asciiWords {
		if slices.Compare(asciiWord[0], []string{""}) == 0 {
			formattedAsciiWord.WriteString("\n")
			continue
		}
		for i := 0; i < 8; i++ {
			for j := 0; j < len(asciiWord); j++ {
				if i < 8-1 {
					if j != len(asciiWord)-1 {
						formattedAsciiWord.WriteString(asciiWord[j][i])
					} else {
						
					}
				} else {
					formattedAsciiWord.WriteString(asciiWord[j][i])
				}
			}
		}
	}
	return formattedAsciiWord.String()
}
