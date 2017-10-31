package webserver

import (
	"encoding/json"
	"errors"
	"game"
	"net/http"
	"strconv"
	"time"
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

const SUGGEST_DURATION time.Duration = 1

func New(w http.ResponseWriter, r *http.Request) {
	sizeString := r.URL.Query().Get("size")

	size, err := strconv.Atoi(sizeString)
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

	timeout := make(chan bool, 1)

	go func() {
		time.Sleep(time.Second * SUGGEST_DURATION)
		timeout <- true
	}()

	path, err := game.SolvePuzzle(grid, initialGrid, timeout)
	panicOnError(err)

	w.Header().Set("Content-Type", "application/json")

	if len(path) > 0 {
		dir, err := game.DirectionFromCoords(grid, path[0])
		panicOnError(err)
		json.NewEncoder(w).Encode(game.ConvertMoveToMoveString(dir))
		return
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("No suggestion found"))
}
