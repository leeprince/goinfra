// Command keys is a chromedp example demonstrating how to send key events to
// an element.
package main

import (
	"context"
	"fmt"
	"github.com/go-vgo/robotgo"
	"log"
	"time"
	
	"github.com/chromedp/chromedp"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/5/6 00:35
 * @Desc:
 */

func KeysV2() {
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
	// create context
	// ctx, cancel := chromedp.NewContext(context.Background())
	// defer cancel()
	// --- 创建无头浏览器：默认-end ---
	
	// 执行任务
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.example.com"),
	)
	if err != nil {
		log.Fatal(err)
	}
	
	/*// 弹窗
	err = chromedp.Run(ctx,
		chromedp.Sleep(time.Second*1),
		chromedp.EvaluateAsDevTools(`alert("Hello, world!");`, nil),
	)
	if err != nil {
		fmt.Println(err)
	}
	
	// 按回车键关闭弹窗:因为上面的alert弹窗会堵塞程序继续运行，所以无法在弹窗出来时关闭弹窗的
	err = chromedp.Run(ctx,
		chromedp.Sleep(time.Second*3),
		chromedp.SendKeys("", kb.Enter), // 模拟按下回车键
	)
	if err != nil {
		log.Fatal(err)
	}*/
	
	// 同样无法关闭 alert 弹窗。具体原因：
	/*
		由于浏览器的安全限制，无法直接通过代码关闭 alert 弹窗。alert 弹窗是阻塞式的，它会阻止 JavaScript 的执行，直到用户关闭弹窗为止。这是为了确保用户能够看到并响应弹窗中的消息。
		如果您需要更高度的控制和自定义弹窗行为，您可以考虑使用模态框或自定义弹窗来替代 alert 弹窗。这样您就可以通过代码来控制弹窗的显示和关闭，而不会受到阻塞的限制。
	*/
	/*
		最终解决办法：请看下面
	*/
	/*go func() {
		fmt.Println("time.Sleep ...")
		time.Sleep(time.Second * 5)
		fmt.Println("time.Sleep >>>")
		err = chromedp.Run(ctx,
			chromedp.SendKeys("", kb.Enter), // 模拟按下回车键
		)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(">kb.Enter")
	
		// 获取当前页面的 DevTools Session ID
		sessionID := chromedp.FromContext(ctx).Target.SessionID
		fmt.Println(">FromContext SessionID", sessionID)
	
		// 执行 JavaScript 代码关闭 alert 弹窗
		err = chromedp.Run(ctx,
			chromedp.EvaluateAsDevTools("window.alert = function() { return true; }", sessionID),
		)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(">window.alert")
	}()*/
	
	/*
		最终解决办法：通过系统按键操作 <<< 完美
	*/
	go func() {
		time.Sleep(time.Second * 3)
		err = robotgo.KeyTap(robotgo.Enter)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(">robotgo.KeyTap(robotgo.Enter) return true")
	}()
	
	// 弹窗
	err = chromedp.Run(ctx,
		chromedp.Sleep(time.Second*1),
		chromedp.EvaluateAsDevTools(`alert("Hello, world!");`, nil),
	)
	if err != nil {
		fmt.Println(err)
	}
}
