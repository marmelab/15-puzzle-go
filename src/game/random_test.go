package game

import (
	"fmt"
	"testing"
)

func TestChoiceSimpleArray(t *testing.T) {
	var wrongArray [1]Coords
	wrongArray[0] = Coords{y: 2, x: 3}
	expectedCoords := Coords{y: 2, x: 3}

	element, err := Choice(wrongArray)
	if ArreCoordsEquals(element, expectedCoords) {
		t.Error("The choice method should return 34")		
	} else if err == nil {
		t.Error("The choice method should not return an error")
	}
}
