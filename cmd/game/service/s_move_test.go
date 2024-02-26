package service

import (
	"context"
	"fmt"
	"testing"
	"wargaming/kitex_gen/game"

	"github.com/cloudwego/kitex/pkg/klog"
)

var corner = []*game.Point{
	{Lat: 35.12, Lng: 114.47},
	{Lat: 38.24, Lng: 114.47},
	{Lat: 38.24, Lng: 122.42},
	{Lat: 35.12, Lng: 122.42},
}

var originPos = game.Point{Lat: 37.52, Lng: 117.21}

var targetPos = game.Point{Lat: 36.46, Lng: 118.07}

func TestGameService_Move(t *testing.T) {
	svc := NewGameService(context.Background())

	resp, err := svc.Move(&game.MoveReq{
		OrginPos:  &originPos,
		TargetPos: &targetPos,
		Obstacle:  make([]*game.Point, 0),
		Corner:    corner,
	})

	if err != nil {
		klog.Error(err)
	}

	for _, p := range resp.Path {
		fmt.Println(p)
	}
}
