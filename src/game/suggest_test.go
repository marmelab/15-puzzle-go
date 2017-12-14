package game

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCompareTwoNodesByCostPositive(t *testing.T) {
	var grid Grid
	var coords []Coords

	n1 := Node{1, 0, grid, coords}
	n2 := Node{10, 0, grid, coords}

	comparison := CompareTwoNodesByCost(n1, n2)

	if comparison != 1 {
		t.Error("The comparison of the two nodes with costs 1 and 10 should return 1")
	}
}

func TestCompareTwoNodesByCostNegative(t *testing.T) {
	var grid Grid
	var coords []Coords

	n1 := Node{10, 0, grid, coords}
	n2 := Node{1, 0, grid, coords}

	comparison := CompareTwoNodesByCost(n1, n2)

	if comparison != -1 {
		t.Error("The comparison of the two nodes with costs 10 and 1 should return -1")
	}
}

func TestCompareTwoNodesByCostEquals(t *testing.T) {
	var grid Grid
	var coords []Coords

	n1 := Node{10, 0, grid, coords}
	n2 := Node{10, 0, grid, coords}

	comparison := CompareTwoNodesByCost(n1, n2)

	if comparison != 0 {
		t.Error("The comparison of the two nodes equals by cost should return 0")
	}
}

func TestIsNodeInListWithLowerCost(t *testing.T) {
	grid1 := Grid{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}
	grid2 := Grid{
		{1, 2, 3},
		{4, 5, 0},
		{7, 8, 6},
	}
	grid3 := Grid{
		{1, 2, 3},
		{4, 0, 5},
		{7, 8, 6},
	}
	grid4 := Grid{
		{1, 0, 3},
		{4, 2, 5},
		{7, 8, 6},
	}
	coords := make([]Coords, 0)

	n1 := Node{0, 0, grid1, coords}
	n2 := Node{2, 0, grid2, coords}
	n3 := Node{3, 0, grid3, coords}
	n4 := Node{4, 0, grid4, coords}

	listNodes := make([]Node, 0)
	listNodes = append(listNodes, n1)
	listNodes = append(listNodes, n2)
	listNodes = append(listNodes, n3)
	listNodes = append(listNodes, n4)

	grid := Grid{
		{1, 0, 3},
		{4, 2, 5},
		{7, 8, 6},
	}

	n := Node{2, 0, grid, coords}

	isInList := IsNodeInListWithLowerCost(listNodes, n)

	if !isInList {
		t.Error("The node should be in the list with a lower cost")
	}
}

func TestNotIsNodeInList(t *testing.T) {
	grid1 := Grid{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}
	grid2 := Grid{
		{1, 2, 3},
		{4, 5, 0},
		{7, 8, 6},
	}
	grid3 := Grid{
		{1, 2, 3},
		{4, 0, 5},
		{7, 8, 6},
	}
	grid4 := Grid{
		{1, 0, 3},
		{4, 2, 5},
		{7, 8, 6},
	}
	coords := make([]Coords, 0)

	n1 := Node{0, 0, grid1, coords}
	n2 := Node{2, 0, grid2, coords}
	n3 := Node{3, 0, grid3, coords}
	n4 := Node{4, 0, grid4, coords}

	listNodes := make([]Node, 0)
	listNodes = append(listNodes, n1)
	listNodes = append(listNodes, n2)
	listNodes = append(listNodes, n3)
	listNodes = append(listNodes, n4)

	grid := Grid{
		{0, 1, 3},
		{4, 2, 5},
		{7, 8, 6},
	}

	n := Node{5, 0, grid, coords}

	isInList := IsNodeInListWithLowerCost(listNodes, n)

	if isInList {
		t.Error("The node should not be in the list")
	}
}

func TestNotIsNodeInListWithLowerCost(t *testing.T) {
	grid1 := Grid{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}
	grid2 := Grid{
		{1, 2, 3},
		{4, 5, 0},
		{7, 8, 6},
	}
	grid3 := Grid{
		{1, 2, 3},
		{4, 0, 5},
		{7, 8, 6},
	}
	grid4 := Grid{
		{1, 0, 3},
		{4, 2, 5},
		{7, 8, 6},
	}
	coords := make([]Coords, 0)

	n1 := Node{0, 0, grid1, coords}
	n2 := Node{2, 0, grid2, coords}
	n3 := Node{3, 0, grid3, coords}
	n4 := Node{4, 0, grid4, coords}

	listNodes := make([]Node, 0)
	listNodes = append(listNodes, n1)
	listNodes = append(listNodes, n2)
	listNodes = append(listNodes, n3)
	listNodes = append(listNodes, n4)

	grid := Grid{
		{1, 0, 3},
		{4, 2, 5},
		{7, 8, 6},
	}

	n := Node{11, 0, grid, coords}

	isInList := IsNodeInListWithLowerCost(listNodes, n)

	if isInList {
		t.Error("The node should not be in the list with a lower cost")
	}
}

func TestSortListWithHeuristic(t *testing.T) {
	grid := Grid{
		{0, 1, 3},
		{4, 2, 5},
		{7, 8, 6},
	}
	coords := make([]Coords, 0)

	n1 := Node{0, 0, grid, coords}
	n2 := Node{1, 1, grid, coords}
	n3 := Node{2, 2, grid, coords}
	n4 := Node{3, 3, grid, coords}

	listNodes := make([]Node, 0)
	listNodes = append(listNodes, n2)
	listNodes = append(listNodes, n4)
	listNodes = append(listNodes, n1)
	listNodes = append(listNodes, n3)

	expectedListNodes := make([]Node, 0)
	expectedListNodes = append(expectedListNodes, n1)
	expectedListNodes = append(expectedListNodes, n2)
	expectedListNodes = append(expectedListNodes, n3)
	expectedListNodes = append(expectedListNodes, n4)

	if !reflect.DeepEqual(SortListWithHeuristic(listNodes), expectedListNodes) {
		t.Error("The list should be sorted")
	}
}

func TestBuildNeighborList(t *testing.T) {
	grid := Grid{
		{1, 2, 3},
		{4, 5, 0},
		{7, 8, 6},
	}
	solvedGrid := Grid{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}

	coords := make([]Coords, 0)
	possibleMoves := make([]Coords, 0)
	possibleMoves = append(possibleMoves, Coords{0, 2})
	possibleMoves = append(possibleMoves, Coords{1, 1})
	possibleMoves = append(possibleMoves, Coords{2, 2})

	n := Node{0, 0, grid, coords}

	g1 := Grid{
		{1, 2, 0},
		{4, 5, 3},
		{7, 8, 6},
	}
	g2 := Grid{
		{1, 2, 3},
		{4, 0, 5},
		{7, 8, 6},
	}
	g3 := Grid{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}

	h1 := TaxicabWithValues(g1, solvedGrid)
	h2 := TaxicabWithValues(g2, solvedGrid)
	h3 := TaxicabWithValues(g3, solvedGrid)

	c1 := make([]Coords, 0)
	c1 = append(c1, Coords{0, 2})
	c2 := make([]Coords, 0)
	c2 = append(c2, Coords{1, 1})
	c3 := make([]Coords, 0)
	c3 = append(c3, Coords{2, 2})

	n1 := Node{1 + h1, h1, g1, c1}
	n2 := Node{1 + h2, h2, g2, c2}
	n3 := Node{1 + h3, h3, g3, c3}

	expectedResult := make([]Node, 0)
	expectedResult = append(expectedResult, n1)
	expectedResult = append(expectedResult, n2)
	expectedResult = append(expectedResult, n3)

	result := BuildNeighborList(n, solvedGrid, possibleMoves)
	if !reflect.DeepEqual(result, expectedResult) {
		t.Error("The neighbor list should be sorted")
	}
}
