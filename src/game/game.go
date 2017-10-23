package game

import (
	"errors"
	"fmt"
	"reflect"
)

type Grid [][]byte

type Coords struct {
	y byte
	x byte
}

const EMPTY_VALUE byte = 0

func AreGridsEquals(grid Grid, grid2 Grid) bool {
	return reflect.DeepEqual(grid, grid2)
}

func AreCoordsEquals(coords Coords, coords2 Coords) bool {
	return reflect.DeepEqual(coords, coords2)
}

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
	return findTileByValue(grid, 0)
}

func ListMovableTiles(grid Grid) []Coords {
	var coordsMovableTiles []Coords

	coordsEmptyTile, err := findEmptyTile(grid)
	if err != nil {
		return coordsMovableTiles
	}

	size := byte(len(grid))
	if coordsEmptyTile.y-1 != 255 {
		coordsMovableTiles = append(coordsMovableTiles, Coords{coordsEmptyTile.y - 1, coordsEmptyTile.x})
	}
	if coordsEmptyTile.x+1 < size {
		coordsMovableTiles = append(coordsMovableTiles, Coords{coordsEmptyTile.y, coordsEmptyTile.x + 1})
	}
	if coordsEmptyTile.y+1 < size {
		coordsMovableTiles = append(coordsMovableTiles, Coords{coordsEmptyTile.y + 1, coordsEmptyTile.x})
	}
	if coordsEmptyTile.x-1 != 255 {
		coordsMovableTiles = append(coordsMovableTiles, Coords{coordsEmptyTile.y, coordsEmptyTile.x - 1})
	}
	return coordsMovableTiles
}

func CoordsFromDirection(grid Grid, dir rune) (Coords, error) {
	var coordsMovableTiles Coords

	coordsEmptyTile, err := findEmptyTile(grid)
	if err != nil {
		return coordsEmptyTile, err
	}

	size := byte(len(grid))
	if dir == 's' || dir == 'S' {
		if coordsEmptyTile.y-1 != 255 {
			coordsMovableTiles.y = coordsEmptyTile.y - 1
			coordsMovableTiles.x = coordsEmptyTile.x
		} else {
			err = errors.New("It's not possible to move 'bottom'")
		}
	} else if dir == 'q' || dir == 'Q' {
		if coordsEmptyTile.x+1 < size {
			coordsMovableTiles.y = coordsEmptyTile.y
			coordsMovableTiles.x = coordsEmptyTile.x + 1
		} else {
			err = errors.New("It's not possible to move 'left'")		
		}
	} else if dir == 'z' || dir == 'Z' {
		if coordsEmptyTile.y+1 < size {
			coordsMovableTiles.y = coordsEmptyTile.y + 1
			coordsMovableTiles.x = coordsEmptyTile.x
		} else {
			err = errors.New("It's not possible to move 'top'")
		}
	} else if dir == 'd' || dir == 'D' {
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

func isTileInMovableTiles(grid Grid, coordsTileToMove Coords) bool {
	for _, value := range ListMovableTiles(grid) {
		if AreCoordsEquals(value, coordsTileToMove) {
			return true
		}
	}
	return false
}

func Move(grid Grid, coordsTileToMove Coords) (Grid, error) {
	if !isTileInMovableTiles(grid, coordsTileToMove) {
		return grid, errors.New(fmt.Sprintf("The tile ate coords (%d, %d) is not movable", coordsTileToMove.y, coordsTileToMove.y))
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

func MoveByValue(grid Grid, value byte) (Grid, error) {
	tileToMoveCoords, err := findTileByValue(grid, value)
	if err != nil {
		return grid, err
	}
	return Move(grid, tileToMoveCoords)
}
