package game

import (
	"errors"
	"reflect"
	"sort"
)

func AddToPriorityList(list []Node, node Node) []Node {
	list = append(list, node)
	sort.Slice(list[:], func(i, j int) bool {
		return list[i].Heuristic < list[j].Heuristic
	})
	return list
}

func RemoveFromPriorityList(list []Node) ([]Node, Node) {
	node := list[0]
	return append(list[:0], list[1:]...), node
}

func SolvePuzzle(shuffledGrid Grid, solvedGrid Grid, timeout chan bool) ([]Coords, error) {
	var coords []Coords

	node := Node{0, TaxicabWithValues(shuffledGrid, solvedGrid), shuffledGrid, coords}
	var openList []Node
	openList = AddToPriorityList(openList, node)

	timedOut := false
	for len(openList) > 0 && !timedOut {
		select {
		case timedOut = <-timeout:
		default:
			openList, node = RemoveFromPriorityList(openList)

			if reflect.DeepEqual(node.Grid, solvedGrid) {
				return node.Moves, nil
			}

			possibleMoves, _ := ListMovableTiles(node.Grid)
			for _, coords := range possibleMoves {
				neighborGrid, _ := Move(node.Grid, coords)
				neighborNode := Node{node.Cost + 1, node.Cost + 1 + Taxicab(neighborGrid, solvedGrid), neighborGrid, append(node.Moves, coords)}

				openList = AddToPriorityList(openList, neighborNode)
			}
		}
	}
	if timedOut {
		return make([]Coords, 0), errors.New("The solver has been stopped by a timeout")
	}
	return make([]Coords, 0), errors.New("No solution found by the deep puzzle algorithm")
}
