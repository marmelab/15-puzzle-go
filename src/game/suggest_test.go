package game

import (
	"testing"
)

func TestTaxicabIdenticalGrids(t *testing.T) {
	startedGrid := Grid{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}
	notShuffledGrid := Grid{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}

	sum := Taxicab(startedGrid, notShuffledGrid)

	if sum != 0 {
		t.Error("The taxicab sum of two identical grid should be equals to 0")
	}
}

func TestTaxicab(t *testing.T) {
	startedGrid := Grid{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}
	shuffledGrid := Grid{
		{1, 5, 2},
		{4, 0, 3},
		{7, 8, 6},
	}

	sum := Taxicab(startedGrid, shuffledGrid)

	if sum != 6 {
		t.Error("The taxicab sum of two identical grid should be equals to 0")
	}
}
