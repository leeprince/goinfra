package main

import (
	"fmt"
	"github.com/leeprince/goinfra/consts"
	"github.com/leeprince/goinfra/utils/keyutil"
	hook "github.com/robotn/gohook"
	"log"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/20 22:28
 * @Desc:
 */

func ListenWindowAndMacpro() {
	hooks := hook.Start()
	defer hook.End()
	
	keyMap, err := keyutil.GetOsKeyButtonRawCode()
	if err != nil {
		log.Fatal("GetOsKeyButtonRawCode", err)
	}
	
	for ev := range hooks {
		if ev.Rawcode == keyMap.MustGet(consts.KeyEsc).Value() {
			//	监听按键-按下
			if ev.Kind == hook.KeyDown {
				fmt.Printf("监听按键:%s-按下.ev:%v\n", consts.KeyEsc, ev)
			}
			
			//	监听按键-按住
			if ev.Kind == hook.KeyHold {
				fmt.Printf("监听按键:%s-按住.ev:%v\n", consts.KeyEsc, ev)
			}
			
			//	监听按键-松开
			if ev.Kind == hook.KeyUp {
				fmt.Printf("监听按键:%s-松开.ev:%v\n", consts.KeyEsc, ev)
			}
		}
		
		if ev.Rawcode == keyMap.MustGet(consts.KeyF1).Value() {
			//	监听按键-按下
			if ev.Kind == hook.KeyDown {
				fmt.Printf("监听按键:%s-按下.ev:%v\n", consts.KeyF1, ev)
			}
			
			//	监听按键-按住
			if ev.Kind == hook.KeyHold {
				fmt.Printf("监听按键:%s-按住.ev:%v\n", consts.KeyF1, ev)
			}
			
			//	监听按键-松开
			if ev.Kind == hook.KeyUp {
				fmt.Printf("监听按键:%s-松开.ev:%v\n", consts.KeyF1, ev)
			}
		}
		
		if ev.Rawcode == keyMap.MustGet(consts.KeyF2).Value() {
			//	监听按键-按下
			if ev.Kind == hook.KeyDown {
				fmt.Printf("监听按键:%s-按下.ev:%v\n", consts.KeyF2, ev)
			}
			
			//	监听按键-按住
			if ev.Kind == hook.KeyHold {
				fmt.Printf("监听按键:%s-按住.ev:%v\n", consts.KeyF2, ev)
			}
			
			//	监听按键-松开
			if ev.Kind == hook.KeyUp {
				fmt.Printf("监听按键:%s-松开.ev:%v\n", consts.KeyF2, ev)
			}
		}
		
		if ev.Rawcode == keyMap.MustGet(consts.KeyF3).Value() {
			//	监听按键-按下
			if ev.Kind == hook.KeyDown {
				fmt.Printf("监听按键:%s-按下.ev:%v\n", consts.KeyF3, ev)
			}
			
			//	监听按键-按住
			if ev.Kind == hook.KeyHold {
				fmt.Printf("监听按键:%s-按住.ev:%v\n", consts.KeyF3, ev)
			}
			
			//	监听按键-松开
			if ev.Kind == hook.KeyUp {
				fmt.Printf("监听按键:%s-松开.ev:%v\n", consts.KeyF3, ev)
			}
		}
		
	}
}
