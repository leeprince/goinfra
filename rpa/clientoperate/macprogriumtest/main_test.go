package main

import (
	"os"
	"runtime"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/4/20 23:13
 * @Desc:
 */

func TestMain(m *testing.M) {
	if runtime.GOOS == "windows" {
		ACTIVE_NAME = "sublime_text.exe"
	}

	os.Exit(m.Run())
}

// 选择窗口
func TestSelectWindow(t *testing.T) {
	SelectWindow()
}

// 输入字符
func TestInputChar(t *testing.T) {
	InputChar()
}

// 鼠标操作
func TestMouse(t *testing.T) {
	Mouse()
}

// 键盘操作
func TestKey(t *testing.T) {
	Key()
}
