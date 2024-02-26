package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"testing"
	"wargaming/cmd/user/infra"
	"wargaming/config"
	"wargaming/kitex_gen/user"
)

func TestUserService_UserInfo(t *testing.T) {
	config.InitForTest()

	infra.Init()
	svc := NewUserService(context.Background())

	register, err := svc.UserInfo(&user.UserInfoReq{
		UserId: 4,
	})
	if err != nil {
		klog.Error(err)
	}

	fmt.Println(register)
}
