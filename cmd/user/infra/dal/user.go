package dal

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"wargaming/pkg/errno"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null;type:varchar(20);comment:'用户名'"`
	Password string `gorm:"not null;type:varchar(255);comment:'用户密码'"`
	Rank     int    `gorm:"not null;type:uint;default:0;comment:'用户等级'"`
}

func CreateUser(ctx context.Context, user *User) error {
	userResp := new(User)

	err := DB.WithContext(ctx).Where("username = ?", user.Username).First(&userResp).Error

	if err == nil {
		return errno.UserExistedError
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if err := DB.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}

	return nil
}

func GetUserByUsername(ctx context.Context, username string) (*User, error) {
	userResp := new(User)

	err := DB.WithContext(ctx).Where("username = ?", username).First(&userResp).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.UserNotFoundError
		}
		return nil, err
	}

	return userResp, nil
}

func GetUserByID(ctx context.Context, userid int64) (*User, error) {
	userResp := new(User)

	err := DB.WithContext(ctx).Where("id = ?", userid).First(&userResp).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.UserNotFoundError
		}
		return nil, err
	}

	return userResp, nil
}
