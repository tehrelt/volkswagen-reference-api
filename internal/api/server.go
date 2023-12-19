package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"strconv"

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
	car.HandleFunc("/", s.handleGetCars()).Methods("GET")
	car.HandleFunc("/{id:[0-9]+}", s.handleGetCar()).Methods("GET")
	car.HandleFunc("/{id:[0-9]+}", s.handleDeleteCar()).Methods("DELETE")
	// car.HandleFunc("/{id:[0-9]+}", s.handleUpdateCar()).Methods("UPDATE")
}

func (s *server) handleCreateCar() http.HandlerFunc {
	type request struct {
		Model       string `json:"model"`
		ReleaseYear int    `json:"release_year"`
		Description string `json:"description,omitempty"`
		ImageLink   string `json:"image_link,omitempty"`
		Bodywork    string `json:"bodywork"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.logger.Debug("cant decode a body of request", slog.Any("err", err))
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		c := &models.CarDto{
			Model:       req.Model,
			ReleaseYear: req.ReleaseYear,
			Description: req.Description,
			ImageLink:   req.ImageLink,
			Bodywork:    req.Bodywork,
		}

		if err := s.store.Car().Create(c); err != nil {
			s.logger.Debug("cant create an car", slog.Any("err", err))
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusOK, c)
	}
}

func (s *server) handleGetCars() http.HandlerFunc {

	type response struct {
		Data  []models.CarOverview `json:"data"`
		Count int                  `json:"count"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		query := r.URL.Query().Get("query")

		data, err := s.store.Car().GetAll()
		if err != nil {
			s.logger.Debug("unexpected error on getting all cars", slog.Any("err", err))
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		re := response{
			Data:  data,
			Count: len(data),
		}
		s.logger.Debug("get a set of cars", slog.Int("length", re.Count), slog.String("query", query))
		s.respond(w, r, http.StatusOK, re)
	}
}

func (s *server) handleGetCar() http.HandlerFunc {
	type response struct {
		Data *models.CarDto `json:"data"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			s.logger.Debug("error when converting id", slog.String("value", (mux.Vars(r)["id"])))
			return
		}

		data, err := s.store.Car().Get(id)
		if err != nil {
			s.logger.Debug("unexpected error on getting all aliases", slog.Any("err", err))
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		re := response{
			Data: data,
		}

		s.logger.Debug("get a car", slog.String("model", re.Data.Model))
		s.respond(w, r, http.StatusOK, re)
	}
}

func (s *server) handleDeleteCar() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			s.logger.Debug("error when converting id", slog.String("value", (mux.Vars(r)["id"])))
			return
		}

		if err := s.store.Car().Delete(id); err != nil {
			s.logger.Debug("unexpected error on getting all aliases", slog.Any("err", err))
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.logger.Debug("successfully deleted a car", slog.Int("car_id", id))
		s.respond(w, r, http.StatusOK, id)
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
