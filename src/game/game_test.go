package game

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBuild(t *testing.T) {
	grid := BuildGrid(3)

	expectedGrid := Grid{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}

	if !reflect.DeepEqual(grid, expectedGrid) {
		t.Error("The grid is not built as expected")
	}
}

func TestListMovableTiles(t *testing.T) {
	grid := Grid{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}
	list, err := ListMovableTiles(grid)
	expectedValue := grid[list[0].Y][list[0].X]
	expectedValue2 := grid[list[1].Y][list[1].X]
	if expectedValue != 6 || expectedValue2 != 8 {
		t.Error(fmt.Sprintf("The movable tiles should be 6 and 8 and not %d and %d", expectedValue, expectedValue2))
	}
	if err != nil {
		t.Error("The ListMovableTiles function should not return an error.")
	}
}

func TestListMovableTiles2(t *testing.T) {
	grid := Grid{
		{1, 2, 3},
		{4, 5, 0},
		{7, 8, 6},
	}
	list, err := ListMovableTiles(grid)
	expectedValue := grid[list[0].Y][list[0].X]
	expectedValue2 := grid[list[1].Y][list[1].X]
	expectedValue3 := grid[list[2].Y][list[2].X]
	if expectedValue != 3 || expectedValue2 != 6 || expectedValue3 != 5 {
		t.Error(fmt.Sprintf("The movable tiles should be 3, 6 and 5 and not %d, %d and %d", expectedValue, expectedValue2, expectedValue3))
	}
	if err != nil {
		t.Error("The ListMovableTiles function should not return an error.")
	}
}

func TestListMovableTilesError(t *testing.T) {
	grid := Grid{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	_, err := ListMovableTiles(grid)
	if err == nil {
		t.Error("The ListMovableTiles function should return an error because there is no movable tile in the grid")
	}
}

func TestListMovableTilesWithoutGoingBack(t *testing.T) {
	grid := Grid{
		{1, 2, 3},
		{4, 5, 0},
		{7, 8, 6},
	}

	list, err := ListMovableTilesWithoutGoingBack(grid, 5)

	if len(list) != 2 {
		t.Error(fmt.Sprintf("The movable tiles should return only two values"))
	}
	if err != nil {
		t.Error("The ListMovableTilesWithoutGoingBack function should not return an error.")
	}
}

func TestCoordsFromDirection(t *testing.T) {
	grid := Grid{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}
	expectedCoords := Coords{1, 2}
	coords, err := CoordsFromDirection(grid, ACTION_MOVE_BOTTOM)
	if !reflect.DeepEqual(coords, expectedCoords) {
		t.Error(fmt.Sprintf("The coords should be equal to { y: %d, x: %d}", expectedCoords.Y, expectedCoords.X))
	} else if err != nil {
		t.Error(fmt.Sprintf("The function should not return an error"))
	}
}

func TestCoordsFromDirectionNotValid(t *testing.T) {
	grid := Grid{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}
	_, err := CoordsFromDirection(grid, ACTION_MOVE_TOP)
	if err == nil {
		t.Error(fmt.Printf("The function should return an error because the direction is not valid"))
	}
}

func TestMove(t *testing.T) {
	grid := Grid{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}
	expectedGrid := Grid{
		{1, 2, 3},
		{4, 5, 0},
		{7, 8, 6},
	}
	coords := Coords{1, 2}
	newGrid, err := Move(grid, coords)

	if err != nil {
		t.Error("The move should not return an error")
	}
	if !reflect.DeepEqual(newGrid, expectedGrid) {
		t.Error("The grid should have changed")
	}
}

func TestMove2(t *testing.T) {
	expectedGrid := Grid{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}
	grid := Grid{
		{1, 2, 3},
		{4, 5, 0},
		{7, 8, 6},
	}
	coords := Coords{2, 2}
	newGrid, err := Move(grid, coords)

	if err != nil {
		t.Error("The move should not return an error")
	}
	if !reflect.DeepEqual(newGrid, expectedGrid) {
		t.Error("The grid should have changed")
	}
}
