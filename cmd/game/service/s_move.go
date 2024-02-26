package service

import (
	"wargaming/kitex_gen/game"
	"wargaming/pkg/h3code"
)

func (s *GameService) Move(req *game.MoveReq) (*game.MoveResp, error) {
	resolution := 6

	// 获取当前图幅六角网格编码
	cornerCells := h3code.GetCellsFromPolygon(req.Corner, resolution)

	// 获取当前障碍物六角网格编码
	obstacleCells := h3code.GetCellsFromPolygon(req.Obstacle, resolution)

	// 获取原始位置六角网格编码
	originPosCell := h3code.GetCellFromPoint(req.OrginPos, resolution)

	// 获取目标移动位置六角网格编码
	targetPosCell := h3code.GetCellFromPoint(req.TargetPos, resolution)

	tiles := h3code.InitTiles(cornerCells, obstacleCells)

	var startTail, endTail *h3code.Tile

	for _, tile := range tiles {
		if tile.Position == originPosCell {
			startTail = tile
		}

		if tile.Position == targetPosCell {
			endTail = tile
		}
	}

	aStar := new(h3code.AStarPathfinding)

	path := aStar.FindPath(startTail, endTail)

	resp := new(game.MoveResp)

	for _, p := range path {
		resp.Path = append(resp.Path, &game.Point{
			Lat: p.Position.LatLng().Lat,
			Lng: p.Position.LatLng().Lng,
		})

	}

	return resp, nil
}
