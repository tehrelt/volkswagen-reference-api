package api

import (
	"database/sql"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/tehrelt/volkswagen-reference-api/internal/middleware"
	"github.com/tehrelt/volkswagen-reference-api/internal/models"
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
	s.router.Use(middleware.CommonMiddleware)
	s.router.Use(middleware.SetRequestId)
	s.router.Use(middleware.LogMiddleware(s.logger))
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))

	car := s.router.PathPrefix("/cars").Subrouter()
	car.HandleFunc("/", s.handleCreateCar()).Methods("POST")
}

func (s *server) handleCreateCar() http.HandlerFunc {
	type request struct {
		Model       string `json:"model"`
		ReleaseYear int    `json:"release_year"`
		Description string `json:"description,omitempty"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.logger.Debug("cant decode a body of request", slog.Any("err", err))
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		c := &models.Car{
			Model:       req.Model,
			ReleaseYear: req.ReleaseYear,
			Description: sql.NullString{
				String: req.Description,
			},
		}

		if err := s.store.Car().Create(c); err != nil {
			s.logger.Debug("cant create an alias", slog.Any("err", err))
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusOK, c)
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
