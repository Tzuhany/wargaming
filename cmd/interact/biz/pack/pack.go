package pack

import (
	"github.com/pkg/errors"
	"wargaming/cmd/interact/biz/model/interact"
	"wargaming/pkg/errno"
)

func BuildBaseResp(err error) *interact.BaseResp {
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

func baseResp(err errno.ErrNo) *interact.BaseResp {
	return &interact.BaseResp{
		Code: err.ErrorCode,
		Msg:  err.ErrorMsg,
	}
}
