package game

func ConvertGridToGridInt(grid Grid) [][]int {
	size := len(grid)
	gridInt := make([][]int, size)

	for y := 0; y < size; y++ {
		gridInt[y] = make([]int, size)
		for x := 0; x < size; x++ {
			gridInt[y][x] = int(grid[y][x])
		}
	}
	return gridInt
}

func ConvertGridIntToGrid(grid [][]int) Grid {
	size := len(grid)
	gridInt := make(Grid, size)

	for y := 0; y < size; y++ {
		gridInt[y] = make([]byte, size)
		for x := 0; x < size; x++ {
			gridInt[y][x] = byte(grid[y][x])
		}
	}
	return gridInt
}

func ConvertMoveStringToMove(MoveString string) byte {
	switch MoveString {
	case "top":
		return ACTION_MOVE_TOP
	case "right":
		return ACTION_MOVE_RIGHT
	case "bottom":
		return ACTION_MOVE_BOTTOM
	case "left":
		return ACTION_MOVE_LEFT
	}
	return ACTION_NONE
}

func ConvertMoveToMoveString(Move byte) string {
	switch Move {
	case ACTION_MOVE_TOP:
		return "top"
	case ACTION_MOVE_RIGHT:
		return "right"
	case ACTION_MOVE_BOTTOM:
		return "bottom"
	case ACTION_MOVE_LEFT:
		return "left"
	}
	return "none"
}
