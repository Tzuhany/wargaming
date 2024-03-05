package jwt

import (
	"common/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type UserClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func CreateToken(userID string, username string) (string, error) {
	expireTime := time.Now().Add(24 * 7 * time.Hour) // 过期时间为7天
	nowTime := time.Now()                            // 当前时间
	claims := UserClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: expireTime,
			},
			IssuedAt: &jwt.NumericDate{
				Time: nowTime,
			},
			Issuer: "wargaming",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Conf.JwtAuth.AccessSecret))
}

func CheckToken(token string) (*UserClaims, error) {
	response, err := jwt.ParseWithClaims(token, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Conf.JwtAuth.AccessSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if resp, ok := response.Claims.(*UserClaims); ok && response.Valid {
		return resp, nil
	}

	return nil, err
}
