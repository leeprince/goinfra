package ginutil

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/leeprince/goinfra/plog"
	"github.com/leeprince/goinfra/utils/byteutil"
	"github.com/leeprince/goinfra/utils/contextutil"
	"github.com/leeprince/goinfra/utils/jsonutil"
	"io"
	"net/http"
	"strings"
	"sync"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/27 14:45
 * @Desc:
 */

var bufferPool = sync.Pool{
	New: func() interface{} { return new(bytes.Buffer) },
}

// respWriter 定义一个存储响应内容的结构体:在结构体中封装了gin的 ResponseWriter，然后在重写的Write方法中，首先向bytes.Buffer中写数据，然后响应
type respWriter struct {
	gin.ResponseWriter
	buf *bytes.Buffer
}

// Write 读取响应数据
func (w *respWriter) Write(b []byte) (int, error) {
	w.buf.Write(b)
	return w.ResponseWriter.Write(b)
}

// 访问日志中间件
func MiddlewareAccessLogAndLogId() gin.HandlerFunc {
	return func(c *gin.Context) {
		logId := contextutil.LogIdByGinContext(c)
		plogEntry := plog.LogID(logId)
		
		route := c.Request.URL.Path
		
		// 无需记录访问日志
		if strings.Contains(route, "ping") ||
			strings.Contains(route, "healthz") ||
			strings.Contains(route, "metrics") {
			c.Next()
			return
		}
		
		method := c.Request.Method
		if method == http.MethodOptions {
			c.Next()
			return
		}
		
		// post 的请求体数据
		reqBytes, err := c.GetRawData()
		if err != nil {
			plogEntry.WithError(err).Info(logId, "读取请求参数失败")
		}
		// get 的请求url参数
		params := c.Request.URL.Query()
		
		var req interface{}
		if len(reqBytes) > 0 {
			req = make(map[string]interface{})
			_ = jsonutil.JsoniterCompatible.Unmarshal(reqBytes, &req)
		} else {
			req = params
		}
		
		plogEntry.WithField("ip", c.ClientIP()).
			WithField("method", method).
			WithField("route", route).
			WithField("params", params).
			WithField("req", req).
			Info("接受请求")
		
		buf := bufferPool.Get().(*bytes.Buffer)
		buf.Reset()
		defer bufferPool.Put(buf)
		
		// 重新将请求参数放到 gin.Context 的 Request.Body 中
		buf.Write(reqBytes)
		c.Request.Body = io.NopCloser(buf)
		
		// 修改 gin context response writer
		w := &respWriter{
			ResponseWriter: c.Writer,
			buf:            buf,
		}
		c.Writer = w
		
		// 进入下一层
		c.Next()
		
		respData := w.buf.Bytes()
		plogEntry.WithField("resp", byteutil.Bytes2String(respData)).
			Info("返回响应")
	}
}
