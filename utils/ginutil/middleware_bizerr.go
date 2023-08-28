package ginutil

import (
	"github.com/gin-gonic/gin"
	"github.com/leeprince/goinfra/perror"
	"github.com/leeprince/goinfra/plog"
	"github.com/leeprince/goinfra/utils/contextutil"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/27 14:42
 * @Desc:
 */

// 错误拦截器(中间件)：业务错误拦截器
func MiddlewareBizErr(code int32, message string) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			logId := contextutil.LogIdByGinContext(c)
			for _, err := range c.Errors {
				bizErr, ok := err.Err.(perror.BizErr)
				if !ok {
					// 如果err是系统错误(error)，则替换为bizerr.Error，并将原err打log
					plog.WithError(err).Error(logId, "系统错误")
					err.Err = perror.NewBizErr(code, message)
				}
				c.JSON(200, gin.H{"code": bizErr.GetCode(), "message": bizErr.GetMessage(), "log_id": logId})
				c.Abort()
			}
		}()
		
		c.Next()
	}
}
