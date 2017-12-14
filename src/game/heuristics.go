package game

import (
	"math"
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
				expectedPos, _ := FindTileByValue(grid2, grid[y][x])
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
				expectedPos, _ := FindTileByValue(grid2, grid[y][x])
				sum += int(math.Abs(float64(y-int(expectedPos.Y))) + math.Abs(float64(x-int(expectedPos.X))))
				sum += size*size - int(grid[y][x])
			}
			x++
		}
		y++
	}
	return sum
}

func TaxicabWithLinearConflict(grid Grid, grid2 Grid) int {
	sum := 0
	size := len(grid)
	y := 0

	for y < size {
		x := 0
		for x < size {
			currentValue := grid[y][x]
			expectedValue := grid2[y][x]
			if currentValue != expectedValue {
				expectedPos, _ := FindTileByValue(grid2, currentValue)
				sum += int(math.Abs(float64(y-int(expectedPos.Y))) + math.Abs(float64(x-int(expectedPos.X))))
				sum += size*size - int(currentValue)
			}
			if x+1 < size {
				followingValue := grid[y][x+1]
				if currentValue > followingValue {
					sum += 2
				}
			}
			x++
		}
		y++
	}
	return sum
}
