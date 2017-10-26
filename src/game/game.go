package game

import (
	"errors"
	"fmt"
	"reflect"
)

type Grid [][]byte

type Coords struct {
	Y byte
	X byte
}

const ACTION_NONE byte = 0
const ACTION_QUIT byte = 1
const ACTION_MOVE_TOP byte = 2
const ACTION_MOVE_RIGHT byte = 3
const ACTION_MOVE_BOTTOM byte = 4
const ACTION_MOVE_LEFT byte = 5
const ACTION_SHUFFLE byte = 6
const ACTION_HELP byte = 7

const EMPTY_VALUE byte = 0

func BuildGrid(size byte) Grid {
	value := EMPTY_VALUE
	grid := make(Grid, size)

	for y := byte(0); y < size; y++ {
		grid[y] = make([]byte, size)
		for x := byte(0); x < size; x++ {
			value++
			if value == size*size {
				grid[y][x] = EMPTY_VALUE
			} else {
				grid[y][x] = value
			}
		}
	}
	return grid
}

func DeepCopyGrid(grid Grid) Grid {
	newGrid := make(Grid, len(grid))
	for y := range grid {
		newGrid[y] = make([]byte, len(grid[y]))
		copy(newGrid[y], grid[y])
	}
	return newGrid
}

func findTileByValue(grid Grid, value byte) (Coords, error) {
	var tile Coords
	size := byte(len(grid))
	y := byte(0)
	for y < size {
		x := byte(0)
		for x < byte(len(grid[y])) {
			if grid[y][x] == value {
				tile.Y = y
				tile.X = x
				return tile, nil
			}
			x++
		}
		y++
	}
	return tile, errors.New("The grid does not contain this tile")
}

func findEmptyTile(grid Grid) (Coords, error) {
	return findTileByValue(grid, EMPTY_VALUE)
}

func ListMovableTiles(grid Grid) ([]Coords, error) {
	var coordsMovableTiles []Coords

	coordsEmptyTile, err := findEmptyTile(grid)
	if err != nil {
		return coordsMovableTiles, err
	}

	size := byte(len(grid))
	if coordsEmptyTile.Y > 0 {
		coordsMovableTiles = append(coordsMovableTiles, Coords{coordsEmptyTile.Y - 1, coordsEmptyTile.X})
	}
	if coordsEmptyTile.X+1 < size {
		coordsMovableTiles = append(coordsMovableTiles, Coords{coordsEmptyTile.Y, coordsEmptyTile.X + 1})
	}
	if coordsEmptyTile.Y+1 < size {
		coordsMovableTiles = append(coordsMovableTiles, Coords{coordsEmptyTile.Y + 1, coordsEmptyTile.X})
	}
	if coordsEmptyTile.X > 0 {
		coordsMovableTiles = append(coordsMovableTiles, Coords{coordsEmptyTile.Y, coordsEmptyTile.X - 1})
	}
	return coordsMovableTiles, nil
}

func ListMovableTilesWithoutGoingBack(grid Grid, previousMovedTile byte) ([]Coords, error) {
	movableTiles, err := ListMovableTiles(grid)
	if err != nil {
		return movableTiles, err
	}

	indexToRemove := -1
	for index, coords := range movableTiles {
		if grid[coords.Y][coords.X] == previousMovedTile {
			indexToRemove = index
			break
		}
	}
	if indexToRemove >= 0 {
		movableTiles = movableTiles[:indexToRemove]
	}
	return movableTiles, nil
}

func CoordsFromDirection(grid Grid, dir byte) (Coords, error) {
	var coordsMovableTiles Coords

	coordsEmptyTile, err := findEmptyTile(grid)
	if err != nil {
		return coordsEmptyTile, err
	}

	size := byte(len(grid))
	switch dir {
	case ACTION_MOVE_TOP:
		if coordsEmptyTile.Y+1 < size {
			coordsMovableTiles.Y = coordsEmptyTile.Y + 1
			coordsMovableTiles.X = coordsEmptyTile.X
		} else {
			err = errors.New("It's not possible to move 'top'")
		}
		break
	case ACTION_MOVE_RIGHT:
		if coordsEmptyTile.X-1 != 255 {
			coordsMovableTiles.Y = coordsEmptyTile.Y
			coordsMovableTiles.X = coordsEmptyTile.X - 1
		} else {
			err = errors.New("It's not possible to move 'right'")
		}
		break
	case ACTION_MOVE_BOTTOM:
		if coordsEmptyTile.Y-1 != 255 {
			coordsMovableTiles.Y = coordsEmptyTile.Y - 1
			coordsMovableTiles.X = coordsEmptyTile.X
		} else {
			err = errors.New("It's not possible to move 'bottom'")
		}
		break
	case ACTION_MOVE_LEFT:
		if coordsEmptyTile.X+1 < size {
			coordsMovableTiles.Y = coordsEmptyTile.Y
			coordsMovableTiles.X = coordsEmptyTile.X + 1
		} else {
			err = errors.New("It's not possible to move 'left'")
		}
		break
	}

	if err != nil {
		return coordsMovableTiles, err
	}
	return coordsMovableTiles, nil
}

func DirectionFromCoords(grid Grid, coords Coords) (byte, error) {
	coordsEmptyTile, err := findEmptyTile(grid)
	if err != nil {
		return ACTION_NONE, err
	}

	Y := float64(coords.Y) - float64(coordsEmptyTile.Y)
	X := float64(coords.X) - float64(coordsEmptyTile.X)

	if Y > 0 {
		return ACTION_MOVE_TOP, nil
	} else if Y < 0 {
		return ACTION_MOVE_BOTTOM, nil
	} else {
		if X > 0 {
			return ACTION_MOVE_LEFT, nil
		} else if X < 0 {
			return ACTION_MOVE_RIGHT, nil
		} else {
			return ACTION_NONE, errors.New("The tile cannot move")
		}
	}
}

func isTileInMovableTiles(grid Grid, coordsTileToMove Coords) (bool, error) {
	movableTiles, err := ListMovableTiles(grid)
	if err != nil {
		return false, err
	}
	for _, value := range movableTiles {
		if reflect.DeepEqual(value, coordsTileToMove) {
			return true, nil
		}
	}
	return false, nil
}

func Move(grid Grid, coordsTileToMove Coords) (Grid, error) {
	isTileMovable, err := isTileInMovableTiles(grid, coordsTileToMove)
	if err != nil {
		return grid, err
	}

	if !isTileMovable {
		return grid, errors.New(fmt.Sprintf("The tile at coords (%d, %d) is not movable", coordsTileToMove.Y, coordsTileToMove.X))
	}

	emptyCoords, err := findEmptyTile(grid)
	if err != nil {
		return grid, err
	}

	newCoords, err := findTileByValue(grid, grid[coordsTileToMove.Y][coordsTileToMove.X])
	if err != nil {
		return grid, err
	}

	newGrid := DeepCopyGrid(grid)
	newGrid[emptyCoords.Y][emptyCoords.X], newGrid[newCoords.Y][newCoords.X] = grid[newCoords.Y][newCoords.X], grid[emptyCoords.Y][emptyCoords.X]
	return newGrid, nil
}
