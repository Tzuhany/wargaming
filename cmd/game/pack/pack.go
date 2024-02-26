package pack

import (
	"errors"
	"wargaming/kitex_gen/game"
	"wargaming/pkg/errno"
)

func BuildBaseResp(err error) *game.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}

	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceError.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *game.BaseResp {
	return &game.BaseResp{
		Code: err.ErrorCode,
		Msg:  err.ErrorMsg,
	}
}
