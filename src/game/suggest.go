package game

import(
	"math"
)

func Taxicab(grid Grid, grid2 Grid) byte {	
	sum := float64(0)
	size := len(grid)
	y := 0

	for y < size {
		x := 0
		for x < size {
			if grid[y][x] != grid2[y][x] {
				expectedPos, _ := findTileByValue(grid2, grid[y][x])
				sum += math.Abs(float64(y - int(expectedPos.y))) + math.Abs(float64(x - int(expectedPos.x)))
			}
			x++
		}
		y++
	}
	return byte(sum)
}
