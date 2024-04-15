package services

import (
	"github.com/gorilla/websocket"
	"github.com/spf13/cast"
	"gowebsocket/config"
	"gowebsocket/logger"
	"time"
)

// Client is a websocket client
type Client struct {
	ID     string
	Socket *websocket.Conn
	Send   chan []byte
}

func (c *Client) Read() {
	defer func() {
		Manager.Unregister <- c
		c.Socket.Close()
	}()
	defer func() {
		if err := recover(); err != nil {
			logger.Error("Read异常退出: ", err)
		}
	}()
	
	for {
		c.Socket.PongHandler()
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			defer Manager.Mux.Unlock()
			Manager.Unregister <- c
			break
		}
		logger.Info("读取到客户端的信息:", string(message))
		Manager.Broadcast <- message
		
	}
}

func (c *Client) Write() {
	defer func() {
		c.Socket.Close()
	}()
	defer func() {
		if err := recover(); err != nil {
			logger.Error("Write异常退出: ", err)
		}
	}()
	
	for {
		select {
		case message, ok := <-c.Send:
			Manager.Mux.Lock()
			if !ok {
				defer Manager.Mux.Unlock()
				c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			logger.Info("发送到客户端的信息:", string(message))
			err := c.Socket.WriteMessage(websocket.TextMessage, message)
			Manager.Mux.Unlock()
			if err != nil {
				logger.Error("发送到客户端的信息失败:", err)
			}
		}
	}
}

func (c *Client) Health() {
	var (
		err error
	)
	defer func() {
		if err := recover(); err != nil {
			logger.Error("Health异常退出: ", err)
		}
	}()
	for {
		Manager.Mux.Lock()
		if err = c.Socket.WriteMessage(websocket.TextMessage, []byte("pong")); err != nil {
			defer Manager.Mux.Unlock()
			return
		}
		Manager.Mux.Unlock()
		time.Sleep(cast.ToDuration(config.Config.Heartbeat) * time.Second)
	}
}
