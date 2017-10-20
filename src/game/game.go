package game

import (
	"errors"
	"reflect"
)

type Grid [][]byte

type Coords struct {
	y int
	x int
}

const EMPTY_VALUE byte = 0

func AreGridsEquals(grid Grid, grid2 Grid) bool {
	return reflect.DeepEqual(grid, grid2)
}

func AreCoordsEquals(coords Coords, coords2 Coords) bool {
	if coords.x == coords2.x && coords.y == coords2.y {
		return true
	}
	return false
}

func IsGridResolved(grid Grid, startedGrid Grid) bool {
	return AreGridsEquals(grid, startedGrid)
}

func BuildGrid(size byte) Grid {
	value := EMPTY_VALUE
	grid := make(Grid, size)

	for y := byte(0); y < size; y++ {
		grid[y] = make([]byte, size)
		for x := byte(0); x < size; x++ {
			value++
			if value == size*size {
				grid[y][x] = 0
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
	size := len(grid)
	y := 0
	for y < size {
		x := 0
		for x < len(grid[y]) {
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

	size := len(grid)
	if coordsEmptyTile.y-1 >= 0 {
		coordsMovableTiles = append(coordsMovableTiles, Coords{coordsEmptyTile.y - 1, coordsEmptyTile.x})
	}
	if coordsEmptyTile.x+1 < size {
		coordsMovableTiles = append(coordsMovableTiles, Coords{coordsEmptyTile.y, coordsEmptyTile.x + 1})
	}
	if coordsEmptyTile.y+1 < size {
		coordsMovableTiles = append(coordsMovableTiles, Coords{coordsEmptyTile.y + 1, coordsEmptyTile.x})
	}
	if coordsEmptyTile.x-1 >= 0 {
		coordsMovableTiles = append(coordsMovableTiles, Coords{coordsEmptyTile.y, coordsEmptyTile.x - 1})
	}
	return coordsMovableTiles
}

func isTileInMovableTiles(grid Grid, coordsTileToMove Coords) bool {
	for _, value := range ListMovableTiles(grid) {
		if AreCoordsEquals(value, coordsTileToMove) {
			return true
		}
	}
	return false
}

func ValueFromCoords(grid Grid, coords Coords) byte {
	return grid[coords.y][coords.x]
}

func Move(grid Grid, coordsTileToMove Coords) (Grid, error) {
	if !isTileInMovableTiles(grid, coordsTileToMove) {
		return grid, errors.New("The tile is not movable")
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
	newGrid[emptyCoords.y][emptyCoords.x] = grid[newCoords.y][newCoords.x]
	newGrid[newCoords.y][newCoords.x] = grid[emptyCoords.y][emptyCoords.x]
	return newGrid, nil
}

func MoveByValue(grid Grid, value byte) (Grid, error) {
	tileToMoveCoords, err := findTileByValue(grid, value)
	if err != nil {
		return grid, err
	}
	return Move(grid, tileToMoveCoords)
}
