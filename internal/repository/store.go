package repository

import "github.com/tehrelt/volkswagen-reference-api/internal/models"

type Store interface {
	Car() CarRepository
}

type CarRepository interface {
	Create(car *models.CarDto) error
	Find(modelName string) (*models.CarDto, error)
	Delete(id int) error
	GetAll() ([]models.CarDto, error)
	Get(id int) (*models.CarDto, error)
}
