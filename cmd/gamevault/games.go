package main

import (
	"errors"
	"fmt"
	"gamevault/internal/data"
	"gamevault/internal/validator"
	"net/http"
)

type GameResponse struct {
	Games data.Game `json:"games"`
}

type GameCreateRequest struct {
	Title     string   `json:"title" validate:"required,min=3,max=100"`
	Year      int32    `json:"year" validate:"required,gte=1900"`
	Genres    []string `json:"genres" validate:"required,dive,required,max=10"`
	Platforms []string `json:"platform" validate:"dive,required,min:1" `
	Developer string   `json:"developer" validate:"required"`
	Publisher string   `json:"publisher" validate:"required"`
	Price     int64    `json:"price" validate:"gte=0"`
}

// @summary Create a new game entry
// @description This endpoint allows the user to create a new game in the system by providing necessary game details.
// @tags Game
// @accept json
// @produce json
// @param game body GameCreateRequest true "Game details"
// @success 200 {object} data.Game "Game created successfully"
// @router /games/ [post]
func (app *application) createGameHandler(w http.ResponseWriter, r *http.Request) {
	var input GameCreateRequest
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	game := &data.Game{
		Title:     input.Title,
		Year:      input.Year,
		Genres:    input.Genres,
		Platforms: input.Platforms,
		Developer: input.Developer,
		Publisher: input.Publisher,
		Price:     input.Price,
	}

	v := validator.New()
	if data.ValidateGame(v, game); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Games.Insert(game)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/movies/%d", game.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"game": game}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// @summary Show details of a specific game
// @description This endpoint retrieves the details of a specific game by its ID.
// @tags Game
// @produce json
// @param id path int true "Game ID"
// @success 200 {object} GameResponse "Game details fetched successfully"
// @failure 404 {object} object "Game not found"
// @router /games/{id} [get]
func (app *application) showGameHandler(w http.ResponseWriter, r *http.Request) {
	game_id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	game, err := app.models.Games.Get(game_id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"game": game}, nil)
	if err != nil {
		app.logger.Println(err)
		app.serverErrorResponse(w, r, err)
	}
}
