package game

import (
	"fmt"
	"testing"
)

func TestCountMisplacedTilesIdenticalGrids(t *testing.T) {
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

	sum := CountMisplacedTiles(startedGrid, notShuffledGrid)

	if sum != 0 {
		t.Error("The misplaced sum of two identical grids should be equals to 0")
	}
}

func TestCountMisplacedTiles(t *testing.T) {
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

	sum := CountMisplacedTiles(startedGrid, shuffledGrid)

	if sum != 5 {
		t.Error("The misplaced tiles counter of two grids should be equals to 5")
	}
}

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
		t.Error("The taxicab sum of two identical grids should be equals to 0")
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
		t.Error("The taxicab sum of two different grids should be equals to 6")
	}
}

func TestTaxicabWithValuesIdenticalGrids(t *testing.T) {
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
		t.Error("The taxicab sum of two identical grids should be equals to 0")
	}
}

func TestTaxicabWithValues(t *testing.T) {
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

	sum := TaxicabWithValues(startedGrid, shuffledGrid)

	if sum != 35 {
		t.Error(fmt.Sprintf("The taxicab sum with values of two different grids should be equals to 35 and not %d", sum))
	}
}