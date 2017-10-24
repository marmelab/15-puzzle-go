package game

import (
	"reflect"
	"testing"
)

func TestChoiceSimpleArray(t *testing.T) {
	wrongArray := make([]Coords, 1)
	wrongArray[0] = Coords{y: 2, x: 3}
	expectedCoords := Coords{y: 2, x: 3}

	element := ChoiceCoords(wrongArray)
	if !reflect.DeepEqual(element, expectedCoords) {
		t.Error("The choice method should return 34")
	}
}
