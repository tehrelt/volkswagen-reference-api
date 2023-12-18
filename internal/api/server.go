package api

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/tehrelt/volkswagen-reference-api/internal/middleware"
	"github.com/tehrelt/volkswagen-reference-api/internal/repository"
)

type server struct {
	store  repository.Store
	logger *slog.Logger
	router *mux.Router
}

func newServer(store repository.Store) *server {
	s := &server{
		store:  store,
		logger: slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})),
		router: mux.NewRouter(),
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	
	s.router.Use(middleware.SetRequestId)
	s.router.Use(middleware.LogMiddleware(s.logger))
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))

}
