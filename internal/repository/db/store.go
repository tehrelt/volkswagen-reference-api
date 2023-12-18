package db

import (
	"database/sql"

	"github.com/tehrelt/volkswagen-reference-api/internal/repository"
)

type Store struct {
	db            *sql.DB
	carRepository *CarRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Car() repository.CarRepository {
	if s.carRepository != nil {
		return s.carRepository
	}

	s.carRepository = &CarRepository{
		store: s,
	}

	return s.carRepository
}
