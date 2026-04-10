package transformer

func (t *Transformer) ConvertCharToAscii(latinWords []string, font string) [][][]string {
	var asciiWords [][][]string

	for _, word := range latinWords {
		currentWord := word
		var asciiWord [][]string

		if currentWord == "" {
			asciiWord = append(asciiWord, []string{""})
			asciiWords = append(asciiWords, asciiWord)
			continue
		}

		for _, char := range currentWord {
			currentChar := char

			asciiChar, err := t.ReadFont(currentChar, font)
			if err != nil {
				return nil
			}

			asciiWord = append(asciiWord, asciiChar)
		}
		asciiWords = append(asciiWords, asciiWord)
	}
	return asciiWords
}
