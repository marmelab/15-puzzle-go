package game

import (
	"math/rand"
	"time"
)

func ChoiceCoords(array []Coords) (Coords, error) {
	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(array)
	return array[n], nil
}