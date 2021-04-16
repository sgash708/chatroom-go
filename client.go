package main

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// writeWait Allow writing_time
	writeWait = 10 * time.Second
	// pongWait Allow reading_time
	pongWait = 60 * time.Second
	// pingPeriod send ping time
	pingPeriod = (pongWait * 9) / 10
	// maxMessageSize
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// connection is middleman
type connection struct {
	ws   *websocket.Conn
	send chan []byte
}

// readPump
func (s subscription) readPump() {
	c := s.conn
	defer func() {
		h.unregister <- s
		// c is main.go's variable
		c.ws.Close()
	}()
	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(ponWait))
	c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, msg, err := c.ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
		m := message{msg, s.room}
		h.broadcast <- m
	}
}

func (c *connection) write(mt int, payload []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(mt, payload)
}
