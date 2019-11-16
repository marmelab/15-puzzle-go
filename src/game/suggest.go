package game

import (
	"errors"
	"sort"
	"time"
)

type Node struct {
	Heuristic int
	Grid      Grid
	Moves     []Coords
	PreviousValue byte
}

const SUGGEST_DURATION time.Duration = 200

func SortListWithHeuristic(list []Node) []Node {
	var listCopied []Node
	listCopied = append(listCopied, list...)
	sort.Slice(listCopied[:], func(i, j int) bool {
		return listCopied[i].Heuristic < listCopied[j].Heuristic
	})
	return listCopied
}

func BuildNeighborList(node Node, solvedGrid Grid, possibleMoves []Coords) []Node {
	var sublist []Node

	for _, coords := range possibleMoves {
		previousValue := node.Grid[coords.Y][coords.X]
		neighborGrid, _ := Move(node.Grid, coords)
		heuristic := TaxicabWithLinearConflict(neighborGrid, solvedGrid)
		var moves []Coords
		moves = append(moves, node.Moves...)
		moves = append(moves, coords)
		neighborNode := Node{heuristic, neighborGrid, moves, previousValue}
		sublist = append(sublist, neighborNode)
	}
	return sublist
}

func BuildDepthList(list []Node, solvedGrid Grid, listChan chan []Node, timeoutChan chan bool) {
	var sublist []Node

	for _, node := range list {
		select {
		case <-timeoutChan:
			return
		default:
			possibleMoves, _ := ListMovableTilesWithoutGoingBack(node.Grid, node.PreviousValue)
			neighboor := BuildNeighborList(node, solvedGrid, possibleMoves)
			sublist = append(sublist, neighboor...)
		}
	}
	listChan <- sublist
}

func findResolvedGridInList(list []Node) (Node, error) {
	var node Node
	sizeSide := len(node.Grid)
	size := (sizeSide) * (sizeSide)
	minCost := ((size * (size + 1)) / 2) - size
	for _, node = range list {
		if node.Heuristic == minCost {
			return node, nil
		}
	}
	return node, errors.New("No solved grid found in list")
}

func isResolved(node Node) bool {
	sizeSide := len(node.Grid)
	size := (sizeSide) * (sizeSide)
	minCost := ((size * (size + 1)) / 2) - size
	return node.Heuristic == minCost
}

func buildInitialList(shuffledGrid Grid, solvedGrid Grid, possibleMoves []Coords) []Node {
	initialNode := Node{TaxicabWithLinearConflict(shuffledGrid, solvedGrid), shuffledGrid, make([]Coords, 0), 0}
	return BuildNeighborList(initialNode, solvedGrid, possibleMoves)
}

func SolvePuzzleD(shuffledGrid Grid, solvedGrid Grid) ([]Coords, error) {
	possibleMoves, _ := ListMovableTiles(shuffledGrid)
	return solvePuzzleD(shuffledGrid, solvedGrid, possibleMoves)
}

func SolvePuzzleDWithPreviousMove(shuffledGrid Grid, solvedGrid Grid, previousValue byte) ([]Coords, error) {
	possibleMoves, _ := ListMovableTilesWithoutGoingBack(shuffledGrid, previousValue)

	return solvePuzzleD(shuffledGrid, solvedGrid, possibleMoves)
}

func solvePuzzleD(shuffledGrid Grid, solvedGrid Grid, possibleMoves []Coords) ([]Coords, error) {
	list := buildInitialList(shuffledGrid, solvedGrid, possibleMoves)
	sortedList := SortListWithHeuristic(list)
	if isResolved(sortedList[0]) {
		return sortedList[0].Moves, nil
	}

	timeoutChan := make(chan bool, 1)
	defer close(timeoutChan)
	listChan := make(chan []Node, 1)
	defer close(listChan)

	solved := false

	go func() {
		time.Sleep(time.Millisecond * SUGGEST_DURATION)
		if solved {
			return
		}
		timeoutChan <- true
	}()

	for {
		go BuildDepthList(list, solvedGrid, listChan, timeoutChan)
		select {
		case <-timeoutChan:
			sortedList := SortListWithHeuristic(list)
			LogNode("solution.log", sortedList[0])
			LogList("list.log", sortedList)
			return sortedList[0].Moves, nil
		case list = <-listChan:
			sortedList := SortListWithHeuristic(list)
			if isResolved(sortedList[0]) {
				solved = true
				return sortedList[0].Moves, nil
			}
		}
	}
}
