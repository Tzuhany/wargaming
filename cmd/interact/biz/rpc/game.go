package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/retry"
	consul "github.com/kitex-contrib/registry-consul"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"wargaming/config"
	"wargaming/kitex_gen/game"
	"wargaming/kitex_gen/game/gameservice"
	"wargaming/pkg/constants"
	"wargaming/pkg/errno"
)

func InitGameRPC() {
	r, err := consul.NewConsulResolver(config.Consul.Addr)
	if err != nil {
		panic(err)
	}

	c, err := gameservice.NewClient(
		constants.GameServiceName,
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

	gameClient = c
}

func Match(ctx context.Context, req *game.MatchReq) (*game.MatchResp, error) {
	resp, err := gameClient.Match(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}

	return resp, nil
}

func Move(ctx context.Context, req *game.MoveReq) (*game.MoveResp, error) {
	resp, err := gameClient.Move(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}

	return resp, nil
}
