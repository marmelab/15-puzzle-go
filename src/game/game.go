package game

type Grid struct {
	Data [][]int
	Size int
}

func AreGridEquals(grid Grid, grid2 Grid) bool {
	if grid.Size != grid2.Size {
		return false
	}
	size := grid.Size
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if grid.Data[i][j] != grid2.Data[i][j] {
				return false
			}
		}
	}
	return true
}

func BuildGrid(size int) Grid {
	value := 0
	data := make([][]int, size)

	for i := 0; i < size; i++ {
		data[i] = make([]int, size)
		for j := 0; j < size; j++ {
			value++
			if value == size*size {
				data[i][j] = 0
			} else {
				data[i][j] = value
			}
		}
	}

	return Grid{
		Data: data,
		Size: size,
	}
}
