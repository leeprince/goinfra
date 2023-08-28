package contextutil

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/leeprince/goinfra/consts"
	"github.com/leeprince/goinfra/utils/idutil"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/26 22:51
 * @Desc:
 */

const (
	GinContextSetHeaderKey = "GinContextSetHeaderKey"
)

func LogIdByGinContext(ginCtx *gin.Context) (logId string) {
	// 从当前 gin.Context 上下文中取
	value, ok := ginCtx.Get(GinContextSetHeaderKey)
	if ok {
		headContext, ok := value.(context.Context)
		if ok {
			readHeaderByContext, err := ReadHeaderByContext(headContext)
			if err == nil {
				return readHeaderByContext.LogId
			}
		}
	}
	
	// 从请求头中取 x-log-id
	logId = ginCtx.GetHeader(consts.HeaderXLogID)
	if logId == "" {
		logId = idutil.UniqIDV3()
	}
	
	// 重新将含有 logId 数据设置到上下文中
	header := &Header{
		LogId: logId,
	}
	haveHeaderCtx := WriteHeaderContext(ginCtx, header)
	
	// 重新将含有 logId 数据的上下文设置到 gin.Context 中
	ginCtx.Set(GinContextSetHeaderKey, haveHeaderCtx)
	
	return
}

func LogIdByContext(ctx *context.Context) (logId string) {
	header, err := ReadHeaderByContext(*ctx)
	if err == nil {
		return header.LogId
	}
	
	logId = idutil.UniqIDV3()
	header = &Header{
		LogId: logId,
	}
	*ctx = WriteHeaderContext(*ctx, header)
	return
}
