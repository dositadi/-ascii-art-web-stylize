package service

import (
	"fmt"
	"net/http"

	"acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
)

func (s *Service) TransformAscii(w http.ResponseWriter, r *http.Request) *models.Error {
	//text := r.FormValue("text")
	//font := r.FormValue("font")

	latinWords, err := s.AsciiTransformer.SplitText("Hello")
	if err != nil {
		return err
	}

	asciiWords := s.AsciiTransformer.ConvertCharToAscii(latinWords, "tinkertoy")

	formattedAscii := s.AsciiTransformer.AsciiFormatter(asciiWords)

	fmt.Println(formattedAscii)

	err = s.AsciiTransformer.RenderAsciiOutput(w, formattedAscii)
	if err != nil {
		return err
	}
	return nil
}
