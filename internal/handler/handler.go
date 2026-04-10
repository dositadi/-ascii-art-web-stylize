package handler

import (
	"net/http"

	"acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
)

type Handler struct {
	Service HandlerService
}

type HandlerService interface {
	RenderHomePage(w http.ResponseWriter, r *http.Request) *models.Error
	RenderAsciiPage(w http.ResponseWriter, r *http.Request) *models.Error
	TransformAscii(w http.ResponseWriter, r *http.Request) *models.Error
}

func NewHandler(h HandlerService) Handler {
	return Handler{
		Service: h,
	}

}

func (h *Handler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	if err := h.Service.RenderHomePage(w, r); err != nil {
		http.Error(w, err.Detail, http.StatusInternalServerError)
		return
	}
}

func (h *Handler) AsciiPageHandler(w http.ResponseWriter, r *http.Request) {
	if err := h.Service.RenderAsciiPage(w, r); err != nil {
		http.Error(w, err.Detail, http.StatusInternalServerError)
		return
	}
}

func (h *Handler) LatinTransformToAscii(w http.ResponseWriter, r *http.Request) {
	if err := h.Service.TransformAscii(w, r); err != nil {
		http.Error(w, err.Detail, http.StatusInternalServerError)
		return
	}
}
