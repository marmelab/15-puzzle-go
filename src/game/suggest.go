package game

import (
	"errors"
	"math"
	"reflect"
	"sort"
)

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

type Node struct {
	Cost      int
	Heuristic int
	Grid      Grid
	Moves     []Coords
}

func CompareTwoNodesByCost(n1 Node, n2 Node) int {
	if n1.Cost < n2.Cost {
		return 1
	} else if n1.Cost == n2.Cost {
		return 0
	} else {
		return -1
	}
}

func IsNodeInListWithLowerCost(list []Node, node Node) bool {
	for _, value := range list {
		if reflect.DeepEqual(node.Grid, value.Grid) && CompareTwoNodesByCost(value, node) < 1 {
			return true
		}
	}
	return false
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
	list = list[1:]
	return list, node
}

func BuildPath(node Node) []Coords {
	return node.Moves
}

func DeepPuzzleAlgorithm(shuffledGrid Grid, solvedGrid Grid) ([]Coords, error) {
	var coords []Coords

	node := Node{0, TaxicabWithValues(shuffledGrid, solvedGrid), shuffledGrid, coords}
	var openList []Node
	openList = AddToPriorityList(openList, node)

	for len(openList) > 0 {
		openList, node = RemoveFromPriorityList(openList)

		if reflect.DeepEqual(node.Grid, solvedGrid) {
			return node.Moves, nil
		}

		possibleMoves, _ := ListMovableTiles(node.Grid)
		for _, coords := range possibleMoves {
			neighboorGrid, _ := Move(node.Grid, coords)
			neighboorNode := Node{node.Cost + 1, node.Cost + 1 + Taxicab(neighboorGrid, solvedGrid), neighboorGrid, append(node.Moves, coords)}

			openList = AddToPriorityList(openList, neighboorNode)
		}
	}

	return coords, errors.New("No solution found by the deep puzzle algorithm")
}
