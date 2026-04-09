package service

import (
	"html/template"
	"net/http"

	"acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
)

type Service struct{}

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
