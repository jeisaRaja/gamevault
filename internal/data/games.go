package data

import (
	"database/sql"
	"fmt"
	"gamevault/internal/validator"
	"time"

	"github.com/lib/pq"
)

type Game struct {
	ID          int64      `json:"id"`
	CreatedAt   time.Time  `json:"-"`
	Title       string     `json:"title"`
	Year        int32      `json:"year"`
	Genres      []Genres   `json:"genres"`
	Platforms    []Platform `json:"platform"`
	Developer   string     `json:"developer"`
	Publisher   string     `json:"publisher"`
	Rating      float32    `json:"rating"`
	RatingCount int32      `json:"rating_count"`
	Price       int64      `json:"price"`
}

func ValidateGame(v *validator.Validator, game *Game) {
	v.Check(game.Title != "", "title", "must be provided")
	v.Check(len(game.Title) <= 500, "title", "must less than 500 bytes long")

	v.Check(game.Year != 0, "year", "must be provided")
	v.Check(game.Year >= 1900, "year", "must be later than 1900")
	v.Check(game.Year <= int32(time.Now().Year()), "year", fmt.Sprintf("must be before %d", time.Now().Year()))

	v.Check(game.Price >= 0, "price", "must be more than $0")

	v.Check(game.Genres != nil, "genres", "must be provided")
	v.Check(len(game.Genres) >= 1, "genres", "must be provided")
	v.Check(len(game.Genres) <= 10, "genres", "must have less or equal than 10 genres")
	v.Check(ValidateGameGenres(game.Genres), "genres", "must be member of valid genre")
	v.Check(validator.Unique(EnumToStringSlice(game.Genres)), "genres", "must not contain duplicate values")

	v.Check(game.Platforms != nil, "platforms", "must be provided")
	v.Check(len(game.Platforms) >= 1, "platforms", "must be provided")
	v.Check(ValidateGamePlatforms(game.Platforms), "platforms", "must be member of valid platforms")
	v.Check(validator.Unique(EnumToStringSlice(game.Platforms)), "platforms", "must not contain duplicate values")

	v.Check(game.Publisher != "", "genres", "must be provided")
}

type GameModel struct {
	DB *sql.DB
}

func (m GameModel) Insert(game *Game) error {
	query := `
    INSERT INTO games (title, year, genres, platform, developer, publisher, price) 
    VALUES ($1, $2, $3, $4, $5, $6, $7)
    RETURNING id, created_at
`
	args := []interface{}{game.Title, game.Year, pq.Array(game.Genres), pq.Array(game.Platforms), game.Developer, game.Publisher, game.Price}

	return m.DB.QueryRow(query, args...).Scan(&game.ID, &game.CreatedAt)
}

func (m GameModel) Get(id int64) (*Game, error) {
	return nil, nil
}

func (m GameModel) Update(game *Game) error {
	return nil
}

func (m GameModel) Delete(id int64) error {
	return nil
}

// for unit test
type MockGameModel struct{}

func (m MockGameModel) Insert(game *Game) error {
	return nil
}

func (m MockGameModel) Get(id int64) (*Game, error) {
	return nil, nil
}

func (m MockGameModel) Update(game *Game) error {
	return nil
}

func (m MockGameModel) Delete(id int64) error {
	return nil
}
