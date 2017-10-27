package webserver

import (
	"encoding/json"
	"game"
	"net/http"
)

type NewParams struct {
	Size byte
}

type MoveParams struct {
	Grid game.Grid
	Dir  byte
}

type SuggestParams struct {
	Grid       game.Grid
	SolvedGrid game.Grid
}

func New(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var newParams NewParams
	err := decoder.Decode(&newParams)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	grid := game.BuildGrid(newParams.Size)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(grid)
}

func Move(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var moveParams MoveParams
	err := decoder.Decode(&moveParams)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	coords, err := game.CoordsFromDirection(moveParams.Grid, moveParams.Dir)
	if err != nil {
		panic(err)
	}
	grid, err := game.Move(moveParams.Grid, coords)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(grid)
}

func Suggest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var suggestParams SuggestParams
	err := decoder.Decode(&suggestParams)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	path, err := game.DeepPuzzleAlgorithm(suggestParams.Grid, suggestParams.SolvedGrid)

	w.Header().Set("Content-Type", "application/json")
	if len(path) > 0 {
		coords := path[0]
		json.NewEncoder(w).Encode(coords)
	} else {
		panic(err)
	}
}
