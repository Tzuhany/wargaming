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

func TestUserService_Login(t *testing.T) {
	config.InitForTest()

	infra.Init()
	svc := NewUserService(context.Background())

	login, err := svc.Login(&user.LoginReq{
		Username: "yzh",
		Password: "123",
	})
	if err != nil {
		klog.Error(err)
	}

	fmt.Println(login)
}
