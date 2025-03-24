package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/livghit/zapp/ui"
)

type Server struct {
	Core   *http.Server
	Router *chi.Mux
	// later DB and Cache
}

func CreateNewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	s.Router.Use(middleware.Logger)
	addRoutes(s.Router)

	s.Core = &http.Server{Addr: ":3000", Handler: s.Router}
	return s
}

func addRoutes(router *chi.Mux) {
	spaHandler, err := ui.SpaHandler()
	if err != nil {
		panic(err)
	}
	router.Handle("/*", spaHandler)
}
