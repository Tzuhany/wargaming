package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	user "user/kitex_gen/user"
	"user/pack"
	"user/service"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	resp = new(user.RegisterResp)

	registerResp, err := service.NewUserService(ctx).Register(req)
	if err != nil {
		klog.Error(err)
		resp.Base = pack.BuildBaseResp(err)
		return
	}

	resp.Base = pack.BuildBaseResp(nil)
	resp.UserId = registerResp.UserId
	resp.Token = registerResp.Token

	return
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	resp = new(user.LoginResp)

	loginResp, err := service.NewUserService(ctx).Login(req)
	if err != nil {
		klog.Error(err)
		resp.Base = pack.BuildBaseResp(err)
		return
	}

	resp.Base = pack.BuildBaseResp(nil)
	resp.UserId = loginResp.UserId
	resp.Token = loginResp.Token

	return

}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoReq) (resp *user.UserInfoResp, err error) {
	resp = new(user.UserInfoResp)

	userInfoResp, err := service.NewUserService(ctx).UserInfo(req)
	if err != nil {
		klog.Error(err)
		resp.Base = pack.BuildBaseResp(err)
		return
	}

	resp.Base = pack.BuildBaseResp(nil)
	resp.User = userInfoResp.User

	return
}
