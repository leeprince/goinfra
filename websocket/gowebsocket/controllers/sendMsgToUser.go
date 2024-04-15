package controllers

import (
	"github.com/gin-gonic/gin"
	"gowebsocket/logger"
	"gowebsocket/params"
	"gowebsocket/services"
)

type WebSocketMsg struct {
}

func (s WebSocketMsg) SendMsgToUser(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("SendMsgToUser异常退出: ", err)
		}
	}()
	var sendMsgToUserReq params.SendMsgToUserReq

	if err := c.ShouldBindJSON(&sendMsgToUserReq); err != nil {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "fail",
		})
		return
	}

	// 发送消息到已通过 wstoken 建立的 websocket 连接上
	wstoken := sendMsgToUserReq.WsToken
	data := sendMsgToUserReq.Data
	conn, ok := services.Manager.Clients[wstoken]
	if ok == false {
		c.JSON(200, gin.H{
			"code":    -2,
			"message": "connect not exist",
		})
		return
	}
	conn.Send <- []byte(data)

	c.JSON(200, gin.H{
		"code":    0,
		"message": "successful.",
	})
}
