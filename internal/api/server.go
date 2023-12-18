package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tehrelt/volkswagen-reference-api/internal/repository"
)

type server struct {
	store  repository.Store
	router *mux.Router
}

func newServer(store repository.Store) *server {
	s := &server{
		store:  store,
		router: mux.NewRouter(),
	}

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
