package webserver

import (
	"encoding/json"
	"errors"
	"game"
	"net/http"
	"strconv"
)

type GridResponse struct {
	Grid [][]int
}

type MoveParams struct {
	Grid [][]int
	Move string
}

type SuggestParams struct {
	Grid       [][]int
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
	queryString := r.URL.Query().Get("size")

	size, err := strconv.Atoi(queryString)
	if err != nil {
		panicOnError(err)
	} else if size < 1 || size >= 10 {
		panicOnError(errors.New("The puzzle size must be between 2 and 10"))
	}

	grid := game.BuildGrid(byte(size))

	json.NewEncoder(w).Encode(GridResponse{Grid: game.ConvertGridToGridInt(grid)})
}

func Move(w http.ResponseWriter, r *http.Request) {
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
	gridJson := r.URL.Query().Get("grid")
	initialGridJson := r.URL.Query().Get("initial_grid")

	var gridInt [][]int
	var initialGridInt [][]int
	err := json.Unmarshal([]byte(gridJson), &gridInt)
	err = json.Unmarshal([]byte(initialGridJson), &initialGridInt)
	panicOnError(err)

	grid := game.ConvertGridIntToGrid(gridInt)
	initialGrid := game.ConvertGridIntToGrid(initialGridInt)
	path, err := game.SolvePuzzle(grid, initialGrid)
	panicOnError(err)

	w.Header().Set("Content-Type", "application/json")

	if len(path) > 0 {
		dir, err := game.DirectionFromCoords(grid, path[0])
		panicOnError(err)
		json.NewEncoder(w).Encode(game.ConvertMoveToMoveString(dir))
	} else {
		panicOnError(errors.New("No suggestion found"))
	}
}
