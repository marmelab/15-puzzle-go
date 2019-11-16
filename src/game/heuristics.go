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

func Distance(y int, goalY int, coeffY int, x int, goalX int, coeffX int) int {
	return int(math.Abs(float64(y-goalY))) * coeffY + int(math.Abs(float64(x-int(goalX)))) * coeffX
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
				sum += size*size + Distance(y, int(expectedPos.Y), 3, x, int(expectedPos.X), 1) * int(currentValue)
			} else {
				sum += int(currentValue)
			}
			if x+1 < size {
				followingXValue := grid[y][x+1]
				if followingXValue != 0 && currentValue > followingXValue {
					sum += 2
				}
			}
			if y+1 < size {
				followingYValue := grid[y+1][x]
				if  followingYValue != 0 && currentValue > followingYValue {
					sum += 2
				}
			}

			x++
		}
		y++
	}
	return sum
}
