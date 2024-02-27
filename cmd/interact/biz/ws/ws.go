package ws

import (
	"github.com/hertz-contrib/websocket"
	"sync"
)

type Connection struct {
	wsConn    *websocket.Conn
	inChan    chan *Message
	outChan   chan *Message
	closeChan chan struct{}
	isClosed  bool
	mutex     sync.Mutex
}

type Message struct {
	MsgType int
	MsgData []byte
}

func InitWsConnection(wsConn *websocket.Conn) *Connection {
	conn := &Connection{
		wsConn:    wsConn,
		inChan:    make(chan *Message),
		outChan:   make(chan *Message),
		closeChan: make(chan struct{}),
	}

	go conn.readLoop()
	go conn.writeLoop()

	return conn
}

func (conn *Connection) ReadMessage() *Message {
	var msg *Message

	select {
	case msg = <-conn.inChan:
	}

	return msg
}

func (conn *Connection) WriteMessage(msg *Message) {

	select {
	case conn.outChan <- msg:
	}
}

func (conn *Connection) Close() {
	conn.wsConn.Close()

	conn.mutex.Lock()
	if !conn.isClosed {
		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock()
}

func (conn *Connection) readLoop() {
	for {
		mt, m, err := conn.wsConn.ReadMessage()
		if err != nil {
			goto ERR
		}

		msg := &Message{
			MsgType: mt,
			MsgData: m,
		}

		select {
		case conn.inChan <- msg:
		case <-conn.closeChan:
			goto ERR
		}
	}

ERR:
	conn.Close()
}

func (conn *Connection) writeLoop() {
	var msg *Message
	var err error

	for {
		msg = <-conn.outChan
		if err = conn.wsConn.WriteMessage(msg.MsgType, msg.MsgData); err != nil {
			goto ERR
		}
	}

ERR:
	conn.Close()
}
