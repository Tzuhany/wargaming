package ws

import (
	"context"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/mitchellh/mapstructure"
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

func (mpc *MessageProcessingCenter) ProcessMessage(ctx context.Context, msg []byte) error {
	// 解析消息数据
	var requestData interface{}
	err := sonic.Unmarshal(msg, &requestData)
	if err != nil {
		hlog.Error(err)
		return err
	}

	var m Message

	err = mapstructure.Decode(requestData, &m)
	if err != nil {
		hlog.Error(err)
		return err
	}

	handler, exists := mpc.handlers[m.Action]
	if !exists {
		return nil
	}

	// 处理消息
	response, err := handler.Handle(ctx, &m)
	if err != nil {
		hlog.Error(err)
		return err
	}

	// 发送响应
	marshal, err := sonic.Marshal(response)
	if err != nil {
		hlog.Error(err)
		return err
	}

	// 获取发送目标
	to := requestData.(map[string]interface{})["to"].(float64)

	client := WsManager.Get(int64(to))

	client.WriteMessage(marshal)

	return nil
}
