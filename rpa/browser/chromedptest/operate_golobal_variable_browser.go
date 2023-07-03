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

var operateGolobalVariableBrowserContext context.Context

// 操作保存在全局变量中的浏览器，实际上是需要保存浏览器的上下文 context
func OperateGolobalVariableBrowser() {
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
	
	operateGolobalVariableBrowserContext = ctx
	
	// 打开目标网页
	err := chromedp.Run(ctx,
		chromedp.Navigate("http://www.baidu.com"),
	)
	if err != nil {
		log.Fatal(err)
	}
	
	time.Sleep(time.Second * 5)
	
	// 打开目标网页
	err = chromedp.Run(operateGolobalVariableBrowserContext,
		chromedp.Navigate("http://www.example.com/"),
	)
	if err != nil {
		log.Fatal(err)
	}
	
	time.Sleep(time.Second * 20)
}
