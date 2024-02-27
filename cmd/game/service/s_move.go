package service

import (
	"github.com/uber/h3-go/v4"
	"wargaming/kitex_gen/game"
	"wargaming/pkg/h3code"
)

func (s *GameService) Move(req *game.MoveReq) (*game.MoveResp, error) {

	var cells []h3.Cell
	var obstacleCells []h3.Cell

	for _, it := range req.Cells {
		cells = append(cells, h3.Cell(it))
	}

	for _, it := range req.Obstacle {
		obstacleCells = append(obstacleCells, h3.Cell(it))
	}

	tiles := h3code.InitTiles(cells, obstacleCells)

	var startTail, endTail *h3code.Tile

	for _, tile := range tiles {
		if tile.Position == h3.Cell(req.OriginCell) {
			startTail = tile
		}

		if tile.Position == h3.Cell(req.TargetCell) {
			endTail = tile
		}
	}

	aStar := new(h3code.AStarPathfinding)

	path := aStar.FindPath(startTail, endTail)

	resp := new(game.MoveResp)

	for _, p := range path {
		resp.Path = append(resp.Path, int64(p.Position))

	}

	return resp, nil
}
