package game

import (
	"fmt"
	"testing"
)

func TestAreGridsEquals(t *testing.T) {
	grid := Grid{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}
	grid2 := Grid{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}

	if !AreGridsEquals(grid, grid2) {
		t.Error("The grids should be equals by value and size")
	}
}

func TestAreGridsNotEqualsBySize(t *testing.T) {
	grid := Grid{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}
	grid2 := Grid{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 0},
	}

	if AreGridsEquals(grid, grid2) {
		t.Error("The grids should ne be equals due to their size")
	}
}

func TestAreGridsNotEqualsByValue(t *testing.T) {
	grid := Grid{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}
	grid2 := Grid{
		{3, 2, 1},
		{4, 5, 6},
		{7, 8, 0},
	}
	if AreGridsEquals(grid, grid2) {
		t.Error("The grids should ne be equals due to their values")
	}
}

func TestAreCoordsEquals(t *testing.T) {
	coords := Coords{
		y: 3,
		x: 2,
	}
	coords2 := Coords{
		y: 3,
		x: 2,
	}
	if !AreCoordsEquals(coords, coords2) {
		t.Error("The coords should be equals")
	}
}

func TestAreCoordsNotEqualsY(t *testing.T) {
	coords := Coords{
		y: 4,
		x: 2,
	}
	coords2 := Coords{
		y: 3,
		x: 2,
	}
	if AreCoordsEquals(coords, coords2) {
		t.Error("The coords should not be equals due to different y")
	}
}

func TestAreCoordsNotEqualsX(t *testing.T) {
	coords := Coords{
		y: 2,
		x: 2,
	}
	coords2 := Coords{
		y: 2,
		x: 1,
	}
	if AreCoordsEquals(coords, coords2) {
		t.Error("The coords should not be equals due to different x")
	}
}

func TestBuild(t *testing.T) {
	grid := BuildGrid(3)

	expectedGrid := Grid{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}

	if !AreGridsEquals(grid, expectedGrid) {
		t.Error("The grid is not built as expected")
	}
}

func TestListMovableTiles(t *testing.T) {
	grid := Grid{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}
	list := ListMovableTiles(grid)
	expectedValue := grid[list[0].y][list[0].x]
	expectedValue2 := grid[list[1].y][list[1].x]
	if expectedValue != 6 || expectedValue2 != 8 {
		t.Error(fmt.Sprintf("The movable tiles should be 6 and 8 and not %d and %d", expectedValue, expectedValue2))
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
	if !AreCoordsEquals(coords, expectedCoords) {
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
	if !AreGridsEquals(newGrid, expectedGrid) {
		t.Error("The grid should have changed")
	}
}

func TestMoveByValue(t *testing.T) {
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
	newGrid, err := MoveByValue(grid, byte(6))

	if err != nil {
		t.Error("The move should not return an error")
	}
	if !AreGridsEquals(newGrid, expectedGrid) {
		t.Error("The grid should have changed")
	}
}
