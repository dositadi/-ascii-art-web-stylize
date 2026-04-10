package transformer

import (
	"bufio"
	"strings"

	"acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
)

func (t *Transformer) SplitText(text string) ([]string, *models.Error) {
	scanner := bufio.NewScanner(strings.NewReader(text))
	var latinWords []string

	for scanner.Scan() {
		latinWord := scanner.Text()
		latinWords = append(latinWords, latinWord)
	}
	return latinWords, nil
}
