package ws

import (
	"context"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/websocket"
	"github.com/mitchellh/mapstructure"
	"wargaming/cmd/interact/biz/model/interact"
	"wargaming/cmd/interact/biz/pack"
)

func Interact(ctx context.Context, c *app.RequestContext, req *interact.InteractReq) {
	var upgrader = websocket.HertzUpgrader{
		CheckOrigin: func(ctx *app.RequestContext) bool {
			return true
		},
		Subprotocols: []string{c.Request.Header.Get("Sec-WebSocket-Protocol")},
	}

	err := upgrader.Upgrade(c, func(conn *websocket.Conn) {
		wsConn := InitWsConnection(conn)

		WsConnManager.Put(req.UserId, &MetaData{
			Conn:   wsConn,
			Status: Free,
		})

		var msg *Message
		var data Data
		for {
			msg = wsConn.ReadMessage()

			err := sonic.Unmarshal(msg.MsgData, &data)
			if err != nil {
				hlog.Error(err)
				return
			}

			switch data.Action {
			case MatchAction:
				mp := data.Data.(map[string]interface{})
				m := MatchData{}

				err := mapstructure.Decode(mp, &m)
				if err != nil {
					hlog.Error(err)
					return
				}

				matchRet, err := MatchHandle(ctx, m)
				if err != nil {
					hlog.Error(err)
					return
				}

				matchRet.Base = pack.BuildBaseResp(err)

				marshal, err := sonic.Marshal(matchRet)
				if err != nil {
					hlog.Error(err)
					return
				}

				wsConn.WriteMessage(&Message{
					MsgType: msg.MsgType,
					MsgData: marshal,
				})

			case MoveAction:
				mp := data.Data.(map[string]interface{})
				m := MoveData{}

				err := mapstructure.Decode(mp, &m)
				if err != nil {
					hlog.Error(err)
					return
				}

				moveRet, err := MoveHandle(ctx, m)
				if err != nil {
					hlog.Error(err)
					return
				}

				moveRet.Base = pack.BuildBaseResp(err)

				marshal, err := sonic.Marshal(moveRet)
				if err != nil {
					hlog.Error(err)
					return
				}

				// 获取对手的连接
				opponentConn := WsConnManager.GetOpponentMetaData(req.UserId).Conn

				opponentConn.WriteMessage(&Message{
					MsgType: msg.MsgType,
					MsgData: marshal,
				})
			}
		}
	})
	if err != nil {
		hlog.Error(err)
	}
}
