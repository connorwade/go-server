package db

import "github.com/google/uuid"

type Music struct {
	ID     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	Artist string    `json:"artist"`
	Price  float64   `json:"price"`
}
