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

// 监听windows操作系统按键
func WindowsOsListenKeyButton() {
	
	hooks := hook.Start()
	
	defer hook.End()
	
	for ev := range hooks {
		
		//	监听键盘-按下
		if ev.Kind == hook.KeyDown {
			
			//	监听w
			if ev.Rawcode == 87 {
				// ...
				fmt.Println("监听w-按下---ev:", ev)
			}
			
			//	监听F11
			if ev.Rawcode == 122 {
				// ...
				fmt.Println("监听F11-按下---ev:", ev)
			}
			
			//	监听F12
			if ev.Rawcode == 123 {
				// ...
				fmt.Println("监听F12-按下---ev:", ev)
				
			}
			
		}
		
		//	监听键盘-按住
		if ev.Kind == hook.KeyHold {
			
			//	监听w
			if ev.Rawcode == 87 {
				// ...
				fmt.Println("监听w-按住---ev:", ev)
			}
			
			//	监听F11
			if ev.Rawcode == 122 {
				// ...
				fmt.Println("监听F11-按住---ev:", ev)
			}
			
			//	监听F12
			if ev.Rawcode == 123 {
				// ...
				fmt.Println("监听F12-按住---ev:", ev)
				
			}
			
		}
		
		//	监听键盘-松开
		if ev.Kind == hook.KeyUp {
			
			//	监听w
			if ev.Rawcode == 87 {
				// ...
				fmt.Println("监听w-松开---ev:", ev)
			}
			
			//	监听F11
			if ev.Rawcode == 122 {
				// ...
				fmt.Println("监听F11-松开---ev:", ev)
			}
			
			//	监听F12
			if ev.Rawcode == 123 {
				// ...
				fmt.Println("监听F12-松开---ev:", ev)
				
			}
			
		}
		
	}
	
}
