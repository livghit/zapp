package server

import (
	"log"
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
	hub := newHub()
	router.Handle("/*", spaHandler)
	// ws handler
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
}

// serveWs handles websocket requests from the peer.
func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}
