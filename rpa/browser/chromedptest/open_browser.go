package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/5/6 00:09
 * @Desc:
 */

// 显示浏览器执行自动化操作
func OpenBrowser() {
	// --- 创建有头浏览器 ---
	// 设置Chrome浏览器的启动参数
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("remote-debugging-port", "9222"),
	)
	allocCtx, allocCancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer allocCancel()
	
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
		chromedp.Navigate("http://www.baidu.com"),
	)
	if err != nil {
		log.Fatal(err)
	}
	
	time.Sleep(time.Second * 20)
}
