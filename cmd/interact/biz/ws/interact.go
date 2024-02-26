package ws

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/websocket"
	"strconv"
	"time"
	"wargaming/cmd/interact/biz/rpc"
	"wargaming/kitex_gen/game"
)

func Interact(ctx context.Context, c *app.RequestContext) {
	var upgrader = websocket.HertzUpgrader{
		CheckOrigin: func(ctx *app.RequestContext) bool {
			return true
		},
		Subprotocols: []string{c.Request.Header.Get("Sec-WebSocket-Protocol")},
	}

	err := upgrader.Upgrade(c, func(conn *websocket.Conn) {
		wsConn := InitWsConnection(conn)

		var data []byte

		for {
			data = wsConn.ReadMessage()

			userId, err := strconv.ParseInt(string(data), 10, 64)
			if err != nil {
				hlog.Error(err)
				return
			}

			var match *game.MatchResp

			for match == nil || match.MatchedUserId == 0 {
				match, err = rpc.Match(ctx, &game.MatchReq{
					UserId: userId,
				})
				time.Sleep(1 * time.Second)
			}

			hlog.Info("match success! -> ", match.MatchedUserId)

			wsConn.WriteMessage(&Message{
				dataType: websocket.TextMessage,
				data:     []byte(strconv.FormatInt(match.MatchedUserId, 10)),
			})
		}
	})
	if err != nil {
		hlog.Error(err)
		return
	}
}
