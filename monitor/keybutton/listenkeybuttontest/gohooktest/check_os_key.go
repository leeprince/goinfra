package main

import (
	"fmt"
	hook "github.com/robotn/gohook"
	"runtime"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/19 12:01
 * @Desc:
 */

// 检查当前操作系统按键的Rawcode
// ！！！注意：不同操作系统的按键`事件`和`Rawcode`是不一样的
// 		1. mac和 windows的大部分按键Rawcode是不一样的
// 		2. mac部分按键`事件`不存在。
// 			1. 如mac pro对于f1~f12不支持`ev.Kind == hook.KeyDown`事件，可能是自己的电脑mac pro的 f1~f12是 Touch Bar(触摸屏按键)所以不支持
func CheckOsKey() {
	fmt.Printf("---检查当前操作系统:`%s`的按键---\n\n", runtime.GOOS)
	evChan := hook.Start()
	defer hook.End()
	
	for ev := range evChan {
		// fmt.Println("hook: ", ev)
		
		if ev.Kind == hook.KeyDown ||
			ev.Kind == hook.KeyHold ||
			ev.Kind == hook.KeyUp {
			fmt.Println(">按键的事件数据 ev：", ev)
		}
	}
}
