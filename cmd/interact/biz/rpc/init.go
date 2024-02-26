package rpc

import "wargaming/kitex_gen/game/gameservice"

var (
	gameClient gameservice.Client
)

func Init() {
	InitGameRPC()
}
