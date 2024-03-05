package dal

import (
	"common/pkg/errno"
	"context"
	"errors"
	"gorm.io/gorm"
)

type User struct {
	UserID   string `gorm:"not null;type:varchar(36);comment:'用户ID'"`
	Username string `gorm:"not null;type:varchar(20);comment:'用户名'"`
	Password string `gorm:"not null;type:varchar(255);comment:'用户密码'"`
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

func GetUserByID(ctx context.Context, userid string) (*User, error) {
	userResp := new(User)

	err := DB.WithContext(ctx).Where("user_id = ?", userid).First(&userResp).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.UserNotFoundError
		}
		return nil, err
	}

	return userResp, nil
}
