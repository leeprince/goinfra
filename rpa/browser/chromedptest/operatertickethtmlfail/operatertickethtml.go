package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/2 10:40
 * @Desc:
 */

const (
	// 要访问的 html页面地址
	navigateRPAHtmlUrl = "http://127.0.0.1:19999/ticket-html-2"
)

func main() {
	
	// --- 创建有头浏览器 ---
	// 设置Chrome浏览器的启动参数
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("remote-debugging-port", "9222"),
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	
	// 创建一个Chrome浏览器实例
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()
	// --- 创建有头浏览器-end ---
	
	// --- 创建无头浏览器：默认 ---
	// // create context
	// ctx, cancel := chromedp.NewContext(context.Background())
	// defer cancel()
	// --- 创建无头浏览器：默认-end ---
	
	// 打开目标网页
	err := chromedp.Run(ctx,
		chromedp.Navigate(navigateRPAHtmlUrl),
	)
	if err != nil {
		log.Fatal(err)
	}
	
	var (
		orderId = "HDTT202307091114133215077033---"
	)
	
	// 非占座票-点击出票失败
	selector := fmt.Sprintf(`document.querySelector("#BookSucTBody%s input.btn.btn-default.btn-lg")`, orderId)
	fmt.Println("非占座票-点击出票失败 selector:", selector)
	err = chromedp.Run(ctx,
		chromedp.Click(selector, chromedp.ByJSPath),
	)
	if err != nil {
		log.Fatal(err)
	}
	
	// 选择失败的原因
	// 需要稍微等待一下，否则弹窗可能还没出来；或者等待元素可见
	// time.Sleep(time.Second * 1)
	// waitSeletorTime := time.Second * 1
	waitSeletorTime := time.Millisecond * 200
	// 检查选择器
	selector = "FailResonGroup"
	fmt.Println("选择失败的原因 selector:", selector)
	selctx, _ := context.WithTimeout(ctx, waitSeletorTime)
	// 等待元素可见.chromedp.ByID 所以"SetBookFailPanel" 前面不带#
	err = chromedp.Run(selctx, chromedp.WaitVisible(selector, chromedp.ByID))
	if err != nil {
		fmt.Println("选择失败的原因 selector 不存在或者超时")
		log.Fatal("WaitVisible ", err)
	}
	fmt.Println("选择失败的原因 selector 已存在")
	err = chromedp.Run(ctx,
		chromedp.Sleep(time.Millisecond*300),
		
		// 按下 1，并且松开
		chromedp.KeyEvent("1"), // 成功
		// chromedp.KeyEvent(("\u0031")), // 成功
		
		chromedp.Sleep(time.Second*1),
		chromedp.KeyEvent("2"),
		
		chromedp.Sleep(time.Second*1),
		chromedp.KeyEvent("3"),
		
		chromedp.Sleep(time.Second*1),
		chromedp.KeyEvent("4"),
		
		chromedp.Sleep(time.Second*1),
		chromedp.KeyEvent("1"),
	
	)
	if err != nil {
		log.Fatal(err)
	}
	
	// 点击确定
	selector = fmt.Sprintf(`//*[@id="SetBookFailPanel"]/div[1]/div/div[3]/button[2]`)
	fmt.Println("设置占座失败-点击确定 selector:", selector)
	err = chromedp.Run(ctx,
		chromedp.Click(selector, chromedp.BySearch),
	)
	
	select {}
}
