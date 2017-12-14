package game

import (
	"errors"
	"reflect"
	"sort"
	"time"
)

type Node struct {
	Cost      int
	Heuristic int
	Grid      Grid
	Moves     []Coords
}

const SUGGEST_DURATION time.Duration = 1

func CompareTwoNodesByCost(n1 Node, n2 Node) int {
	if n1.Cost < n2.Cost {
		return 1
	}
	if n1.Cost == n2.Cost {
		return 0
	}
	return -1
}

func IsNodeInListWithLowerCost(list []Node, node Node) bool {
	for _, value := range list {
		if reflect.DeepEqual(node.Grid, value.Grid) && CompareTwoNodesByCost(value, node) < 1 {
			return true
		}
	}
	return false
}

func IsNodeInListOfListWithLowerCost(list [][]Node, node Node) bool {
	for _, subList := range list {
		if IsNodeInListWithLowerCost(subList, node) {
			return true
		}
	}
	return false
}

func SortListWithHeuristic(list []Node) []Node {
	var listCopied []Node
	listCopied = append(listCopied, list...)
	sort.Slice(listCopied[:], func(i, j int) bool {
		return listCopied[i].Heuristic < listCopied[j].Heuristic
	})
	return listCopied
}

func BuildNeighborList(node Node, solvedGrid Grid, possibleMoves []Coords) []Node {
	var subList []Node

	for _, coords := range possibleMoves {
		neighborGrid, _ := Move(node.Grid, coords)
		heuristic := TaxicabWithLinearConflict(neighborGrid, solvedGrid)
		neighborNode := Node{node.Cost + 1 + heuristic, heuristic, neighborGrid, append(node.Moves, coords)}
		subList = append(subList, neighborNode)
	}
	return subList
}

func BuildDepthList(list []Node, solvedGrid Grid, listChan chan []Node, timeoutChan chan bool) {
	var sublist []Node

	for _, node := range list {
		select {
		case <-timeoutChan:
			return
		default:
			previousMoveCoords := node.Moves[len(node.Moves)-1]
			previousValue := node.Grid[previousMoveCoords.Y][previousMoveCoords.X]
			possibleMoves, _ := ListMovableTilesWithoutGoingBack(node.Grid, previousValue)
			for _, neighbor := range BuildNeighborList(node, solvedGrid, possibleMoves) {
				sublist = append(sublist, neighbor)
			}
		}
	}
	listChan <- sublist
}

func findResolvedGridInList(list []Node) (Node, error) {
	var node Node
	for _, node = range list {
		if node.Heuristic == 0 {
			return node, nil
		}
	}
	return node, errors.New("No solved grid found in list")
}

func buildInitialList(shuffledGrid Grid, solvedGrid Grid, possibleMoves []Coords) [][]Node {
	initialNode := Node{0, TaxicabWithLinearConflict(shuffledGrid, solvedGrid), shuffledGrid, make([]Coords, 0)}
	list := make([][]Node, 0)
	sublist := BuildNeighborList(initialNode, solvedGrid, possibleMoves)
	list = append(list, sublist)
	return list
}

func SolvePuzzleD(shuffledGrid Grid, solvedGrid Grid) ([]Coords, error) {
	possibleMoves, _ := ListMovableTiles(shuffledGrid)
	return solvePuzzleD(shuffledGrid, solvedGrid, possibleMoves)
}

func SolvePuzzleDWithPreviousMove(shuffledGrid Grid, solvedGrid Grid, lastMove Coords) ([]Coords, error) {
	previousValue := shuffledGrid[lastMove.Y][lastMove.X]
	possibleMoves, _ := ListMovableTilesWithoutGoingBack(shuffledGrid, previousValue)

	return solvePuzzleD(shuffledGrid, solvedGrid, possibleMoves)
}

func solvePuzzleD(shuffledGrid Grid, solvedGrid Grid, possibleMoves []Coords) ([]Coords, error) {
	list := buildInitialList(shuffledGrid, solvedGrid, possibleMoves)
	node, err := findResolvedGridInList(list[0])
	if err == nil {
		return node.Moves, nil
	}

	timeoutChan := make(chan bool, 1)
	defer close(timeoutChan)
	listChan := make(chan []Node, 1)
	defer close(listChan)

	solved := false

	go func() {
		time.Sleep(time.Second * SUGGEST_DURATION)
		if solved {
			return
		}
		timeoutChan <- true
	}()

	depth := 0
	for {
		go BuildDepthList(list[depth], solvedGrid, listChan, timeoutChan)

		select {
		case <-timeoutChan:
			sortedList := SortListWithHeuristic(list[depth-1])
			return sortedList[0].Moves, nil
		case sublist := <-listChan:
			node, err := findResolvedGridInList(sublist)
			if err == nil {
				solved = true
				return node.Moves, nil
			}

			list = append(list, sublist)
		}
		depth++
	}
}
