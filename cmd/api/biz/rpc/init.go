package rpc

import (
	"wargaming/kitex_gen/user/userservice"
)

var (
	userClient userservice.Client
)

func Init() {
	InitUserRPC()
}
