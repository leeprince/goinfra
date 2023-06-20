package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/6/20 17:04
 * @Desc:
 */

func SetMaxWindow() {
	var err error

	// --- 创建有头浏览器 ---
	// 设置Chrome浏览器的启动参数
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("remote-debugging-port", "9222"),
		chromedp.WindowSize(1920, 1040),
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

	// 获取当前屏幕大小
	var width, height int
	err = chromedp.Run(ctx, chromedp.EvaluateAsDevTools(`window.screen.availWidth`, &width),
		chromedp.EvaluateAsDevTools(`window.screen.availHeight`, &height))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("width:", width)
	fmt.Println("height:", height)

	// 打开目标网页
	err = chromedp.Run(ctx,
		chromedp.Navigate("http://www.baidu.com"),
	)
	if err != nil {
		log.Fatal(err)
	}

	select {}
}
