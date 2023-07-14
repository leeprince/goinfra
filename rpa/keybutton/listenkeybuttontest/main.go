package main

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/12 21:09
 * @Desc:
 */

import (
	"github.com/micmonay/keybd_event"
)

func main() {
	// 创建一个键盘事件模拟器
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	// 设置快捷键为 F1
	kb.SetKeys(keybd_event.VK_F1)

	/*// 监听快捷键被按下的事件
	for {
		if kb.HasCTRLPressed() && kb.HasALTPressed() && kb.HasKeyPressed(keybd_event.VK_T) {
			fmt.Println("快捷键 Ctrl+Alt+T按下了")
		}
		time.Sleep(100 * time.Millisecond)
	}*/
}
