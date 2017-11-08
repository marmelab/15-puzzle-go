package webserver

import (
	"encoding/json"
	"errors"
	"game"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

type GridResponse struct {
	InitialGrid [][]int
	Grid        [][]int
}

type MoveParams struct {
	InitialGrid [][]int
	Grid        [][]int
	Move        string
}

type MoveTileParams struct {
	InitialGrid [][]int
	Grid        [][]int
	TileNumber  int
}

type MoveResponse struct {
	Grid      [][]int
	IsVictory bool
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

	initialGrid := game.BuildGrid(byte(size))
	grid, _ := game.Shuffle(game.DeepCopyGrid(initialGrid))

	json.NewEncoder(w).Encode(GridResponse{InitialGrid: game.ConvertGridToGridInt(initialGrid), Grid: game.ConvertGridToGridInt(grid)})
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

	victory := reflect.DeepEqual(moveParams.Grid, moveParams.InitialGrid)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(MoveResponse{Grid: game.ConvertGridToGridInt(grid), IsVictory: victory})
}

func MoveTile(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var moveTileParams MoveTileParams
	err := decoder.Decode(&moveTileParams)
	panicOnError(err)
	defer r.Body.Close()

	coords, err := game.FindTileByValue(game.ConvertGridIntToGrid(moveTileParams.Grid), byte(moveTileParams.TileNumber))
	panicOnError(err)

	grid, err := game.Move(game.ConvertGridIntToGrid(moveTileParams.Grid), coords)
	panicOnError(err)

	gridInt := game.ConvertGridToGridInt(grid)
	victory := reflect.DeepEqual(gridInt, moveTileParams.InitialGrid)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(MoveResponse{Grid: gridInt, IsVictory: victory})
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
