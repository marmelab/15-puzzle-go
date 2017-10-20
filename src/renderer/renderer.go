package renderer

import (
	"game"
	"fmt"
)

func Hello() string {
	return "Welcome to the 15 puzzle game.\n"
}

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
