package ws

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/mitchellh/mapstructure"
	"time"
	"wargaming/cmd/interact/biz/rpc"
	"wargaming/kitex_gen/game"
	"wargaming/pkg/constants"
	"wargaming/pkg/errno"
)

type ActionHandler interface {
	Handle(ctx context.Context, req interface{}) (*Data, error)
}

type MatchHandler struct{}

func (h *MatchHandler) Handle(ctx context.Context, req interface{}) (*Data, error) {

	var matchData MatchData

	err := mapstructure.Decode(req, &matchData)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	WsConnManager.SetStatus(matchData.UserId, Matching)

	var match *game.MatchResp

	delay := constants.MatchInitialDelay
	retryCount := 0

	for match == nil || match.MatchedUserId == 0 {
		match, err = rpc.Match(ctx, &game.MatchReq{
			UserId: matchData.UserId,
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

	WsConnManager.SetStatus(matchData.UserId, Gaming)
	WsConnManager.SetOpponent(matchData.UserId, match.MatchedUserId)

	return &Data{
		Action: MatchAction,
		Data: MatchData{
			UserId: match.MatchedUserId,
		},
	}, nil
}

type MoveHandler struct{}

func (h *MoveHandler) Handle(ctx context.Context, req interface{}) (*Data, error) {

	var moveData MoveData

	err := mapstructure.Decode(req, &moveData)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	return &Data{
		Action: MoveAction,
		Data:   moveData,
	}, nil
}
