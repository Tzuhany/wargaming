package main

import (
	"common/config"
	"common/pkg/constants"
	"common/pkg/utils"
	"common/tracer"
	"flag"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	consul "github.com/kitex-contrib/registry-consul"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"net"
	user "user/kitex_gen/user/userservice"
)

var (
	listenAddr string // listen port
	configFile = flag.String("config", "config.yaml", "config file")
)

func Init() {
	flag.Parse()

	config.InitConfig(*configFile)

	tracer.InitJaeger(config.Conf.Service.Name)

	// log
	klog.SetLevel(klog.Level(config.Conf.Log.Level))
}

func main() {
	Init()

	r, err := consul.NewConsulRegister(config.Conf.Consul.Addr)
	if err != nil {
		panic(err)
	}

	for index, addr := range config.Conf.Service.Addr {
		if ok := utils.AddrCheck(addr); ok {
			listenAddr = addr
			break
		}

		if index == len(config.Conf.Service.Addr)-1 {
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
			ServiceName: config.Conf.Service.Name,
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
