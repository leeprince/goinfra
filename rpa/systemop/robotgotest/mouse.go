package main

import (
	"github.com/go-vgo/robotgo"
	"github.com/leeprince/goinfra/perror"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/4 00:30
 * @Desc:
 */

// 鼠标操作
func Mouse() {
	err = robotgo.ActiveName(ACTIVE_NAME)
	perror.ErrPanic(err)
	
	robotgo.MoveSmooth(100, 200, 1.0, 2.0, 2000)
	
	robotgo.MouseSleep = 2000
	robotgo.MoveClick(100, 400)
	
	robotgo.MouseSleep = 2000
	robotgo.MoveClick(200, 400, "right", true)
	
	robotgo.MovesClick(200, 600, "right", true)
}
