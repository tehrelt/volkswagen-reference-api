package api

import (
	"github.com/gorilla/mux"
	"github.com/tehrelt/volkswagen-reference-api/internal/repository"
)

type server struct {
	store  repository.Store
	router *mux.Router
}
