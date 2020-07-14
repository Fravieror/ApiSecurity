package server

import (
	"net/http"

	"apiSecurity/token/infrastructure"

	"github.com/gorilla/mux"
)

type api struct {
	router http.Handler
}

type Server interface {
	Router() http.Handler
}

func New() Server {
	a := &api{}

	r := mux.NewRouter()
	r.HandleFunc("/token", infrastructure.GenerateToken).Methods(http.MethodPost)
	r.HandleFunc("/verify", infrastructure.VerifyToken).Methods(http.MethodPost)
	// r.HandleFunc("/ping", infrastructure.Ping).Methods(http.MethodGet)
	/* Gracias a Gorilla podemos usar expresiones regulares para asegurarnos
	 de antemano que los parámetros pasados cumplen con la regla que queremos.
	r.HandleFunc("/gophers/{ID:[a-zA-Z0-9_]+}", a.fetchGopher).Methods(http.MethodGet)*/
	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}
