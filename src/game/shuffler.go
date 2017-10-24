package game

import (
	"time"
)

func Shuffle(grid Grid) Grid {
	timer := time.NewTimer(time.Second * 1)

	gridShuffled := DeepCopyGrid(grid)
	go func() {
		for {
			movableTiles, err := ListMovableTiles(gridShuffled)
			if err != nil {
				panic(err)
			}
			tileToMove, _ := ChoiceCoords(movableTiles)
			gridShuffled, _ = Move(gridShuffled, tileToMove)
		}
	}()
	<-timer.C
	return gridShuffled
}
