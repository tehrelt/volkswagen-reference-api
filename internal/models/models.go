package models

import "database/sql"

type Car struct {
	ID          int            `json:"id"`
	Model       string         `json:"model"`
	ReleaseYear int            `json:"release_year"`
	Description sql.NullString `json:"description"`
	Bodywork    string         `json:"bodywork"`
	ImageLink   sql.NullString `json:"image_link"`
}

type CarDto struct {
	ID          int    `json:"id"`
	Model       string `json:"model"`
	ReleaseYear int    `json:"release_year"`
	Description string `json:"description"`
	Bodywork    string `json:"bodywork"`
	ImageLink   string `json:"image_link"`
}

type CarOverview struct {
	Id    int    `json:"id"`
	Model string `json:"model"`
	Thumb string `json:"image_link"`
}
