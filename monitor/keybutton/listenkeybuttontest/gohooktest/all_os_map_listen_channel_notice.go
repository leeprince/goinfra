package main

import (
	"fmt"
	"github.com/leeprince/goinfra/consts"
	"github.com/leeprince/goinfra/consts/constval"
	"github.com/leeprince/goinfra/utils/sliceutil"
	hook "github.com/robotn/gohook"
	"log"
	"runtime"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/20 22:28
 * @Desc:
 */

var (
	F1KeyUpCh chan struct{}
	F2KeyUpCh chan struct{}
	F3KeyUpCh chan struct{}
)

func ListenWindowAndMacproChannelNotice() {
	hooks := hook.Start()
	defer hook.End()
	
	var keyMap *constval.StringUint16Group
	if !sliceutil.InString(runtime.GOOS, []string{
		consts.GOOSDarwin,
		consts.GOOSWindows,
	}) {
		log.Fatal("暂不支持该操作系统")
	}
	keyMap = consts.WindowsOSKeyButtonRawcode
	if runtime.GOOS == consts.GOOSDarwin {
		keyMap = consts.DarwinOSKeyButtonRawcode
	}
	
	for ev := range hooks {
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
