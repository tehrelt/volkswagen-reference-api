package db

import "github.com/tehrelt/volkswagen-reference-api/internal/models"

type CarRepository struct {
	store *Store
}

func (r *CarRepository) Create(car *models.Car) error {
	return r.store.db.QueryRow(
		"INSERT INTO car (model, release_year, description) VALUES ($1, $2, $3) RETURNING id",
		car.Model,
		car.ReleaseYear,
		car.Description.String,
	).Scan(&car.ID)
}

func (r *CarRepository) Find(model string) (*models.Car, error) {

	car := &models.Car{}

	if err := r.store.db.QueryRow(
		"SELECT id, model, description FROM car WHERE model = $1",
		model,
	).Scan(&car.ID, &car.Model, &car.Description); err != nil {
		return nil, err
	}

	return car, nil
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

	rows, err := r.store.db.Query("SELECT id, model, description FROM car")
	if err != nil {
		return nil, err
	}

	var t models.Car
	var dto models.CarDto
	for rows.Next() {
		rows.Scan(&t.ID, &t.Model, &t.Description)

		dto = models.CarDto{
			ID:    t.ID,
			Model: t.Model,
		}

		if t.Description.Valid {
			dto.Description = t.Description.String
		} else {
			dto.Description = ""
		}

		cars = append(cars, dto)
	}

	return cars, nil
}

func (r *CarRepository) Get(id int) (*models.Car, error) {
	car := &models.Car{}

	if err := r.store.db.QueryRow(
		"SELECT id, model, description FROM car WHERE id = $1",
		id,
	).Scan(&car.ID, &car.Model, &car.Description); err != nil {
		return nil, err
	}

	return car, nil
}
