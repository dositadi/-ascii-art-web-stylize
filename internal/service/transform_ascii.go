package service

import (
	"net/http"
	"slices"
	"strconv"

	"acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
)

func (s *Service) TransformAscii(w http.ResponseWriter, r *http.Request) *models.Error {
	text := r.FormValue("text")
	font := r.FormValue("font")

	latinWords, err := s.AsciiTransformer.SplitText(text)
	if err != nil {
		return err
	}

	asciiWords := s.AsciiTransformer.ConvertCharToAscii(latinWords, font)

	width := 0
	line := 0

	for _, asciiWord := range asciiWords {
		if len(asciiWord) > width {
			width = len(asciiWord)
		}

		if slices.Compare(asciiWord[0], []string{""}) == 0 {
			line++
		}
	}

	formattedAscii := s.AsciiTransformer.AsciiFormatter(asciiWords)

	err = s.AsciiTransformer.RenderAsciiOutput(w, formattedAscii, font, strconv.Itoa(width), strconv.Itoa(line))
	if err != nil {
		return err
	}
	return nil
}
