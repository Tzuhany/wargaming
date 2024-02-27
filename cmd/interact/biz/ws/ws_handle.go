package ws

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"time"
	"wargaming/cmd/interact/biz/rpc"
	"wargaming/kitex_gen/game"
	"wargaming/pkg/constants"
	"wargaming/pkg/errno"
)

func MatchHandle(ctx context.Context, msg MatchData) (*Data, error) {

	WsConnManager.SetStatus(msg.UserId, Matching)

	var match *game.MatchResp
	var err error

	delay := constants.MatchInitialDelay
	retryCount := 0

	for match == nil || match.MatchedUserId == 0 {
		match, err = rpc.Match(ctx, &game.MatchReq{
			UserId: msg.UserId,
		})

		if err != nil {
			hlog.Error(err)
			return nil, err
		}

		time.Sleep(delay)
		retryCount++

		if delay < constants.MaxMatchRetryDelay {
			delay *= 2 // 延迟时间翻倍
		}

		if delay > constants.MaxMatchRetryDelay {
			return nil, errno.MatchTimeoutError
		}
	}

	WsConnManager.SetStatus(msg.UserId, Gaming)
	WsConnManager.SetOpponent(msg.UserId, match.MatchedUserId)

	return &Data{
		Action: MatchAction,
		Data: MatchData{
			UserId: match.MatchedUserId,
		},
	}, nil
}

func MoveHandle(ctx context.Context, msg MoveData) (*Data, error) {
	return &Data{
		Action: MoveAction,
		Data:   msg,
	}, nil
}
