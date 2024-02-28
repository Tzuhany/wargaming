package ws

import (
	"github.com/hertz-contrib/websocket"
	"sync"
)

const (
	Matching = iota
	Free
	Gaming
)

type Client struct {
	status    int
	wsConn    *websocket.Conn
	inChan    chan []byte
	outChan   chan []byte
	closeChan chan struct{}
	isClosed  bool
	mutex     sync.Mutex
}

func NewClient(wsConn *websocket.Conn) *Client {
	client := &Client{
		status:    Free,
		wsConn:    wsConn,
		inChan:    make(chan []byte),
		outChan:   make(chan []byte),
		closeChan: make(chan struct{}),
	}

	go client.readLoop()
	go client.writeLoop()

	return client
}

func (c *Client) ReadMessage() []byte {
	var msg []byte

	select {
	case msg = <-c.inChan:
	}

	return msg
}

func (c *Client) WriteMessage(msg []byte) {

	select {
	case c.outChan <- msg:
	}
}

func (c *Client) Close() {
	c.wsConn.Close()

	c.mutex.Lock()
	if !c.isClosed {
		close(c.closeChan)
		c.isClosed = true
	}
	c.mutex.Unlock()
}

func (c *Client) readLoop() {
	for {
		_, m, err := c.wsConn.ReadMessage()
		if err != nil {
			goto ERR
		}

		select {
		case c.inChan <- m:
		case <-c.closeChan:
			goto ERR
		}
	}

ERR:
	c.Close()
}

func (c *Client) writeLoop() {
	var msg []byte
	var err error

	for {
		msg = <-c.outChan
		if err = c.wsConn.WriteMessage(websocket.TextMessage, msg); err != nil {
			goto ERR
		}
	}

ERR:
	c.Close()
}

func (c *Client) SetStatus(status int) {
	c.status = status
}
