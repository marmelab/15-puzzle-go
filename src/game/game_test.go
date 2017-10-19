package game

import (
	"testing"
)

func TestAreGridEquals(t *testing.T) {
	grid := Grid{
		Data: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 0},
		},
		Size: 3,
	}

	grid2 := Grid{
		Data: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 0},
		},
		Size: 3,
	}

	grid3 := Grid{
		Data: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 0},
		},
		Size: 2,
	}

	grid4 := Grid{
		Data: [][]int{
			{3, 2, 1},
			{4, 5, 6},
			{7, 8, 0},
		},
		Size: 3,
	}

	if !AreGridEquals(grid, grid2) {
		t.Error("The grids should be equals by value and size")
	}

	if AreGridEquals(grid, grid3) {
		t.Error("The grids should ne be equals due to their size")
	}

	if AreGridEquals(grid, grid4) {
		t.Error("The grids should ne be equals due to their values")
	}
}

func TestBuild(t *testing.T) {
	grid := BuildGrid(3)

	expectedGrid := Grid{
		Data: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 0},
		},
		Size: 3,
	}

	if !AreGridEquals(grid, expectedGrid) {
		t.Error("The grid is not built as expected")
	}
}
