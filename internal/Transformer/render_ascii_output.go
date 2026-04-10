package transformer

import (
	"html/template"
	"net/http"

	"acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
)

func (t *Transformer) RenderAsciiOutput(w http.ResponseWriter, formattedAscii string) *models.Error {
	temp, err := template.New("ascii_output.html").ParseFiles("web/static/ascii_output.html")
	if err != nil {
		return &models.Error{
			Err:    "Unable to load page",
			Detail: err.Error(),
		}
	}

	asciiOutputData := struct {
		FormattedAsciiText string
	}{}

	if err = temp.ExecuteTemplate(w, "ascii-output", asciiOutputData); err != nil {
		return &models.Error{
			Err:    "Unable to load page",
			Detail: err.Error(),
		}
	}
	return nil
}
