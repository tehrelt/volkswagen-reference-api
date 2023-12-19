package models

import "database/sql"

type Car struct {
	ID          int            `json:"id"`
	Model       string         `json:"model"`
	ReleaseYear int            `json:"release_year"`
	Description sql.NullString `json:"description"`
	ImageLink   sql.NullString `json:"image_link"`
}

type CarDto struct {
	ID          int    `json:"id"`
	Model       string `json:"model"`
	ReleaseYear int    `json:"release_year"`
	Description string `json:"description"`
	ImageLink   string `json:"image_link"`
}
