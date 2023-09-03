package signalutil

import (
	"os"
	"os/signal"
	"syscall"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/9/2 13:27
 * @Desc:
 */

var OsSigal chan os.Signal

func InitOsSigal() {
	// 优雅关闭
	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	OsSigal = make(chan os.Signal, 1)
	signal.Notify(OsSigal,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGKILL,
	)
}
