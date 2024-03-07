package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"golang.org/x/crypto/bcrypt"
	"wargaming/cmd/user/infra/dal"
	"wargaming/common/errno"
	"wargaming/common/jwt"
	"wargaming/kitex_gen/user"
)

func (s *UserService) Login(ctx context.Context, req *user.LoginReq) (*user.LoginResp, error) {
	userModel, err := dal.FindUserByUsername(ctx, req.Username)

	if err != nil {
		return nil, err
	}

	if bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(req.Password)) != nil {
		return nil, errno.AuthorizationFailedError
	}

	token, err := jwt.CreateToken(userModel.UID, req.Username)
	if err != nil {
		klog.Error(err)
		return nil, err
	}

	return &user.LoginResp{
		UserId: userModel.UID,
		Token:  token,
	}, nil
}
