package router

import (
	"github.com/gin-gonic/gin"
	"gowebsocket/controllers"
	"net/http"
	"time"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	
	router.NoRoute(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    http.StatusNotFound,
			"message": "接口不存在",
			"params":  map[string]string{},
		})
	})
	
	router.NoMethod(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    http.StatusMethodNotAllowed,
			"message": "请求方式不允许",
			"params":  map[string]string{},
		})
	})
	
	router.Any("/", func(c *gin.Context) {
		c.Header("server", "gowebsocket service /1.0.0")
		c.JSON(200, gin.H{
			"message": "gowebsocket service " + time.Now().Format("2006-01-02 15:04:05"),
		})
	})
	
	api := router.Group("")
	{
		api.GET("/healthz", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
			return
		})
		
		api.GET("/ws", controllers.WebSocket{}.WsHandler)
		api.POST("/pushws", controllers.WebSocketMsg{}.SendMsgToUser)
	}
	
	return router
}
