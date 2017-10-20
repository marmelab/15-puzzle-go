package game

type Grid [][]int

func AreGridEquals(grid Grid, grid2 Grid) bool {
	if len(grid) != len(grid2) {
		return false
	}
	size := len(grid)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if grid[i][j] != grid2[i][j] {
				return false
			}
		}
	}
	return true
}

func BuildGrid(size int) Grid {
	value := 0
	grid := make([][]int, size)

	for i := 0; i < size; i++ {
		grid[i] = make([]int, size)
		for j := 0; j < size; j++ {
			value++
			if value == size*size {
				grid[i][j] = 0
			} else {
				grid[i][j] = value
			}
		}
	}
	return grid
}
