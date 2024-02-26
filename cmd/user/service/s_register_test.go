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

func TestUserService_Register(t *testing.T) {

	config.InitForTest()

	infra.Init()
	svc := NewUserService(context.Background())

	register, err := svc.Register(&user.RegisterReq{
		Username: "yzh",
		Password: "123456",
	})
	if err != nil {
		klog.Error(err)
	}

	fmt.Println(register)
}
