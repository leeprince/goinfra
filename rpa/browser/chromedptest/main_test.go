package main

import "testing"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/4/22 15:13
 * @Desc:
 */

func TestShowBroswer(t *testing.T) {
	ShowBrowser()
}

func TestCloseBrowserV1(t *testing.T) {
	go CloseBrowserV1()
	select {}
}
func TestCloseBrowserV2(t *testing.T) {
	go CloseBrowserV2()
	select {}
}
func TestCloseBrowserV3(t *testing.T) {
	go CloseBrowserV3()
	select {}
}

func TestSearch(t *testing.T) {
	Search()
}

func TestKeys(t *testing.T) {
	Keys()
}

func TestKeysV1(t *testing.T) {
	KeysV1()
}

func TestVisible(t *testing.T) {
	Visible()
}

func TestMonitorHttp(t *testing.T) {
	MonitorHttp()
}

func TestMonitorHttpV1(t *testing.T) {
	MonitorHttpV1()
}
