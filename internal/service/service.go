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
	RenderAsciiOutput(w http.ResponseWriter, formattedAscii, font, width, lines string) *models.Error
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

func (s *Service) RenderAsciiPage(w http.ResponseWriter, r *http.Request) *models.Error {
	temp, err := template.New("index.html").ParseFiles("web/templates/index.html")
	if err != nil {
		return &models.Error{
			Err:    "Server Error",
			Detail: err.Error(),
		}
	}

	asciiPageData := struct {
		AsciiPageRoute string
		HomePageRoute  string
	}{
		AsciiPageRoute: "/home/ascii-art",
		HomePageRoute:  "/home",
	}

	err = temp.Execute(w, asciiPageData)
	if err != nil {
		return &models.Error{
			Err:    "Server Error",
			Detail: err.Error(),
		}
	}
	return nil
}
