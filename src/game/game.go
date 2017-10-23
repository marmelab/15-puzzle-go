package game

import (
	"errors"
	"fmt"
	"reflect"
	"unicode"
)

type Grid [][]byte

type Coords struct {
	y byte
	x byte
}

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
				tile.y = y
				tile.x = x
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
	if coordsEmptyTile.y > 0 {
		coordsMovableTiles = append(coordsMovableTiles, Coords{coordsEmptyTile.y - 1, coordsEmptyTile.x})
	}
	if coordsEmptyTile.x+1 < size {
		coordsMovableTiles = append(coordsMovableTiles, Coords{coordsEmptyTile.y, coordsEmptyTile.x + 1})
	}
	if coordsEmptyTile.y+1 < size {
		coordsMovableTiles = append(coordsMovableTiles, Coords{coordsEmptyTile.y + 1, coordsEmptyTile.x})
	}
	if coordsEmptyTile.x > 0 {
		coordsMovableTiles = append(coordsMovableTiles, Coords{coordsEmptyTile.y, coordsEmptyTile.x - 1})
	}
	return coordsMovableTiles, nil
}

func CoordsFromDirection(grid Grid, dir rune) (Coords, error) {
	var coordsMovableTiles Coords

	coordsEmptyTile, err := findEmptyTile(grid)
	if err != nil {
		return coordsEmptyTile, err
	}

	size := byte(len(grid))
	if unicode.ToLower(dir) == 's' {
		if coordsEmptyTile.y-1 != 255 {
			coordsMovableTiles.y = coordsEmptyTile.y - 1
			coordsMovableTiles.x = coordsEmptyTile.x
		} else {
			err = errors.New("It's not possible to move 'bottom'")
		}
	} else if unicode.ToLower(dir) == 'q' {
		if coordsEmptyTile.x+1 < size {
			coordsMovableTiles.y = coordsEmptyTile.y
			coordsMovableTiles.x = coordsEmptyTile.x + 1
		} else {
			err = errors.New("It's not possible to move 'left'")
		}
	} else if unicode.ToLower(dir) == 'z' {
		if coordsEmptyTile.y+1 < size {
			coordsMovableTiles.y = coordsEmptyTile.y + 1
			coordsMovableTiles.x = coordsEmptyTile.x
		} else {
			err = errors.New("It's not possible to move 'top'")
		}
	} else if unicode.ToLower(dir) == 'd' {
		if coordsEmptyTile.x-1 != 255 {
			coordsMovableTiles.y = coordsEmptyTile.y
			coordsMovableTiles.x = coordsEmptyTile.x - 1
		} else {
			err = errors.New("It's not possible to move 'right'")
		}
	}
	if err != nil {
		return coordsMovableTiles, err
	}
	return coordsMovableTiles, nil
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
		return grid, errors.New(fmt.Sprintf("The tile at coords (%d, %d) is not movable", coordsTileToMove.y, coordsTileToMove.y))
	}

	emptyCoords, err := findEmptyTile(grid)
	if err != nil {
		return grid, err
	}

	newCoords, err := findTileByValue(grid, grid[coordsTileToMove.y][coordsTileToMove.x])
	if err != nil {
		return grid, err
	}

	newGrid := DeepCopyGrid(grid)
	newGrid[emptyCoords.y][emptyCoords.x], newGrid[newCoords.y][newCoords.x] = grid[newCoords.y][newCoords.x], grid[emptyCoords.y][emptyCoords.x]
	return newGrid, nil
}
