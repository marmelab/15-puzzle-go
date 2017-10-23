package events

import (
	"game"
	"renderer"
	"fmt"
)

func RenderListener(gridChan chan game.Grid) {
	for {
		grid := <-gridChan
		renderer.ClearTerminal()
		fmt.Println(renderer.DrawGrid(grid))
		fmt.Println("Move the tiles with the arrow keys or or press (S) to shuffle or press Esc to exit")
	}
}
