package ginutil

import (
	"github.com/gin-gonic/gin"
	"github.com/leeprince/goinfra/perror"
	"github.com/leeprince/goinfra/plog"
	"net/http"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/6 10:16
 * @Desc:
 */

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/6 09:43
 * @Desc:
 */

func ResponseSuccess(ctx *gin.Context, logID string, data any) {
	ctx.JSON(http.StatusOK, BaseResponse{
		BaseCommonResponse: BaseCommonResponse{
			Code:    perror.Success.GetCode(),
			Message: perror.Success.GetMessage(),
			LogID:   logID,
		},
		Data: data,
	})
	return
}

func ResponseError(ctx *gin.Context, logID string, err error) {
	plog.LogID(logID).WithError(err).WithField("method", "ResponseError").Error("request")
	
	code := perror.Fail.GetCode()
	codeMessage := perror.Fail.GetMessage()
	
	bizErr, ok := err.(perror.BizErr)
	if ok {
		code = bizErr.GetCode()
		codeMessage = bizErr.GetMessage()
	}
	
	ctx.JSON(http.StatusOK, BaseResponse{
		BaseCommonResponse: BaseCommonResponse{
			Code:    code,
			Message: codeMessage,
			LogID:   logID,
		},
		Data: nil,
	})
	return
}
