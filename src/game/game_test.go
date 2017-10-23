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
	expectedValue := grid[list[0].y][list[0].x]
	expectedValue2 := grid[list[1].y][list[1].x]
	if expectedValue != 6 || expectedValue2 != 8 {
		t.Error(fmt.Sprintf("The movable tiles should be 6 and 8 and not %d and %d", expectedValue, expectedValue2))
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

func TestCoordsFromDirection(t *testing.T) {
	grid := Grid{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}
	expectedCoords := Coords{
		y: 1,
		x: 2,
	}
	coords, err := CoordsFromDirection(grid, 'S')
	if !reflect.DeepEqual(coords, expectedCoords) {
		t.Error(fmt.Sprintf("The coords should be equal to { y: %d, x: %d}", expectedCoords.y, expectedCoords.x))
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
	_, err := CoordsFromDirection(grid, 'Z')
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
	coords := Coords{
		y: 1,
		x: 2,
	}
	newGrid, err := Move(grid, coords)

	if err != nil {
		t.Error("The move should not return an error")
	}
	if !reflect.DeepEqual(newGrid, expectedGrid) {
		t.Error("The grid should have changed")
	}
}
