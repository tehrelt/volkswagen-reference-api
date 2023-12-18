package repository

import "github.com/tehrelt/volkswagen-reference-api/internal/models"

type Store interface {
	Car() CarRepository
}

type CarRepository interface {
	Create(car *models.Car) error
	Find(modelName string) (*models.Car, error)
	Delete(id int) error
	GetAll() ([]models.Car, error)
	Get(id int) (*models.Car, error)
}
