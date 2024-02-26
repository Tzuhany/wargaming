// Code generated by hertz generator.

package main

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/hertz-contrib/registry/consul"
	"wargaming/cmd/interact/biz/rpc"
	"wargaming/config"
	"wargaming/pkg/constants"
	"wargaming/pkg/mw"
	"wargaming/pkg/tracer"
	"wargaming/pkg/utils"
)

var (
	listenAddr string // listen port
)

func Init() {
	config.Init(constants.InteractService)

	tracer.InitJaeger(constants.InteractService)

	rpc.Init()

	// log
	hlog.SetLevel(hlog.LevelDebug)
}

func main() {

	Init()

	consulClient, err := consulapi.NewClient(consulapi.DefaultConfig())
	if err != nil {
		panic(err)
	}
	c := consul.NewConsulRegister(consulClient)

	// get available port from config set
	for index, addr := range config.Service.AddrList {
		if ok := utils.AddrCheck(addr); ok {
			listenAddr = addr
			break
		}

		if index == len(config.Service.AddrList)-1 {
			hlog.Fatal("not available port from config")
		}
	}

	fmt.Println(utils.NewNetAddr("tcp", listenAddr))

	r := server.New(
		server.WithHostPorts(listenAddr),
		server.WithHandleMethodNotAllowed(true),
		server.WithMaxRequestBodySize(1<<31),
		server.WithRegistry(c, &registry.Info{
			ServiceName: constants.InteractService,
			Addr:        utils.NewNetAddr("tcp", listenAddr),
			Weight:      10,
			Tags:        nil,
		}),
	)

	r.Use(
		mw.Recovery(),
	)

	register(r)
	r.Spin()
}
