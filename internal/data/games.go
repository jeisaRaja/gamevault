package data

import (
	"fmt"
	"gamevault/internal/validator"
	"time"
)

type Game struct {
	ID          int64                `json:"id"`
	CreatedAt   time.Time            `json:"-"`
	Title       string               `json:"title"`
	Year        int32                `json:"year,omitempty"`
	Genres      []validator.Genres   `json:"genres,omitempty"`
	Platform    []validator.Platform `json:"platform,omitempty"`
	Developer   string               `json:"developer,omitempty"`
	Publisher   string               `json:"publisher,omitempty"`
	Rating      float32              `json:"rating,omitempty"`
	RatingCount int32                `json:"rating_count"`
	Price       int64                `json:"price,omitempty"`
}

func ValidateGame(v *validator.Validator, game *Game) {
	v.Check(game.Title != "", "title", "must be provided")
	v.Check(len(game.Title) <= 500, "title", "must less than 500 bytes long")

	v.Check(game.Year != 0, "year", "must be provided")
	v.Check(game.Year >= 1900, "year", "must be greater than 1900")
	v.Check(game.Year <= int32(time.Now().Year()), "year", fmt.Sprintf("must be before %d", time.Now().Year()))

	v.Check(game.Price >= 0, "price", "must be more than $0")

	v.Check(game.Genres != nil, "genres", "must be provided")
	v.Check(len(game.Genres) >= 1, "genres", "must be provided")
	v.Check(len(game.Genres) <= 10, "genres", "must have less or equal than 10 genres")
	v.Check(validator.ValidateGameGenres(game.Genres), "genres", "must be member of valid genre")
	v.Check(validator.Unique(validator.EnumToStringSlice(game.Genres)), "genres", "must not contain duplicate values")

	v.Check(game.Platform != nil, "genres", "must be provided")
	v.Check(len(game.Platform) >= 1, "genres", "must be provided")
	v.Check(len(game.Platform) <= 10, "genres", "must have less or equal than 10 genres")
	v.Check(validator.Unique(validator.EnumToStringSlice(game.Platform)), "genres", "must not contain duplicate values")

	v.Check(game.Publisher != "", "genres", "must be provided")
}
