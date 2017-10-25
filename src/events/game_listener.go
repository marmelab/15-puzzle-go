package events

import (
	"fmt"
	"game"
	"reflect"
	"renderer"
)

func GameListener(doneChan chan bool, inputChan chan byte, msgChan chan Message, startedGrid game.Grid) {
	count := 0

	grid := game.DeepCopyGrid(startedGrid)

	for {
		msgChan <- Message{fmt.Sprintf("Turn %d\n%s\nMove the tiles with the arrow keys or or press (S) to shuffle or press Esc to exit", count, renderer.DrawGrid(grid)), true}

		action := <-inputChan

		if action == game.ACTION_QUIT {
			doneChan <- false
			break
		}

		if action == game.ACTION_SHUFFLE {
			msgChan <- Message{"Shuffling...", true}
			grid = game.Shuffle(grid)
			count = 0
		} else {
			newCoords, err := game.CoordsFromDirection(grid, action)
			if err != nil {
				msgChan <- Message{err.Error(), false}
				continue
			}
			grid, err = game.Move(grid, newCoords)
			if err != nil {
				msgChan <- Message{"The move is not possible, please try another direction", false}
				continue
			}
			count++
			if reflect.DeepEqual(grid, startedGrid) {
				msgChan <- Message{fmt.Sprintf("%s\nGGWP, you solved the puzzle in %d turn(s)!", renderer.DrawGrid(startedGrid), count), true}
				doneChan <- true
				break
			}
		}
	}
}
