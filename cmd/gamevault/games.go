package main

import (
	"gamevault/internal/data"
	"net/http"
	"time"
)

func (app *application) createGameHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title     string   `json:"title"`
		Year      int32    `json:"year,omitempty"`
		Genres    []string `json:"genres,omitempty"`
		Platform  []string `json:"platform,omitempty"`
		Developer string   `json:"developer,omitempty"`
		Publisher string   `json:"publisher,omitempty"`
		Price     int64    `json:"price,omitempty"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// fmt.Fprintf(w, "%+v\n", input)
	app.writeJSON(w, http.StatusOK, envelope{"game": input}, nil)
}

func (app *application) showGameHandler(w http.ResponseWriter, r *http.Request) {
	game_id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	game := data.Game{
		Title:     "Soma",
		Year:      2021,
		CreatedAt: time.Now(),
		Genres:    []string{"action", "fantasy"},
		Developer: "ubisoft",
		Publisher: "ubisoft",
		ID:        game_id,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"game": game}, nil)
	if err != nil {
		app.logger.Println(err)
		app.serverErrorResponse(w, r, err)
	}
}
