package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/leeprince/goinfra/perror"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/4 00:30
 * @Desc:
 */

// 屏幕操作
func Screen() {
	err := robotgo.ActiveName(ACTIVE_NAME)
	perror.ErrPanic(err)
	
	bit := robotgo.CaptureScreen()
	fmt.Println("abitMap...", bit)
	gbit := robotgo.ToBitmap(bit)
	fmt.Println("bitmap...", gbit.Width)
	
	gbitMap := robotgo.CaptureGo()
	fmt.Println("Go CaptureScreen...", gbitMap.Width)
	fmt.Println("...", gbitMap.Width, gbitMap.BytesPerPixel)
	
	img := robotgo.CaptureImg()
	robotgo.Save(img, "save.png")
}
