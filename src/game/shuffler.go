package game

import (
	"time"
)

const SHUFFLE_DURATION time.Duration = 50

func Shuffle(grid Grid) (Grid, int) {
	timer := time.NewTimer(time.Millisecond * SHUFFLE_DURATION)

	gridShuffled := DeepCopyGrid(grid)

	count := 0
	var tileToMove Coords
	go func() {
		movableTiles, err := ListMovableTiles(gridShuffled)

		for {
			if err != nil {
				panic(err)
			}

			tileToMove = ChoiceCoords(movableTiles, time.Now().Unix())
			gridShuffled, _ = Move(gridShuffled, tileToMove)
			movableTiles, err = ListMovableTilesWithoutGoingBack(gridShuffled, gridShuffled[tileToMove.Y][tileToMove.X])
			count++
		}
	}()
	<-timer.C
	return gridShuffled, count
}
