package h3code

import (
	"wargaming/kitex_gen/game"

	"github.com/uber/h3-go/v4"
)

func GetCellsFromPolygon(points []*game.Point, resolution int) []h3.Cell {
	var latLngArr []h3.LatLng
	for _, p := range points {
		latLngArr = append(latLngArr, h3.NewLatLng(p.Lat, p.Lng))
	}
	cells := h3.PolygonToCells(
		h3.GeoPolygon{
			GeoLoop: latLngArr,
		},
		resolution,
	)

	return cells
}

func GetCellFromPoint(point *game.Point, resolution int) h3.Cell {
	cell := h3.LatLngToCell(
		h3.LatLng{
			Lat: point.Lat,
			Lng: point.Lng,
		},
		resolution,
	)

	return cell
}
