package events

import (
	"fmt"
	"game"
	"renderer"
	"time"
)

const SUGGEST_DURATION time.Duration = 1

func SuggestListener(msgChan chan Message, grid game.Grid, startedGrid game.Grid) {
	suggestMsg := make(chan string, 1)
	defer close(suggestMsg)

	go func() {
		path, _ := game.SolvePuzzle(grid, startedGrid)
		if len(path) > 0 {
			suggestion := path[0]
			suggestMsg <- fmt.Sprintf("> Suggest: move tile n°%d by pressing the %s arrow", grid[suggestion.Y][suggestion.X], renderer.DrawMove(grid, suggestion))
		}
	}()

	select {
	case msg := <-suggestMsg:
		msgChan <- Message{msg, false}
		return
	case <-time.After(time.Second * SUGGEST_DURATION):
		return
	}
}
