package main

import (
	"fmt"
	"net/http"
)

func (app *application) createGameHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new game")
}

func (app *application) showGameHandler(w http.ResponseWriter, r *http.Request) {
	game_id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show details of game %d\n", game_id)
}
