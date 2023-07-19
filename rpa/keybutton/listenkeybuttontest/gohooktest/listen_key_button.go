package main

import (
	"fmt"
	hook "github.com/robotn/gohook"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/19 11:59
 * @Desc:
 */

// 监听按键
func ListenKeyButton() {
	
	hooks := hook.Start()
	
	defer hook.End()
	
	// 下面通过 `check_os_key.go` 检查不同操作系统的按键相关数据
	/* Windows
	# F11：Rawcode=122
	
	# F12：Rawcode=123
	*/
	/* Mac
	# F11：Rawcode=103
	
	# F12：Rawcode=111
	*/
	
	for ev := range hooks {
		
		//	监听键盘-按下
		if ev.Kind == hook.KeyDown {
			
			//	监听w
			if ev.Rawcode == 87 {
				// ...
				fmt.Println("监听w-按下")
			}
			
			//	监听F11
			if ev.Rawcode == 122 {
				// ...
				fmt.Println("监听F11-按下")
			}
			
			//	监听F12
			if ev.Rawcode == 123 {
				// ...
				fmt.Println("监听F12-按下")
				
			}
			
		}
		
		//	监听键盘-松开
		if ev.Kind == hook.KeyUp {
			
			//	监听w
			if ev.Rawcode == 87 {
				// ...
				fmt.Println("监听w-松开")
			}
			
			//	监听F11
			if ev.Rawcode == 122 {
				// ...
				fmt.Println("监听F11-松开")
			}
			
			//	监听F12
			if ev.Rawcode == 123 {
				// ...
				fmt.Println("监听F12-松开")
				
			}
			
		}
		
	}
	
}
