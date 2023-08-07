package main

import (
	"github.com/go-vgo/robotgo"
	"github.com/leeprince/goinfra/perror"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/4 00:28
 * @Desc:
 */

// 选择窗口
func SelectWindow() {
	// err = robotgo.ActivePid(33213)
	// ACTIVE_NAME 可以通过`ps -ef | grep sublime` 快速查找
	err = robotgo.ActiveName(ACTIVE_NAME)
	perror.ErrPanic(err)
}
