package ginservice

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/leeprince/goinfra/utils/fileutil"
	"gitlab.yewifi.com/golden-cloud/common"
	"gitlab.yewifi.com/golden-cloud/common/gclog"
	"io"
	"net/http"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/1/18 18:31
 * @Desc:
 */

// 接收前端请求的文件url，响应并设置请求头为可以直接下载的文件
func DownloadFile(ctx *gin.Context) {
	logID := common.LogIdByCtx(ctx)
	logEntry := gclog.WithField("method", "DownloadFile")
	logEntry.Info(logID, "info")

	// 获取 URL 参数
	fileUrl := ctx.Query("file_url")
	if fileUrl == "" {
		logEntry.Error(logID, "file_url empty")
		ctx.String(http.StatusOK, `解析请求参数错误`)
		return
	}

	// 2. 发起GET请求到指定URL获取数据
	fileResp, err := http.Get(fileUrl)
	if err != nil {
		logEntry.WithError(err).Error(logID, "Get err")
		ctx.String(http.StatusOK, `请求文件 url 失败`)
		return
	}
	defer fileResp.Body.Close()

	// 将数据转换为字节流
	data, err := io.ReadAll(fileResp.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filePathName, err := fileutil.GetPathAndName(fileUrl)
	if err != nil {
		logEntry.WithError(err).Error(logID, "GetPathAndName err")
		ctx.String(http.StatusOK, `获取 url 文件名失败`)
		return
	}

	// 设置响应头，以便浏览器识别为下载流
	// 设置响应头以触发文件下载
	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filePathName)) // 在这里设置文件名

	// 将字节流写入响应体
	ctx.Data(http.StatusOK, "application/octet-stream", data) // 设置内容类型，这里是通用二进制流

	logEntry.Info(logID, "response")

	return
}
