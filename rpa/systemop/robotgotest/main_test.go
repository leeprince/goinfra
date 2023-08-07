package main

import (
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/4/20 23:13
 * @Desc:
 */

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
func TestKeyEnter(t *testing.T) {
	KeyEnter()
}

// 键盘操作
func TestKeyEnterActiveName(t *testing.T) {
	KeyEnterActiveName()
}

// 键盘操作
func TestKeyUpDown(t *testing.T) {
	KeyUpDown()
}

// 键盘操作
func TestKeyChar(t *testing.T) {
	KeyChar()
}

// 键盘操作
func TestKeyEnterActiveNameDarwinCtrlF(t *testing.T) {
	KeyEnterActiveNameDarwinCtrlF()
}

// 键盘操作
func TestKeyEnterActiveNameWindowsCtrlF(t *testing.T) {
	KeyEnterActiveNameWindowsCtrlF()
}

// 屏幕操作
func TestScreen(t *testing.T) {
	Screen()
}
