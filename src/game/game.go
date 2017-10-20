package game

type Grid [][]byte

const EMPTY_VALUE byte = 0

func AreGridsEquals(grid Grid, grid2 Grid) bool {
	size := len(grid)
	if size != len(grid2) {
		return false
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if grid[i][j] != grid2[i][j] {
				return false
			}
		}
	}
	return true
}

func BuildGrid(size byte) Grid {
	value := EMPTY_VALUE
	grid := make([][]byte, size)

	for i := byte(0); i < size; i++ {
		grid[i] = make([]byte, size)
		for j := byte(0); j < size; j++ {
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
