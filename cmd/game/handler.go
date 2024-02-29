package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"wargaming/cmd/game/pack"
	"wargaming/cmd/game/service"
	game "wargaming/kitex_gen/game"
)

// GameServiceImpl implements the last service interface defined in the IDL.
type GameServiceImpl struct{}

// Match implements the GameServiceImpl interface.
func (s *GameServiceImpl) Match(ctx context.Context, req *game.MatchReq) (resp *game.MatchResp, err error) {
	resp = new(game.MatchResp)

	moveResp, err := service.NewGameService(ctx).Match(req)
	if err != nil {
		klog.Error(err)
		resp.Base = pack.BuildBaseResp(err)
		return
	}

	resp.Base = pack.BuildBaseResp(nil)
	resp.MatchedUserId = moveResp.MatchedUserId

	return
}

// Move implements the GameServiceImpl interface.
func (s *GameServiceImpl) Move(ctx context.Context, req *game.FindPathReq) (resp *game.FindPathResp, err error) {
	resp = new(game.FindPathResp)

	moveResp, err := service.NewGameService(ctx).Move(req)
	if err != nil {
		klog.Error(err)
		resp.Base = pack.BuildBaseResp(err)
		return
	}

	resp.Base = pack.BuildBaseResp(nil)
	resp.Path = moveResp.Path

	return
}
