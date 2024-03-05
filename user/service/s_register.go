package service

import (
	"common/jwt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"user/infra/dal"
	"user/kitex_gen/user"
)

func (s *UserService) Register(req *user.RegisterReq) (*user.RegisterResp, error) {
	hashBytes, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	userID := uuid.New()

	userModel := &dal.User{
		Username: req.Username,
		Password: string(hashBytes),
		UserID:   userID.String(),
	}

	err := dal.CreateUser(s.ctx, userModel)
	if err != nil {
		klog.Error(err)
		return nil, err
	}

	token, err := jwt.CreateToken(userModel.UserID, userModel.Username)
	if err != nil {
		klog.Error(err)
		return nil, err
	}

	return &user.RegisterResp{
		UserId: userModel.UserID,
		Token:  token,
	}, nil
}
