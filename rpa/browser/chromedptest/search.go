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
 * @Date:   2023/5/6 00:17
 * @Desc:
 */

// 打开网址搜索
func Search() {
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

	// 设置超时时间
	ctx, cancel = context.WithTimeout(ctx, time.Second*20)
	defer cancel()

	// 打开网页
	log.Println("打开网页")
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.baidu.com"),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 等待页面加载
	log.Println("等待页面加载")
	err = chromedp.Run(ctx,
		// wait for footer element is visible (ie, page is loaded)
		// chromedp.WaitVisible(`#kw`, chromedp.ByID),
		chromedp.WaitVisible(`/html/body/div[1]/div[1]/div[5]/div/div/form/span[1]/input`),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 输入关键字
	log.Println("输入关键字")
	err = chromedp.Run(ctx,
		chromedp.SendKeys(`#kw`, "golang", chromedp.ByID),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 点击搜索
	log.Println("点击搜索")
	err = chromedp.Run(ctx,
		chromedp.Click("#su", chromedp.ByID),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 等待搜索结果加载完成
	log.Println("等待搜索结果加载完成")
	err = chromedp.Run(ctx,
		// chromedp.Sleep(3*time.Second),
		chromedp.WaitVisible("#searchTag", chromedp.ByID),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 获取搜索结果
	log.Println("获取搜索结果")
	var result string
	err = chromedp.Run(ctx,
		// chromedp.Text("#searchTag", &result, chromedp.ByID),
		chromedp.InnerHTML("#searchTag", &result, chromedp.ByID),
		// chromedp.InnerHTML(".tag-container_ksKXH .tag-wrapper_1sGop", &result, chromedp.NodeVisible),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)

	// 保持打开的Chrome浏览器示例不主动退出
	log.Println("保持打开的Chrome浏览器示例不主动退出")
	err = chromedp.Run(ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			select {}
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
}
