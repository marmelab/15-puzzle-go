package events

import (
	"fmt"
	"game"
	"reflect"
	"renderer"
	"time"
)

const TIME_SLEEP time.Duration = 800

func GameListener(doneChan chan bool, inputChan chan byte, msgChan chan Message, startedGrid game.Grid) {
	grid := game.DeepCopyGrid(startedGrid)
	turnCounter := 0
	for {
		msgChan <- Message{fmt.Sprintf("Turn %d\n%s\nMove the tiles with the arrow keys\nor press (S) to shuffle\nor press (H) to be helped\nor press Esc to exit", turnCounter, renderer.DrawGrid(grid)), true}

		action := <-inputChan

		if action == game.ACTION_QUIT {
			doneChan <- false
			break
		}

		if action == game.ACTION_SHUFFLE {
			msgChan <- Message{"Shuffling...", true}
			grid, _ = game.Shuffle(grid)
			time.Sleep(time.Millisecond * TIME_SLEEP)
			turnCounter = 1
		} else if action == game.ACTION_HELP {
			if turnCounter == 0 {
				msgChan <- Message{"\n> Shuffle the puzzle before asking help :)", false}
				time.Sleep(time.Millisecond * TIME_SLEEP)
				continue
			}
			go SuggestListener(msgChan, grid, startedGrid)
		} else {
			newCoords, err := game.CoordsFromDirection(grid, action)
			if err != nil {
				msgChan <- Message{fmt.Sprintf("\n> %s", err.Error()), false}
				time.Sleep(time.Millisecond * TIME_SLEEP)
				continue
			}
			grid, err = game.Move(grid, newCoords)
			if err != nil {
				msgChan <- Message{"\n> The move is not possible, please try another direction", false}
				time.Sleep(time.Millisecond * TIME_SLEEP)
				continue
			}
			turnCounter++
			if reflect.DeepEqual(grid, startedGrid) {
				msgChan <- Message{fmt.Sprintf("Congratulations\n\n%s\nGGWP, you solved the puzzle in %d turn(s)!", renderer.DrawGrid(startedGrid), turnCounter), true}
				doneChan <- true
				break
			}
		}
	}
}
