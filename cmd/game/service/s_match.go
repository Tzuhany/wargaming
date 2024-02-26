package service

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
	"math"
	"strconv"
	"wargaming/cmd/game/rpc"
	"wargaming/cmd/user/infra/cache"
	"wargaming/kitex_gen/game"
	"wargaming/kitex_gen/user"
)

func (s *GameService) Match(req *game.MatchReq) (*game.MatchResp, error) {
	resp := new(game.MatchResp)

	// 获取该玩家的详细信息
	userInfo, err := rpc.UserInfo(s.ctx, &user.UserInfoReq{
		UserId: req.UserId,
	})
	if err != nil || userInfo == nil {
		klog.Error(err)
		return nil, err
	}

	// 将该用户放入匹配池
	err = cache.Rdb.ZAdd(s.ctx, "match", redis.Z{
		Score:  float64(userInfo.User.Rank),
		Member: userInfo.User.Id,
	}).Err()
	if err != nil {
		klog.Error(err)
		return nil, err
	}

	// 获取匹配池中指定范围内的玩家
	rankRange := userInfo.User.Rank - 100
	if rankRange < 0 {
		rankRange = 0
	}
	rankRangeEnd := userInfo.User.Rank + 100

	matchedUsers, err := cache.Rdb.ZRangeByScoreWithScores(s.ctx, "match", &redis.ZRangeBy{
		Min: fmt.Sprintf("%d", rankRange),
		Max: fmt.Sprintf("%d", rankRangeEnd),
	}).Result()
	if err != nil {
		klog.Error(err)
		return nil, err
	}

	if len(matchedUsers) > 0 {

		minGap := math.MaxFloat64

		for _, m := range matchedUsers {
			matchedUserId, err := strconv.ParseInt(m.Member.(string), 10, 64)
			if err != nil {
				return nil, err
			}

			if math.Abs(float64(userInfo.User.Rank)-m.Score) < minGap && matchedUserId != userInfo.User.Id {
				minGap = math.Abs(float64(userInfo.User.Rank) - m.Score)
				resp.MatchedUserId = matchedUserId
			}
		}

		err = cache.Rdb.ZRem(s.ctx, "match", userInfo.User.Id, resp.MatchedUserId).Err()
		if err != nil {
			klog.Error(err)
			return nil, err
		}

		return resp, nil
	}

	return resp, nil
}
