package ws

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/websocket"
	"wargaming/cmd/interact/biz/model/interact"
)

func Interact(ctx context.Context, c *app.RequestContext, req *interact.InteractReq) {
	var upgrader = websocket.HertzUpgrader{
		CheckOrigin: func(ctx *app.RequestContext) bool {
			return true
		},
		Subprotocols: []string{c.Request.Header.Get("Sec-WebSocket-Protocol")},
	}

	processingCenter := NewMessageProcessingCenter()
	processingCenter.RegisterHandler(MatchAction, &MatchHandler{})
	processingCenter.RegisterHandler(MoveAction, &MoveHandler{})

	err := upgrader.Upgrade(c, func(conn *websocket.Conn) {
		wsConn := InitWsConnection(conn)

		WsConnManager.Put(req.UserId, &MetaData{
			Conn:   wsConn,
			Status: Free,
		})

		var msg *Message
		for {
			msg = wsConn.ReadMessage()

			if err := processingCenter.ProcessMessage(ctx, wsConn, msg, req.UserId); err != nil {
				return
			}
		}
	})
	if err != nil {
		hlog.Error(err)
	}
}
