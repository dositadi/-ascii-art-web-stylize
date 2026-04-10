package transformer

import (
	"html/template"
	"net/http"

	"acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
)

func (t *Transformer) RenderAsciiOutput(w http.ResponseWriter, formattedAscii, font, width, lines string) *models.Error {
	temp, err := template.New("ascii_output.html").ParseFiles("web/static/ascii_output.html")
	if err != nil {
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

	asciiPageData := struct {
		AsciiPageRoute string
	}{
		AsciiPageRoute: "/home/ascii-art",
	}

	if w.Header().Get("HX-Request") == "true" {
		if err = temp.ExecuteTemplate(w, "ascii-output", asciiOutputData); err != nil {
			return &models.Error{
				Err:    "Unable to load page",
				Detail: err.Error(),
			}
		}
	} else {
		err = temp.Execute(w, asciiPageData)
		if err != nil {
			return &models.Error{
				Err:    "Server Error",
				Detail: err.Error(),
			}
		}
	}
	return nil
}
