package game

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConvertGridToGridInt(t *testing.T) {
	grid := Grid{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}

	expectedGridInt := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}

	if !reflect.DeepEqual(ConvertGridToGridInt(grid), expectedGridInt) {
		t.Error("The grid converted is not the same as the expected one")
	}
}

func TestConvertGridIntToGrid(t *testing.T) {
	gridInt := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}

	expectedGrid := Grid{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}

	if !reflect.DeepEqual(ConvertGridIntToGrid(gridInt), expectedGrid) {
		t.Error("The grid converted is not the same as the expected one")
	}
}

func TestConvertMoveStringToMoveTop(t *testing.T) {
	expectedMove := ACTION_MOVE_TOP

	if ConvertMoveStringToMove("top") != expectedMove {
		t.Error(fmt.Sprintf("The move converted should be equals to %d", ACTION_MOVE_TOP))
	}
}

func TestConvertMoveStringToMoveRight(t *testing.T) {
	expectedMove := ACTION_MOVE_RIGHT

	if ConvertMoveStringToMove("right") != expectedMove {
		t.Error(fmt.Sprintf("The move converted should be equals to %d", ACTION_MOVE_RIGHT))
	}
}

func TestConvertMoveStringToMoveBottom(t *testing.T) {
	expectedMove := ACTION_MOVE_BOTTOM

	if ConvertMoveStringToMove("bottom") != expectedMove {
		t.Error(fmt.Sprintf("The move converted should be equals to %d", ACTION_MOVE_BOTTOM))
	}
}

func TestConvertMoveStringToMoveLeft(t *testing.T) {
	expectedMove := ACTION_MOVE_LEFT

	if ConvertMoveStringToMove("left") != expectedMove {
		t.Error(fmt.Sprintf("The move converted should be equals to %d", ACTION_MOVE_LEFT))
	}
}

func TestConvertMoveStringToMoveNone(t *testing.T) {
	expectedMove := ACTION_NONE

	if ConvertMoveStringToMove("qsdsdq") != expectedMove {
		t.Error(fmt.Sprintf("The move converted should be equals to %d", ACTION_NONE))
	}
}

func TestConvertMoveToMoveStringTop(t *testing.T) {
	expectedMove := "top"

	if ConvertMoveToMoveString(ACTION_MOVE_TOP) != expectedMove {
		t.Error(fmt.Sprintf("The move converted should be equals to %s", "top"))
	}
}

func TestConvertMoveToMoveStringRight(t *testing.T) {
	expectedMove := "right"

	if ConvertMoveToMoveString(ACTION_MOVE_RIGHT) != expectedMove {
		t.Error(fmt.Sprintf("The move converted should be equals to %s", "right"))
	}
}

func TestConvertMoveToMoveStringBottom(t *testing.T) {
	expectedMove := "bottom"

	if ConvertMoveToMoveString(ACTION_MOVE_BOTTOM) != expectedMove {
		t.Error(fmt.Sprintf("The move converted should be equals to %s", "bottom"))
	}
}

func TestConvertMoveToMoveStringLeft(t *testing.T) {
	expectedMove := "left"

	if ConvertMoveToMoveString(ACTION_MOVE_LEFT) != expectedMove {
		t.Error(fmt.Sprintf("The move converted should be equals to %s", "left"))
	}
}

func TestConvertMoveToMoveStringNone(t *testing.T) {
	expectedMove := "none"

	if ConvertMoveToMoveString(ACTION_NONE) != expectedMove {
		t.Error(fmt.Sprintf("The move converted should be equals to %s", "none"))
	}
}
