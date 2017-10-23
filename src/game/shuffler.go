package game

import (
	"time"
)

func Shuffle(grid Grid) Grid{
	timer := time.NewTimer(time.Second * 1)
	
	gridShuffled := DeepCopyGrid(grid)
	go func() {
		for {
			movableTiles := ListMovableTiles(gridShuffled)
			tileToMove, _ := ChoiceCoords(movableTiles)
			gridShuffled, _ = Move(gridShuffled, tileToMove)
		}
	}()
	<- timer.C
	return gridShuffled
}
