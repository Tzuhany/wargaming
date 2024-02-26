package service

import (
	"wargaming/cmd/user/infra/dal"
	"wargaming/kitex_gen/user"
)

func (s *UserService) UserInfo(req *user.UserInfoReq) (*user.UserInfoResp, error) {
	userModel, err := dal.GetUserByID(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &user.UserInfoResp{
		User: &user.User{
			Id:       int64(userModel.ID),
			Username: userModel.Username,
			Rank:     int64(userModel.Rank),
		},
	}, nil
}
