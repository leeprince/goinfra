package main

import "testing"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/19 11:58
 * @Desc:
 */

func TestCheckOsKey(t *testing.T) {
	CheckOsKey()
	
	select {}
}

func TestWindowsOsListenKeyButton(t *testing.T) {
	WindowsOsListenKeyButton()
	
	select {}
}

func TestMacProOsListenKeyButton(t *testing.T) {
	MacProOsListenKeyButton()
	
	select {}
}
