package data

import "time"

type Game struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Platform  []string  `json:"platform,omitempty"`
	Developer string    `json:"developer,omitempty"`
	Publisher string    `json:"publisher,omitempty"`
	Rating    float32   `json:"rating,omitempty"`
	Price     int64     `json:"price,omitempty"`
}
