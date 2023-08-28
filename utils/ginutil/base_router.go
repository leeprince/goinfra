package ginutil

import (
	"github.com/gin-gonic/gin"
	"github.com/leeprince/goinfra/plog"
	"github.com/leeprince/goinfra/utils/contextutil"
	"net/http"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/28 17:06
 * @Desc:
 */

func SetBaseRouter(engine *gin.Engine) {
	// 未定义路由时
	engine.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{})
	})
	
	// 路由方法不匹配时
	engine.NoMethod(func(context *gin.Context) {
		context.JSON(http.StatusMethodNotAllowed, gin.H{})
	})
	
	// 默认访问路径
	engine.Any("/", func(c *gin.Context) {
		c.Header("Server", "prince/1.0.0")
		c.JSON(200, gin.H{
			"code":    200,
			"message": "ok",
			"log_id":  contextutil.LogIdByGinContext(c),
			"data":    map[string]string{},
		})
		plog.Info("/")
	})
	
	engine.Any("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
		plog.Info("/ping")
	})
}
