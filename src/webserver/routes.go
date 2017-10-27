package webserver

import (
	"encoding/json"
	"game"
	"strconv"
	"fmt"
	"net/http"
	"errors"
)

type GridResponse struct {
	Grid [][]int
}

type MoveParams struct {
	Grid [][]int
	Move string
}

type SuggestParams struct {
	Grid [][]int
	SolvedGrid [][]int
}

type SuggestResponse struct {
	Move string
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func New(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Route 'new' called")
	
	queryString := r.URL.Query().Get("size")
	
	size, err :=strconv.Atoi(queryString)
	if err != nil {
		panicOnError(err)
	} else if size < 1 || size >= 10 {
		panicOnError(errors.New("The puzzle size must be between 2 and 10"))
	}

	grid := game.BuildGrid(byte(size))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(GridResponse{Grid: game.ConvertGridToGridInt(grid)})
}

func Move(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Route 'move' called")
	
	decoder := json.NewDecoder(r.Body)

	var moveParams MoveParams
	err := decoder.Decode(&moveParams)
	panicOnError(err)
	defer r.Body.Close()

	coords, err := game.CoordsFromDirection(game.ConvertGridIntToGrid(moveParams.Grid), game.ConvertMoveStringToMove(moveParams.Move))
	panicOnError(err)

	grid, err := game.Move(game.ConvertGridIntToGrid(moveParams.Grid), coords)
	panicOnError(err)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(GridResponse{Grid: game.ConvertGridToGridInt(grid)})
}

func Suggest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Route 'suggest' called")
	
	decoder := json.NewDecoder(r.Body)

	var suggestParams SuggestParams
	err := decoder.Decode(&suggestParams)
	panicOnError(err)

	defer r.Body.Close()
	
	grid := game.ConvertGridIntToGrid(suggestParams.Grid)
	solvedGrid := game.ConvertGridIntToGrid(suggestParams.SolvedGrid)
	path, err := game.DeepPuzzleAlgorithm(grid, solvedGrid)

	w.Header().Set("Content-Type", "application/json")

	if len(path) > 0 {
		dir, err := game.DirectionFromCoords(grid, path[0])
		panicOnError(err)
		json.NewEncoder(w).Encode(game.ConvertMoveToMoveString(dir))
	} else {
		panicOnError(errors.New("No suggestion found"))
	}
}
