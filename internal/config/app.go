package config

import (
	"fmt"
	"log"
	"net/http"

	"acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/internal/handler"
	"acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/internal/service"
)

type App struct {
	Router *http.ServeMux
}

func (a *App) InitializeRoute() {
	service := service.NewService()
	handler := handler.NewHandler(&service)

	a.Router = http.NewServeMux()
	a.Router.HandleFunc("GET /home", handler.HomeHandler)
	a.Router.HandleFunc("GET /home/ascii-art", handler.AsciiPageHandler)
}

func (a *App) InitializeFileServer() {
	pageCSSFileServer := http.FileServer(http.Dir("web/static"))
	a.Router.Handle("GET /web/static/", http.StripPrefix("/web/static/", pageCSSFileServer))

	asciiPageFileServer := http.FileServer(http.Dir("web/templates"))
	a.Router.Handle("GET /web/templates/", http.StripPrefix("/web/templates/", asciiPageFileServer))
}

func (a *App) Run() {
	a.InitializeRoute()
	a.InitializeFileServer()
	server := http.Server{
		Addr:    ":8081",
		Handler: a.Router,
	}
	fmt.Println("Server running at port 8082")

	log.Fatal(server.ListenAndServe())
}
