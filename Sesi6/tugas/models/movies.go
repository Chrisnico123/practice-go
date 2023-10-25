package models

type Movies struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Genre       string `json:"genre" db:"genre"`
}