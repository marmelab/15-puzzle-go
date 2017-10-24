package events

import (
	"fmt"
	"game"
	"reflect"
	"renderer"
)

func GameListener(doneChan chan bool, inputChan chan byte, gridChan chan game.Grid, startedGrid game.Grid) {
	grid := game.DeepCopyGrid(startedGrid)
	gridChan <- grid
	for {
		action := <-inputChan
		if action == game.ACTION_QUIT {
			doneChan <- false
			return
		}

		if action == game.ACTION_SHUFFLE {
			renderer.ClearTerminal()
			fmt.Println("Shuffling...")
			grid = game.Shuffle(grid)
		} else {
			newCoords, err := game.CoordsFromDirection(grid, action)
			if err != nil {
				fmt.Println(err)
				continue
			}
			grid, err = game.Move(grid, newCoords)
			if err != nil {
				fmt.Println("The move is not possible, please try another direction")
				continue
			}
			if reflect.DeepEqual(grid, startedGrid) {
				doneChan <- true
			}
		}
		gridChan <- game.DeepCopyGrid(grid)
	}
}
