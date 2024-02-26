package h3code

import (
	"fmt"
	"testing"
	"wargaming/kitex_gen/game"
)

var points = []*game.Point{
	{Lat: 35.12, Lng: 114.47},
	{Lat: 38.24, Lng: 114.47},
	{Lat: 38.24, Lng: 122.42},
	{Lat: 35.12, Lng: 122.42},
}

var startPoint = game.Point{Lat: 37.52, Lng: 117.21}

var endPoint = game.Point{Lat: 36.46, Lng: 118.07}

var obstaclePoints []*game.Point

func TestAStarPathfinding_FindPath(t *testing.T) {
	cells := GetCellsFromPolygon(points, 6)

	obstacles := GetCellsFromPolygon(obstaclePoints, 6)

	startCell := GetCellFromPoint(&startPoint, 6)

	endCell := GetCellFromPoint(&endPoint, 6)

	tiles := InitTiles(cells, obstacles)

	var startTail, endTail *Tile

	for _, tile := range tiles {
		if tile.Position == startCell {
			startTail = tile
		}

		if tile.Position == endCell {
			endTail = tile
		}
	}

	aStar := new(AStarPathfinding)

	path := aStar.FindPath(startTail, endTail)

	for _, t := range path {
		fmt.Println(t.Position.LatLng())
	}

	fmt.Println(path)
}
