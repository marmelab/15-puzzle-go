package renderer

import (
	"fmt"
	"game"
)

func DrawLine(startSymb string, stopSymb string, sepSymb string, tileSize int) string {
	line := "\n" + startSymb

	tileLine := "────"

	for i := 0; i < tileSize; i++ {
		line += tileLine
		if i != tileSize-1 {
			line += sepSymb
		}
	}

	line += stopSymb + "\n"

	return line
}

func DrawGrid(grid game.Grid) string {
	size := len(grid)

	horizontalLine := DrawLine("├", "┤", "┼", size)
	firstHorizontalLine := DrawLine("┌", "┐", "┬", size)
	lastHorizontalLine := DrawLine("└", "┘", "┴", size)

	gridToShow := firstHorizontalLine

	var tileStr string

	for i := 0; i < size; i++ {
		gridToShow += "│"
		for j := 0; j < size; j++ {
			tile := grid[i][j]
			if tile > 0 && tile < 10 {
				tileStr = fmt.Sprintf("  %d │", tile)
			} else if tile == 0 {
				tileStr = "    │"
			} else {
				tileStr = fmt.Sprintf(" %d │", tile)
			}
			gridToShow += tileStr
		}
		if i != size-1 {
			gridToShow += horizontalLine
		}
	}
	gridToShow += lastHorizontalLine
	return gridToShow
}

func DrawMove(grid game.Grid, coords game.Coords) string {
	dir, err := game.DirectionFromCoords(grid, coords)
	if err != nil {
		return ""
	}
	switch dir {
	case game.ACTION_MOVE_TOP:
		return "top"
	case game.ACTION_MOVE_RIGHT:
		return "right"
	case game.ACTION_MOVE_BOTTOM:
		return "bottom"
	case game.ACTION_MOVE_LEFT:
		return "left"
	}
	return ""
}
