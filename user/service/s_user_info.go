package service

import (
	"user/infra/dal"
	"user/kitex_gen/user"
)

func (s *UserService) UserInfo(req *user.UserInfoReq) (*user.UserInfoResp, error) {
	userModel, err := dal.GetUserByID(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &user.UserInfoResp{
		User: &user.User{
			UserId:   userModel.UserID,
			Username: userModel.Username,
		},
	}, nil
}
