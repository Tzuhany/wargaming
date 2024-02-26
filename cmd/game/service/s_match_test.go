package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"testing"
	"wargaming/kitex_gen/game"
)

func TestGameService_Match(t *testing.T) {
	svc := NewGameService(context.Background())

	resp, err := svc.Match(&game.MatchReq{
		UserId: 1,
	})
	if err != nil {
		klog.Error(err)
		return
	}

	klog.Info(resp.MatchedUserId)
}
