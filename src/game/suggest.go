package game

import (
	"errors"
	"math"
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

func CountMisplacedTiles(grid Grid, grid2 Grid) int {
	sum := 0
	size := len(grid)
	y := 0
	for y < size {
		x := 0
		for x < size {
			if grid[y][x] != grid2[y][x] {
				sum++
			}
			x++
		}
		y++
	}
	return sum
}

func Taxicab(grid Grid, grid2 Grid) int {
	sum := 0
	size := len(grid)
	y := 0

	for y < size {
		x := 0
		for x < size {
			if grid[y][x] != grid2[y][x] {
				expectedPos, _ := findTileByValue(grid2, grid[y][x])
				sum += int(math.Abs(float64(y-int(expectedPos.Y))) + math.Abs(float64(x-int(expectedPos.X))))
			}
			x++
		}
		y++
	}
	return sum
}

func TaxicabWithValues(grid Grid, grid2 Grid) int {
	sum := 0
	size := len(grid)
	y := 0

	for y < size {
		x := 0
		for x < size {
			if grid[y][x] != grid2[y][x] {
				expectedPos, _ := findTileByValue(grid2, grid[y][x])
				sum += int(math.Abs(float64(y-int(expectedPos.Y))) + math.Abs(float64(x-int(expectedPos.X))))
				sum += size*size - int(grid[y][x])
			}
			x++
		}
		y++
	}
	return sum
}

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

func BuildNeighborList(node Node, solvedGrid Grid, possibleMoves []Coords) []Node {
	var subList []Node

	for _, coords := range possibleMoves {
		neighborGrid, _ := Move(node.Grid, coords)
		heuristic := TaxicabWithValues(neighborGrid, solvedGrid)
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

func findGridResolvedInList(list []Node) (Node, error) {
	var node Node
	for _, node = range list {
		if node.Heuristic == 0 {
			return node, nil
		}
	}
	return node, errors.New("No solved grid found in list")
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
	initialNode := Node{0, TaxicabWithValues(shuffledGrid, solvedGrid), shuffledGrid, make([]Coords, 0)}

	timeoutChan := make(chan bool, 1)
	defer close(timeoutChan)
	listChan := make(chan []Node, 1)
	defer close(listChan)

	list := make([][]Node, 0)

	sublist := BuildNeighborList(initialNode, solvedGrid, possibleMoves)
	node, err := findGridResolvedInList(sublist)
	if err == nil {
		return node.Moves, nil
	}

	list = append(list, sublist)
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
		case sublist = <-listChan:
			node, err := findGridResolvedInList(sublist)
			if err == nil {
				solved = true
				return node.Moves, nil
			}

			list = append(list, sublist)
		}
		depth++
	}
}
