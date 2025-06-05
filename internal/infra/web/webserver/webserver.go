package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	return &WebServer{
		Router:        router,
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(method, path string, handler http.HandlerFunc) {
	switch method {
	case "GET":
		s.Router.Get(path, handler)
	case "POST":
		s.Router.Post(path, handler)
	case "PUT":
		s.Router.Put(path, handler)
	case "DELETE":
		s.Router.Delete(path, handler)
	}
}

func (s *WebServer) Start() {
	http.ListenAndServe(s.WebServerPort, s.Router)
}
