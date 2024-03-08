package connector

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"wargaming/frame/game"
	"wargaming/frame/net"
	"wargaming/frame/remote"
)

type Connector struct {
	isRunning bool
	wsManager *net.Manager     // 管理 websocket 连接
	handlers  net.LogicHandler // 逻辑处理器，保存所有的业务逻辑处理方法
	remoteCli remote.Client    // 用于和 nats 通信
}

func Default() *Connector {
	return &Connector{
		handlers: make(net.LogicHandler),
	}
}

func (c *Connector) Run(serverId string) {
	if !c.isRunning {
		//启动websocket和nats
		c.wsManager = net.NewManager()
		c.wsManager.ConnectorHandlers = c.handlers
		//启动nats nats server不会存储消息
		c.remoteCli = remote.NewNatsClient(serverId, c.wsManager.RemoteReadChan)
		c.remoteCli.Run()
		c.wsManager.RemoteCli = c.remoteCli
		c.Serve(serverId)
	}
}

func (c *Connector) Close() {
	if c.isRunning {
		//关闭websocket和nats
		c.wsManager.Close()
	}
}

func (c *Connector) Serve(serverId string) {
	klog.Infof("run connector:%v", serverId)
	c.wsManager.ServerId = serverId
	connectorConfig := game.Conf.GetConnector(serverId)
	if connectorConfig == nil {
		klog.Fatalf("no connector config found")
	}
	addr := fmt.Sprintf("%s:%d", connectorConfig.Host, connectorConfig.ClientPort)
	c.isRunning = true
	c.wsManager.Run(addr)
}

func (c *Connector) RegisterHandler(handlers net.LogicHandler) {
	c.handlers = handlers
}
