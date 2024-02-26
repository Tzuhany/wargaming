package service

import "context"

type GameService struct {
	ctx context.Context
}

func NewGameService(ctx context.Context) *GameService {
	return &GameService{ctx: ctx}
}
