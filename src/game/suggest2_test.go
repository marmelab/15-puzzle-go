package game

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddToPriorityListEnd(t *testing.T) {
	var coords []Coords

	grid := Grid{
		{1, 0, 3},
		{4, 2, 5},
		{7, 8, 6},
	}
	n1 := Node{0, 0, grid, coords}
	n2 := Node{2, 1, grid, coords}
	n3 := Node{3, 2, grid, coords}
	n4 := Node{4, 3, grid, coords}

	listNodes := make([]Node, 0)
	listNodes = append(listNodes, n1)
	listNodes = append(listNodes, n2)
	listNodes = append(listNodes, n3)
	listNodes = append(listNodes, n4)

	grid2 := Grid{
		{0, 1, 3},
		{4, 2, 5},
		{7, 8, 6},
	}
	n := Node{11, 0, grid2, coords}

	listNodes = AddToPriorityList(listNodes, n)

	if reflect.DeepEqual(listNodes[4].Grid, n.Grid) {
		t.Error("The node should be at the end of the priority list")
	}
}

func TestAddToPriorityListTop(t *testing.T) {
	var coords []Coords

	grid := Grid{
		{1, 0, 3},
		{4, 2, 5},
		{7, 8, 6},
	}

	n1 := Node{0, 0, grid, coords}
	n2 := Node{2, 1, grid, coords}
	n3 := Node{3, 2, grid, coords}
	n4 := Node{4, 3, grid, coords}

	var listNodes []Node
	listNodes = append(listNodes, n1)
	listNodes = append(listNodes, n2)
	listNodes = append(listNodes, n3)
	listNodes = append(listNodes, n4)

	grid2 := Grid{
		{0, 1, 3},
		{4, 2, 5},
		{7, 8, 6},
	}
	n := Node{0, 0, grid2, coords}

	listNodes = AddToPriorityList(listNodes, n)

	if reflect.DeepEqual(listNodes[0].Grid, n.Grid) {
		t.Error("The node should be at the top of the priority list")
	}
}

func TestRemoveFromPriorityList(t *testing.T) {
	coords := make([]Coords, 0)
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

	n1 := Node{0, 0, grid1, coords}
	n2 := Node{2, 1, grid2, coords}
	n3 := Node{3, 2, grid3, coords}
	n4 := Node{4, 3, grid4, coords}

	listNodes := make([]Node, 0)
	listNodes = append(listNodes, n1)
	listNodes = append(listNodes, n2)
	listNodes = append(listNodes, n3)
	listNodes = append(listNodes, n4)

	var n Node
	listNodes, n = RemoveFromPriorityList(listNodes)

	listNodesSize := len(listNodes)
	if listNodesSize != 3 {
		t.Error(fmt.Sprintf("The length's list should now be equals to 3 and not %d", listNodesSize))
	}

	if n.Cost != 0 || n.Heuristic != 0 || !reflect.DeepEqual(n.Grid, grid1) {
		t.Error(fmt.Sprintf("The returned node should be the first one and not {Cost: %d, Heuristic: %d}", n.Cost, n.Heuristic))
	}
}
