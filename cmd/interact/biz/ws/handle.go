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
	Handle(ctx context.Context, req *Message) (*Message, error)
}

type MatchHandler struct{}

func (h *MatchHandler) Handle(ctx context.Context, req *Message) (*Message, error) {

	var matchData MatchReq

	err := mapstructure.Decode(req.Data, &matchData)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	WsManager.Get(req.From).SetStatus(Matching)

	var match *game.MatchResp

	delay := constants.MatchInitialDelay
	retryCount := 0

	for match == nil || match.MatchedUserId == 0 {
		match, err = rpc.Match(ctx, &game.MatchReq{
			UserId: matchData.UserID,
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

	WsManager.Get(req.From).SetStatus(Gaming)

	return &Message{
		From:   req.From,
		To:     req.To,
		Action: MatchAction,
		Data: MatchResp{
			MatchedUserID: match.MatchedUserId,
		},
	}, nil
}

type MoveHandler struct{}

func (h *MoveHandler) Handle(ctx context.Context, req *Message) (*Message, error) {

	var moveData MoveData

	err := mapstructure.Decode(req, &moveData)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	return &Message{
		From:   req.From,
		To:     req.To,
		Action: MoveAction,
		Data:   moveData,
	}, nil
}
