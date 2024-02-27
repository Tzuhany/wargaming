package ws

import (
	"context"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"wargaming/cmd/interact/biz/pack"
)

type MessageProcessingCenter struct {
	handlers map[int]ActionHandler
}

func NewMessageProcessingCenter() *MessageProcessingCenter {
	return &MessageProcessingCenter{
		handlers: make(map[int]ActionHandler),
	}
}

func (mpc *MessageProcessingCenter) RegisterHandler(action int, handler ActionHandler) {
	mpc.handlers[action] = handler
}

func (mpc *MessageProcessingCenter) ProcessMessage(ctx context.Context, wsConn *Connection, msg *Message, userId int64) error {
	// 解析消息数据
	var requestData interface{}
	err := sonic.Unmarshal(msg.MsgData, &requestData)
	if err != nil {
		hlog.Error(err)
		return err
	}

	action := requestData.(map[string]interface{})["action"].(float64)

	handler, exists := mpc.handlers[int(action)]
	if !exists {
		return nil
	}

	if action != MatchAction {
		wsConn = WsConnManager.GetOpponentMetaData(userId).Conn
	}

	// 处理消息
	response, err := handler.Handle(ctx, requestData)
	if err != nil {
		hlog.Error(err)
		return err
	}

	response.Base = pack.BuildBaseResp(err)

	// 发送响应
	marshal, err := sonic.Marshal(response)
	if err != nil {
		hlog.Error(err)
		return err
	}

	wsConn.WriteMessage(&Message{
		MsgType: msg.MsgType,
		MsgData: marshal,
	})

	return nil
}
