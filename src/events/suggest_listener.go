package events

import (
	"game"
	"renderer"
	"fmt"
)

func SuggestListener(msgChan chan Message, gridChan chan game.Grid, startedGrid game.Grid) {
	for {
		grid := <-gridChan
		path, _ := game.DeepPuzzleAlgorithm2(grid, startedGrid)
		if len(path) > 0 {
			suggestion := path[0]
			msgChan <- Message{fmt.Sprintf("> Suggest: move tile n°%d by pressing the %s arrow", grid[suggestion.Y][suggestion.X], renderer.DrawMove(grid, suggestion)), false}						
		}
	}
}
