package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
	"wargaming/cmd/user/infra/dal"
	"wargaming/common/jwt"
	"wargaming/kitex_gen/user"
)

func (s *UserService) Register(ctx context.Context, req *user.RegisterReq) (*user.RegisterResp, error) {
	hashBytes, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	uid := uuid.New().String()

	userModel := &dal.User{
		Username:   req.Username,
		Password:   string(hashBytes),
		UID:        uid,
		CreateTime: time.Now(),
	}

	err := dal.Insert(s.ctx, userModel)
	if err != nil {
		klog.Error(err)
		return nil, err
	}

	token, err := jwt.CreateToken(uid, req.Username)
	if err != nil {
		klog.Error(err)
		return nil, err
	}

	return &user.RegisterResp{
		UserId: uid,
		Token:  token,
	}, nil
}
