package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	consul "github.com/kitex-contrib/registry-consul"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"net"
	"wargaming/cmd/user/infra"
	"wargaming/common/config"
	"wargaming/common/constants"
	"wargaming/common/tracer"
	"wargaming/common/utils"
	user "wargaming/kitex_gen/user/userservice"
)

var (
	listenAddr string // listen port
)

func Init() {
	config.Init(constants.UserServiceName)

	tracer.InitJaeger(constants.UserServiceName)

	infra.Init()

	// log
	klog.SetLevel(klog.LevelDebug)
}

func main() {

	Init()

	r, err := consul.NewConsulRegister(config.Consul.Addr)
	if err != nil {
		panic(err)
	}

	// get available port from config set
	for index, addr := range config.Service.Addr {
		if ok := utils.AddrCheck(addr); ok {
			listenAddr = addr
			break
		}

		if index == len(config.Service.Addr)-1 {
			klog.Fatal("not available port from config")
		}
	}

	addr, err := net.ResolveTCPAddr("tcp", listenAddr)
	if err != nil {
		panic(err)
	}

	svr := user.NewServer(
		new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: constants.UserServiceName,
		}),
		server.WithMuxTransport(),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithSuite(trace.NewDefaultServerSuite()),
		server.WithLimit(&limit.Option{
			MaxConnections: constants.MaxConnections,
			MaxQPS:         constants.MaxQPS,
		}),
	)

	if err = svr.Run(); err != nil {
		panic(err)
	}
}
