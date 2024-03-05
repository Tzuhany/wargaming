package service

import (
	"common/jwt"
	"common/pkg/errno"
	"github.com/cloudwego/kitex/pkg/klog"
	"golang.org/x/crypto/bcrypt"
	"user/infra/dal"
	"user/kitex_gen/user"
)

func (s *UserService) Login(req *user.LoginReq) (*user.LoginResp, error) {
	userModel, err := dal.GetUserByUsername(s.ctx, req.Username)

	if err != nil {
		return nil, err
	}

	if bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(req.Password)) != nil {
		return nil, errno.AuthorizationFailedError
	}

	token, err := jwt.CreateToken(userModel.UserID, req.Username)
	if err != nil {
		klog.Error(err)
		return nil, err
	}

	return &user.LoginResp{
		UserId: userModel.UserID,
		Token:  token,
	}, nil
}
