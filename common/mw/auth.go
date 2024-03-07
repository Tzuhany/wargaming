package mw

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"wargaming/common/errno"
	"wargaming/common/jwt"
	"wargaming/common/response"
)

func AuthToken() app.HandlerFunc {

	return func(ctx context.Context, c *app.RequestContext) {
		token := c.Request.Header.Get("Authorization")

		if len(token) == 0 {
			token = c.Request.Header.Get("Sec-WebSocket-Protocol")
		}

		if len(token) == 0 {
			hlog.CtxInfof(ctx, "token is not exist clientIP: %v\n", c.ClientIP())
			response.SendFailResponse(c, errno.AuthorizationFailedError)
			c.Abort()
			return
		}

		hlog.CtxInfof(ctx, "token: %v clientIP: %v\n", token, c.ClientIP())

		claims, err := jwt.CheckToken(token)

		if err != nil {
			hlog.CtxInfof(ctx, "token is invalid clientIP: %v\n", c.ClientIP())
			response.SendFailResponse(c, errno.AuthorizationFailedError)
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Set("id", claims.Id)

		c.Next(ctx)
	}
}
