package events

import (
	"game"
	"reflect"
	"renderer"
)

func GameListener(doneChan chan bool, inputChan chan byte, msgChan chan Message, startedGrid game.Grid) {
	grid := game.DeepCopyGrid(startedGrid)

	for {
		msgChan <- Message{renderer.DrawGrid(grid) + "\nMove the tiles with the arrow keys or or press (S) to shuffle or press Esc to exit", true}

		action := <-inputChan

		if action == game.ACTION_QUIT {
			doneChan <- false
			return
		}

		if action == game.ACTION_SHUFFLE {
			msgChan <- Message{"Shuffling...", true}
			grid = game.Shuffle(grid)
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
			if reflect.DeepEqual(grid, startedGrid) {
				msgChan <- Message{renderer.DrawGrid(startedGrid) + "\nGGWP, you solved the puzzle!", true}
				doneChan <- true
				return
			}
		}
	}
}
