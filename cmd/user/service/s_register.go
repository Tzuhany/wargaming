package service

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"golang.org/x/crypto/bcrypt"
	"wargaming/cmd/user/infra/dal"
	"wargaming/kitex_gen/user"
	"wargaming/pkg/jwt"
)

func (s *UserService) Register(req *user.RegisterReq) (*user.RegisterResp, error) {
	hashBytes, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	userModel := &dal.User{
		Username: req.Username,
		Password: string(hashBytes),
	}

	err := dal.CreateUser(s.ctx, userModel)
	if err != nil {
		klog.Error(err)
		return nil, err
	}

	token, err := jwt.CreateToken(int64(userModel.ID), req.Username)
	if err != nil {
		klog.Error(err)
		return nil, err
	}

	return &user.RegisterResp{
		UserId: int64(userModel.ID),
		Token:  token,
	}, nil
}
