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

var (
	EscKeyUpCh chan struct{}
	F1KeyUpCh  chan struct{}
	F2KeyUpCh  chan struct{}
	F3KeyUpCh  chan struct{}
)

func ListenWindowAndMacproChannelNotice() {
	hooks := hook.Start()
	defer hook.End()
	
	keyMap, err := keyutil.GetOsKeyButtonRawCode()
	if err != nil {
		log.Fatal("GetOsKeyButtonRawCode", err)
	}
	
	for ev := range hooks {
		if ev.Rawcode == keyMap.MustGet(consts.KeyEsc).Value() &&
			ev.Kind == hook.KeyUp {
			EscKeyUpCh <- struct{}{}
			
			//	监听按键-松开
			fmt.Printf("监听按键:%s-松开.ev:%v\n", consts.KeyEsc, ev)
		}
		
		if ev.Rawcode == keyMap.MustGet(consts.KeyF1).Value() &&
			ev.Kind == hook.KeyUp {
			F1KeyUpCh <- struct{}{}
			
			//	监听按键-松开
			fmt.Printf("监听按键:%s-松开.ev:%v\n", consts.KeyF1, ev)
		}
		
		if ev.Rawcode == keyMap.MustGet(consts.KeyF2).Value() &&
			ev.Kind == hook.KeyUp {
			F2KeyUpCh <- struct{}{}
			
			//	监听按键-松开
			fmt.Printf("监听按键:%s-松开.ev:%v\n", consts.KeyF2, ev)
		}
		
		if ev.Rawcode == keyMap.MustGet(consts.KeyF3).Value() &&
			ev.Kind == hook.KeyUp {
			F3KeyUpCh <- struct{}{}
			
			//	监听按键-松开
			fmt.Printf("监听按键:%s-松开.ev:%v\n", consts.KeyF3, ev)
		}
		
	}
}
