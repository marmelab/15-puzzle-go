package events

import (
	"fmt"
	"game"
	"renderer"
	"time"
)

const SUGGEST_DURATION time.Duration = 1

func SuggestListener(msgChan chan Message, grid game.Grid, startedGrid game.Grid) {
	timeout := make(chan bool, 1)

	go func() {
		time.Sleep(time.Second * SUGGEST_DURATION)
		timeout <- true
	}()

	go func() {
		path, err := game.SolvePuzzle(grid, startedGrid, timeout)
		var msg string
		if err != nil {
			msg = err.Error()
		} else if len(path) > 0 {
			suggestion := path[0]
			msg = fmt.Sprintf("Move tile n°%d by pressing the %s arrow", grid[suggestion.Y][suggestion.X], renderer.DrawMove(grid, suggestion))
		}
		msgChan <- Message{fmt.Sprintf("\n> Suggest: %s", msg), false}
	}()
}
