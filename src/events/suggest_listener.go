package events

import (
	"fmt"
	"game"
	"renderer"
)

func SuggestListener(msgChan chan Message, grid game.Grid, startedGrid game.Grid) {
	path, err := game.SolvePuzzleD(grid, startedGrid)
	var msg string
	if err != nil {
		msg = err.Error()
	} else if len(path) > 0 {
		suggestion := path[0]
		msg += fmt.Sprintf("\nMove tile n°%d by pressing the %s arrow", grid[suggestion.Y][suggestion.X], renderer.DrawMove(grid, suggestion))
	}
	msgChan <- Message{fmt.Sprintf("\n> Suggest: %s", msg), false}
}
