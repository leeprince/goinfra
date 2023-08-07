package main

import (
	"github.com/go-vgo/robotgo"
	"github.com/leeprince/goinfra/perror"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/4 00:30
 * @Desc:
 */

// 键盘操作
func KeyEnter() {
	time.Sleep(time.Second * 2)
	
	robotgo.KeySleep = 2000
	err = robotgo.KeyTap(robotgo.Enter)
	perror.ErrLogFatal(err)
}

// 键盘操作
func KeyEnterActiveName() {
	err = robotgo.ActiveName(ACTIVE_NAME)
	perror.ErrLogFatal(err)
	
	time.Sleep(time.Second * 2)
	
	robotgo.KeySleep = 2000
	err = robotgo.KeyTap(robotgo.Enter)
	perror.ErrLogFatal(err)
}

// 键盘操作
func KeyUpDown() {
	time.Sleep(time.Second * 2)
	
	robotgo.KeySleep = 1000
	err = robotgo.KeyTap(robotgo.Up)
	perror.ErrLogFatal(err)
	
	robotgo.KeySleep = 1000
	err = robotgo.KeyTap(robotgo.Down)
	perror.ErrLogFatal(err)
}

// 键盘操作
func KeyChar() {
	err = robotgo.KeyTap(robotgo.Key1)
	perror.ErrLogFatal(err)
	
	err = robotgo.KeyTap(robotgo.KeyA)
	perror.ErrLogFatal(err)
}

// 键盘操作
func KeyEnterActiveNameDarwinCtrlF() {
	err = robotgo.ActiveName(ACTIVE_NAME)
	perror.ErrLogFatal(err)
	
	time.Sleep(time.Second * 2)
	
	// --- mac: command+f
	/*
		报错：
		fatal error: unexpected signal during runtime execution
		[signal SIGSEGV: segmentation violation code=0x1 addr=0x8 pc=0x41aeda8]
	
		runtime stack:
		runtime.throw({0x4214fca?, 0x0?})
			/Users/leeprince/.g/go/src/runtime/panic.go:992 +0x71
		runtime.sigpanic()
			/Users/leeprince/.g/go/src/runtime/signal_unix.go:802 +0x3a9
	
		goroutine 35 [syscall]:
		runtime.cgocall(0x41aed80, 0xc000058cb8)
			/Users/leeprince/.g/go/src/runtime/cgocall.go:157 +0x5c fp=0xc000058c90 sp=0xc000058c58 pc=0x40073bc
		github.com/go-vgo/robotgo._Cfunc_keyCodeForChar(0x69)
			_cgo_gotypes.go:703 +0x47 fp=0xc000058cb8 sp=0xc000058c90 pc=0x41ab147
		github.com/go-vgo/robotgo.checkKeyCodes({0x4208a5b?, 0x2?})
			/Users/leeprince/go/pkg/mod/github.com/go-vgo/robotgo@v1.0.0-rc1/key.go:351 +0x7f fp=0xc000058d10 sp=0xc000058cb8 pc=0x41ab79f
		github.com/go-vgo/robotgo.keyTaps({0x4208a5b, 0x1}, {0xc00009a000?, 0x1?, 0x41ca4e0?}, 0x1?)
			/Users/leeprince/go/pkg/mod/github.com/go-vgo/robotgo@v1.0.0-rc1/key.go:409 +0x4d fp=0xc000058d48 sp=0xc000058d10 pc=0x41abd2d
		github.com/go-vgo/robotgo.KeyTap({0x4208a5b, 0x1}, {0xc000058f30?, 0x2, 0x2?})
			/Users/leeprince/go/pkg/mod/github.com/go-vgo/robotgo@v1.0.0-rc1/key.go:528 +0x579 fp=0xc000058ef8 sp=0xc000058d48 pc=0x41ac339
		github.com/leeprince/goinfra/rpa/clientoperate/macprogriumtest.Key()
			/Users/leeprince/www/go/goinfra/rpa/clientoperate/macprogriumtest/main.go:74 +0x205 fp=0xc000058f60 sp=0xc000058ef8 pc=0x41ae2a5
		github.com/leeprince/goinfra/rpa/clientoperate/macprogriumtest.TestKey(0x0?)
	*/
	err = robotgo.KeyTap(robotgo.KeyF, robotgo.Cmd)
	perror.ErrLogFatal(err)
	// 等于下面的操作
	// err = robotgo.KeyToggle(robotgo.Cmd, "down")
	// perror.ErrLogFatal(err)
	// err = robotgo.KeyTap(robotgo.KeyF)
	// perror.ErrLogFatal(err)
	// err = robotgo.KeyToggle(robotgo.Cmd, "up")
	// perror.ErrLogFatal(err)
	// ---
}

// 键盘操作
func KeyEnterActiveNameWindowsCtrlF() {
	err = robotgo.ActiveName(ACTIVE_NAME)
	perror.ErrLogFatal(err)
	
	time.Sleep(time.Second * 2)
	
	// --- windows: ctrl+f  >> windows 完美执行完成
	/*
		go version go1.18.5 windows/amd64
	*/
	/*
		gcc.exe (x86_64-posix-sjlj-rev0, Built by MinGW-W64 project) 8.1.0
			Copyright (C) 2018 Free Software Foundation, Inc.
			This is free software; see the resource for copying conditions.  There is NO
			warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
	*/
	err = robotgo.KeyTap(robotgo.KeyF, robotgo.Ctrl)
	perror.ErrLogFatal(err)
	// 等于下面的操作
	// err = robotgo.KeyToggle(robotgo.Ctrl, "down")
	// perror.ErrLogFatal(err)
	// err = robotgo.KeyTap(robotgo.KeyF)
	// perror.ErrLogFatal(err)
	// err = robotgo.KeyToggle(robotgo.Ctrl, "up")
	// perror.ErrLogFatal(err)
	// ---
	
}
