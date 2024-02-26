package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/retry"
	consul "github.com/kitex-contrib/registry-consul"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"wargaming/config"
	"wargaming/kitex_gen/user"
	"wargaming/kitex_gen/user/userservice"
	"wargaming/pkg/constants"
	"wargaming/pkg/errno"
)

func InitUserRPC() {
	r, err := consul.NewConsulResolver(config.Consul.Addr)
	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		constants.UserServiceName,
		client.WithMuxConnection(constants.MuxConnection),
		client.WithRPCTimeout(constants.RPCTimeout),
		client.WithConnectTimeout(constants.ConnectTimeout),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
		client.WithSuite(trace.NewDefaultClientSuite()),
		client.WithLoadBalancer(loadbalance.NewWeightedRoundRobinBalancer()),
	)

	if err != nil {
		panic(err)
	}

	userClient = c
}

func UserInfo(ctx context.Context, req *user.UserInfoReq) (*user.UserInfoResp, error) {
	resp, err := userClient.UserInfo(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}

	return resp, nil
}
