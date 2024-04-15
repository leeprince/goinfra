package controllers

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/4/15 16:47
 * @Desc:	创建连接：通过 HTTP 协议升级到 WebSocket 协议；发起连接的请求格式为：ws://host:port/ws
 */

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gowebsocket/logger"
	"gowebsocket/services"
	"net/http"
)

type WebSocket struct {
}

var (
	upgrader = websocket.Upgrader{
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func (s WebSocket) WsHandler(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("WsHandler异常退出: ", err)
		}
	}()
	wstoken := c.Query("wstoken")
	
	if wstoken == "" {
		http.NotFound(c.Writer, c.Request)
		return
	}
	
	// 完成ws协议的握手操作
	// Upgrade: 升级将 HTTP 服务器连接升级到 WebSocket 协议。
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	
	// 可以添加用户信息验证
	client := &services.Client{
		ID:     services.Manager.CreatId(wstoken),
		Socket: conn,
		Send:   make(chan []byte),
	}
	services.Manager.Register <- client
	// 接收消息
	go client.Read()
	// 发送消息
	go client.Write()
	// 启动线程，心跳检测
	go client.Health()
}
