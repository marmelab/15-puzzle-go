package renderer

import (
	"game"
	"testing"
)

func TestDrawGrid(t *testing.T) {
	expectedGridRendered := `
┌────┬────┬────┬────┐
│  1 │  2 │  3 │  4 │
├────┼────┼────┼────┤
│  5 │  6 │  7 │  8 │
├────┼────┼────┼────┤
│  9 │ 10 │ 11 │ 12 │
├────┼────┼────┼────┤
│ 13 │ 14 │ 15 │    │
└────┴────┴────┴────┘
`

	gridRendered := DrawGrid(game.BuildGrid(4))

	if expectedGridRendered != gridRendered {
		t.Error("The grid layout is not corresponding to expected one")
	}

	expectedGridRendered2 := `
┌────┬────┬────┐
│  1 │  2 │  3 │
├────┼────┼────┤
│  4 │  5 │  6 │
├────┼────┼────┤
│  7 │  8 │    │
└────┴────┴────┘
`

	gridRendered2 := DrawGrid(game.BuildGrid(3))

	if expectedGridRendered2 != gridRendered2 {
		t.Error("The grid layout is not corresponding to expected one")
	}
}
