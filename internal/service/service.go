package service

import (
	"html/template"
	"net/http"

	transformer "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/internal/Transformer"
	"acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
)

type Service struct {
	AsciiTransformer TransformAsciiText
}

type TransformAsciiText interface {
	SplitText(text string) ([]string, *models.Error)
	ConvertCharToAscii(latinWords []string, font string) [][][]string
	AsciiFormatter(asciiWords [][][]string) string
	ReadFont(char rune, font string) ([]string, *models.Error)
	RenderAsciiOutput(w http.ResponseWriter, formattedAscii string) *models.Error
}

func NewService() Service {
	return Service{
		AsciiTransformer: &transformer.Transformer{},
	}
}

func (s *Service) RenderHomePage(w http.ResponseWriter, r *http.Request) *models.Error {
	temp, err := template.New("index.html").ParseFiles("web/static/index.html")
	if err != nil {
		return &models.Error{
			Err:    "Server Error",
			Detail: err.Error(),
		}
	}

	homePageData := struct {
		AsciiPageRoute string
	}{
		AsciiPageRoute: "/home/ascii-art",
	}

	err = temp.Execute(w, homePageData)
	if err != nil {
		return &models.Error{
			Err:    "Server Error",
			Detail: err.Error(),
		}
	}
	return nil
}
