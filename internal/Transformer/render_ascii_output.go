package transformer

import (
	"fmt"
	"html/template"
	"net/http"

	"acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
)

func (t *Transformer) RenderAsciiOutput(w http.ResponseWriter, formattedAscii, font, width, lines string) *models.Error {
	temp, err := template.New("ascii_output.html").ParseFiles("web/static/ascii_output.html")
	if err != nil {
		fmt.Println("1", err.Error())
		return &models.Error{
			Err:    "Unable to load page",
			Detail: err.Error(),
		}
	}

	asciiOutputData := struct {
		AsciiOutput string
		Font        string
		Width       string
		Lines       string
	}{
		AsciiOutput: formattedAscii,
		Font:        font,
		Width:       width,
		Lines:       lines,
	}

	if err = temp.ExecuteTemplate(w, "ascii-output", asciiOutputData); err != nil {
		fmt.Println(err.Error())
		return &models.Error{
			Err:    "Unable to load page",
			Detail: err.Error(),
		}
	}
	return nil
}
