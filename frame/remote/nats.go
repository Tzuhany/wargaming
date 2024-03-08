package remote

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/nats-io/nats.go"
	"wargaming/frame/game"
)

type NatsClient struct {
	serverId string
	conn     *nats.Conn
	readChan chan []byte
}

func (c *NatsClient) Run() error {
	var err error
	c.conn, err = nats.Connect(game.Conf.ServersConf.Nats.Url)
	if err != nil {
		klog.Errorf("connect nats server fail,err:%v", err)
		return err
	}
	go c.sub()
	return nil
}

func (c *NatsClient) Close() error {
	if c.conn != nil {
		c.conn.Close()
	}
	return nil
}

func NewNatsClient(serverId string, readChan chan []byte) *NatsClient {
	return &NatsClient{
		serverId: serverId,
		readChan: readChan,
	}
}

func (c *NatsClient) sub() {
	_, err := c.conn.Subscribe(c.serverId, func(msg *nats.Msg) {
		//收到的其他nats client发送的消息
		c.readChan <- msg.Data
	})
	if err != nil {
		klog.Errorf("nats sub err:%v", err)
	}
}

func (c *NatsClient) SendMsg(dst string, data []byte) error {
	if c.conn != nil {
		return c.conn.Publish(dst, data)
	}
	return nil
}
