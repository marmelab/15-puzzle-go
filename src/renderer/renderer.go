package renderer

import (
	"game"
	"strconv"
)

func Hello() string {
	return "Welcome to the 15 puzzle game.\n"
}

func BuildLine(startSymb string, stopSymb string, sepSymb string, tileSize int) string {
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

func BuildGrid(grid game.Grid) string {
	size := len(grid)

	horizontalLine := BuildLine("├", "┤", "┼", size)
	firstHorizontalLine := BuildLine("┌", "┐", "┬", size)
	lastHorizontalLine := BuildLine("└", "┘", "┴", size)

	gridToShow := firstHorizontalLine

	var tile_str string

	for i := 0; i < size; i++ {
		gridToShow += "│"
		for j := 0; j < size; j++ {
			tile := grid[i][j]
			if tile > 0 && tile < 10 {
				tile_str = "  " + strconv.Itoa(tile) + " │"
			} else if tile == 0 {
				tile_str = "    │"
			} else {
				tile_str = " " + strconv.Itoa(tile) + " │"
			}
			gridToShow += tile_str
		}
		if i != size-1 {
			gridToShow += horizontalLine
		}
	}
	gridToShow += lastHorizontalLine
	return gridToShow
}
