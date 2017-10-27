package game

import (
	"fmt"
	"reflect"
	"testing"
)

func TestChoiceCoordsOneElement(t *testing.T) {
	coordsArray := make([]Coords, 1)
	coordsArray[0] = Coords{2, 3}
	expectedCoords := Coords{2, 3}

	element := ChoiceCoords(coordsArray, 1)
	if !reflect.DeepEqual(element, expectedCoords) {
		t.Error("The choice method should return {y: 2, x: 3}")
	}
}

func TestChoiceCoordsMultipleElements(t *testing.T) {
	coordsArray := make([]Coords, 4)
	coordsArray[0] = Coords{1, 1}
	coordsArray[1] = Coords{1, 2}
	coordsArray[2] = Coords{2, 1}
	coordsArray[3] = Coords{2, 2}

	expectedCoords := Coords{2, 1}

	element := ChoiceCoords(coordsArray, 1)
	if !reflect.DeepEqual(element, expectedCoords) {
		t.Error(fmt.Sprintf("The choice method should return {y: 1, x: 2} and not {y: %d; x: %d}", element.X, element.Y))
	}
}
