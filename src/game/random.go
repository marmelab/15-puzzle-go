package game

import (
	"math/rand"
)

func ChoiceCoords(array []Coords, seed int64) Coords {
	rand.Seed(seed)
	n := rand.Int() % len(array)
	return array[n]
}

func ChoiceCoordsNoSeed(array []Coords) Coords {
	n := rand.Int() % len(array)
	return array[n]
}
