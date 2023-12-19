package db

import (
	"log"

	"github.com/tehrelt/volkswagen-reference-api/internal/models"
)

type CarRepository struct {
	store *Store
}

func (r *CarRepository) Create(car *models.CarDto) error {
	// log.Print(car);
	return r.store.db.QueryRow(
		"INSERT INTO car (model, release_year, description, image) VALUES ($1, $2, $3, $4) RETURNING id",
		car.Model,
		car.ReleaseYear,
		car.Description,
		car.ImageLink,
	).Scan(&car.ID)
}

func (r *CarRepository) Find(model string) (*models.CarDto, error) {

	car := &models.Car{}

	if err := r.store.db.QueryRow(
		"SELECT id, model, description, image FROM car WHERE model = $1",
		model,
	).Scan(&car.ID, &car.Model, &car.Description, &car.ImageLink); err != nil {
		return nil, err
	}
	dto := &models.CarDto{
		ID:          car.ID,
		Model:       car.Model,
		ReleaseYear: car.ReleaseYear,
		Description: car.Description.String,
		ImageLink:   car.ImageLink.String,
	}
	return dto, nil
}

func (r *CarRepository) Delete(id int) error {
	stmt, err := r.store.db.Prepare("DELETE FROM car WHERE id = ?")
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(id); err != nil {
		return err
	}

	return nil
}

func (r *CarRepository) GetAll() ([]models.CarDto, error) {
	var cars []models.CarDto

	rows, err := r.store.db.Query("SELECT id, model, release_year, description, image FROM car")
	if err != nil {
		return nil, err
	}

	var t models.Car
	var dto models.CarDto
	for rows.Next() {
		rows.Scan(&t.ID, &t.Model, &t.ReleaseYear, &t.Description, &t.ImageLink)

		dto = models.CarDto{
			ID:          t.ID,
			ReleaseYear: t.ReleaseYear,
			Model:       t.Model,
			Description: t.Description.String,
			ImageLink:   t.ImageLink.String,
		}

		cars = append(cars, dto)
	}

	return cars, nil
}

func (r *CarRepository) Get(id int) (*models.CarDto, error) {
	car := &models.Car{}

	if err := r.store.db.QueryRow(
		"SELECT id, model, release_year, description, image FROM car WHERE id = $1",
		id,
	).Scan(&car.ID, &car.Model, &car.ReleaseYear, &car.Description, &car.ImageLink); err != nil {
		return nil, err
	}

	log.Print(car)

	dto := &models.CarDto{
		ID:          car.ID,
		Model:       car.Model,
		ReleaseYear: car.ReleaseYear,
		Description: car.Description.String,
		ImageLink:   car.ImageLink.String,
	}

	return dto, nil
}
