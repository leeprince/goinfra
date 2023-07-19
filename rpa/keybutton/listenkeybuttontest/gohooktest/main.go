package main

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/19 9:52
 * @Desc:
 */

import (
	"fmt"
	hook "github.com/robotn/gohook"
)

func main() {

	hooks := hook.Start()

	defer hook.End()

	for ev := range hooks {

		//	监听键盘弹起
		if ev.Kind == hook.KeyUp {

			//	监听F11
			if ev.Rawcode == 122 {
				//...
				fmt.Println("监听F11-松开")
			}

			//	监听F12
			if ev.Rawcode == 123 {
				//...
				fmt.Println("监听F12-松开")

			}

		}

	}

}
