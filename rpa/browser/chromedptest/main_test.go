package main

import "testing"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/4/22 15:13
 * @Desc:
 */

// 打开浏览器
func TestOpenBrowser(t *testing.T) {
	OpenBrowser()
}

// 关闭浏览器
func TestCloseBrowserV1(t *testing.T) {
	go CloseBrowserV1()
}
func TestCloseBrowserV2(t *testing.T) {
	CloseBrowserV2()
}
func TestCloseBrowserV3(t *testing.T) {
	CloseBrowserV3()
}

// 打开浏览器后，到百度中执行搜索
func TestSearch(t *testing.T) {
	Search()
}

// 按键输入
func TestKeys(t *testing.T) {
	Keys()
}

// 按键输入
func TestKeysV1(t *testing.T) {
	KeysV1()
}

// 在浏览器上设置窗口最大化
func TestSetMaxWindow(t *testing.T) {
	SetMaxWindow()
}

// 在浏览器上设置弹窗
func TestSetBrowserAlertWindow(t *testing.T) {
	SetBrowserAlertWindow()
}

// 等待指定选择器可见
func TestVisible(t *testing.T) {
	Visible()
}

// 监听http 网络
func TestMonitorHttp(t *testing.T) {
	MonitorHttp()
}

// 监听http 网络
func TestMonitorHttpV1(t *testing.T) {
	MonitorHttpV1()
}

// 操作保存在全局变量中的浏览器，实际上是需要保存浏览器的上下文 context
func TestOperateGolobalVariableBrowser(t *testing.T) {
	OperateGolobalVariableBrowser()
}
