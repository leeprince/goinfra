package main

import (
	"fmt"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/19 11:58
 * @Desc:
 */

func TestCheckOsKey(t *testing.T) {
	CheckOsKey()
	
	select {}
}

func TestWindowsOsListen(t *testing.T) {
	WindowsOsListen()
	
	select {}
}

func TestMacProOsListen(t *testing.T) {
	MacProOsListen()
	
	select {}
}

func TestListenWindowAndMacpro(t *testing.T) {
	ListenWindowAndMacpro()
	
	select {}
}

func TestListenWindowAndMacproChannelNotice(t *testing.T) {
	EscKeyUpCh = make(chan struct{}, 0)
	F1KeyUpCh = make(chan struct{}, 0)
	F2KeyUpCh = make(chan struct{}, 0)
	F3KeyUpCh = make(chan struct{}, 0)
	
	go ListenWindowAndMacproChannelNotice()
	
	go func() {
		for {
			select {
			case <-EscKeyUpCh:
				fmt.Println("接收到按键 Esc被按下的通知")
			case <-F1KeyUpCh:
				fmt.Println("接收到按键 F1被按下的通知")
			case <-F2KeyUpCh:
				fmt.Println("接收到按键 F2被按下的通知")
			case <-F3KeyUpCh:
				fmt.Println("接收到按键 F3被按下的通知")
			}
		}
	}()
	
	/*// 多个协程监听。只会随机进入一个协程的通道
	go func() {
		for {
			select {
			case <-EscKeyUpCh:
				fmt.Println("1 接收到按键 Esc被按下的通知")
			case <-F1KeyUpCh:
				fmt.Println("1 接收到按键 F1被按下的通知")
			case <-F2KeyUpCh:
				fmt.Println("1 接收到按键 F2被按下的通知")
			case <-F3KeyUpCh:
				fmt.Println("1 接收到按键 F3被按下的通知")
			}
		}
	
	}()*/
	
	select {}
}
