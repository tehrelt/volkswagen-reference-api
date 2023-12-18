package models

import "database/sql"

type Car struct {
	ID          int            `json:"id"`
	Model       string         `json:"model"`
	ReleaseYear int            `json:"release_year"`
	Description sql.NullString `json:"description"`
}
