package game

import (
	"testing"
)

func TestAreGridsEquals(t *testing.T) {
	grid := Grid {
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
	grid := Grid {
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
	grid := Grid {
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
