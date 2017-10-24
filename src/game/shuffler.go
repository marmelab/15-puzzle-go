package game

import (
	"time"
)

const SHUFFLE_DURATION time.Duration = 1

func Shuffle(grid Grid) Grid {
	timer := time.NewTimer(time.Second * SHUFFLE_DURATION)

	gridShuffled := DeepCopyGrid(grid)
	go func() {
		for {
			movableTiles, err := ListMovableTiles(gridShuffled)
			if err != nil {
				panic(err)
			}
			tileToMove := ChoiceCoords(movableTiles, time.Now().Unix())
			gridShuffled, _ = Move(gridShuffled, tileToMove)
		}
	}()
	<-timer.C
	return gridShuffled
}
