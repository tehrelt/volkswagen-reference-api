package api

import "github.com/gorilla/mux"

type server struct {
	store  repository.Store
	router *mux.Router
}
