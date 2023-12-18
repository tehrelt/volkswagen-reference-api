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
		"SELECT id, model, description FROM cars WHERE model = $1",
		model,
	).Scan(&car.ID, &car.Model, &car.Description); err != nil {
		return nil, err
	}

	return car, nil
}

func (r *CarRepository) Delete(id int) error {
	stmt, err := r.store.db.Prepare("DELETE FROM cars WHERE id = ?")
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(id); err != nil {
		return err
	}

	return nil
}

func (r *CarRepository) GetAll() ([]models.Car, error) {
	var cars []models.Car

	rows, err := r.store.db.Query("SELECT id, model, description FROM cars")
	if err != nil {
		return nil, err
	}

	var t models.Car
	for rows.Next() {
		rows.Scan(&t.ID, &t.Model, &t.Description)

		cars = append(cars, t)
	}

	return cars, nil
}

func (r *CarRepository) Get(id int) (*models.Car, error) {
	car := &models.Car{}

	if err := r.store.db.QueryRow(
		"SELECT id, model, description FROM cars WHERE id = $1",
		id,
	).Scan(&car.ID, &car.Model, &car.Description); err != nil {
		return nil, err
	}

	return car, nil
}
