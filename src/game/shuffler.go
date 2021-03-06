package game

import (
	"time"
)

const SHUFFLE_DURATION time.Duration = 50

type ShuffleResult struct {
	Grid       Grid
	CountMoves int
}

func Shuffle(grid Grid) (Grid, int) {
	timer := time.NewTimer(time.Millisecond * SHUFFLE_DURATION)

	shuffleChan := make(chan ShuffleResult, 1)
	defer close(shuffleChan)

	go func() {
		var tileToMove Coords
		gridShuffled := DeepCopyGrid(grid)
		movableTiles, err := ListMovableTiles(gridShuffled)
		countMoves := 0

		for {
			select {
			case <-timer.C:
				shuffleChan <- ShuffleResult{gridShuffled, countMoves}
				return
			default:
				if err != nil {
					panic(err)
				}

				tileToMove = ChoiceCoordsNoSeed(movableTiles)
				gridShuffled, _ = Move(gridShuffled, tileToMove)
				movableTiles, err = ListMovableTilesWithoutGoingBack(gridShuffled, gridShuffled[tileToMove.Y][tileToMove.X])
				countMoves++
			}
		}
	}()

	result := <-shuffleChan
	return result.Grid, result.CountMoves
}
