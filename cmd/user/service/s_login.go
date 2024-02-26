package service

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"golang.org/x/crypto/bcrypt"
	"wargaming/cmd/user/infra/dal"
	"wargaming/kitex_gen/user"
	"wargaming/pkg/errno"
	"wargaming/pkg/jwt"
)

func (s *UserService) Login(req *user.LoginReq) (*user.LoginResp, error) {
	userModel, err := dal.GetUserByUsername(s.ctx, req.Username)

	if err != nil {
		return nil, err
	}

	if bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(req.Password)) != nil {
		return nil, errno.AuthorizationFailedError
	}

	token, err := jwt.CreateToken(int64(userModel.ID), req.Username)
	if err != nil {
		klog.Error(err)
		return nil, err
	}

	return &user.LoginResp{
		UserId: int64(userModel.ID),
		Token:  token,
	}, nil
}
