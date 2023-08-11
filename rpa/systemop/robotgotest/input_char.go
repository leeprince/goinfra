package main

import (
	"github.com/go-vgo/robotgo"
	"github.com/leeprince/goinfra/perror"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/4 00:29
 * @Desc:
 */

// 输入字符
func InputChar() {
	err = robotgo.ActiveName(ACTIVE_NAME)
	perror.ErrPanic(err)
	
	time.Sleep(time.Second * 2)
	
	robotgo.TypeStr("Hello, world!")
}
