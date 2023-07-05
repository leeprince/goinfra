// Command keys is a chromedp example demonstrating how to send key events to
// an element.
package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp/kb"
	"log"
	"time"
	
	"github.com/chromedp/chromedp"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/5/6 00:35
 * @Desc:
 */

func KeysV1() {
	StartHttpServer()
	
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
	
	// run task list
	err := chromedp.Run(ctx,
		chromedp.Navigate(fmt.Sprintf("http://localhost:%d", *port)),
		chromedp.WaitVisible(`#input1`, chromedp.ByID),
		chromedp.WaitVisible(`#textarea1`, chromedp.ByID),
	)
	if err != nil {
		log.Fatal(err)
	}
	
	var val1 string
	time.Sleep(time.Second * 1)
	err = chromedp.Run(ctx,
		chromedp.SetValue(`#input1`, "#input1 value", chromedp.ByID),
		chromedp.Value(`#input1`, &val1, chromedp.ByID),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("#input1 value: %s", val1)
	
	var val2 string
	err = chromedp.Run(ctx,
		// chromedp.SendKeys(`#textarea1`, kb.End+"\b\b\n\naoeu\n\ntest1\n\nblah2\n\n\t\t\t\b\bother box!\t\ntest4", chromedp.ByID),
		chromedp.SendKeys(`#textarea1`, kb.End+"\nprince hello world!", chromedp.ByID),
		chromedp.Value(`#textarea1`, &val2, chromedp.ByID),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("#textarea1 value: %s", val2)
	
	var val3 string
	time.Sleep(time.Second * 1)
	err = chromedp.Run(ctx,
		chromedp.SetValue(`#input2`, "#input2 value", chromedp.ByID),
		chromedp.Value(`#input2`, &val3, chromedp.ByID),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("#textarea1 value: %s", val2)
	
	// var val4 string
	// time.Sleep(time.Second * 1)
	// err = chromedp.Run(ctx,
	// 	chromedp.SendKeys(`#select1`, kb.ArrowDown, chromedp.ByID),
	// 	chromedp.Value(`#select1`, &val4, chromedp.ByID),
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("#select1 value val4: %s", val4)
	
	var val5 string
	/*
		chromedp抛弃了Select方法，是因为该方法在某些情况下会出现一些问题，例如：
	
		在某些网页中，下拉列表的选项可能是动态生成的，而Select方法只能选择态的选项。
		在某些网页中，下拉列表的选项可能是隐藏的，而Select方法无法选择隐藏的选项。
		了解决这些问题，chromedp推荐使用Click和Keys方法来模拟用户的点击输入操作，以便更好地制定下拉列表的选择。
	*/
	/*
		最新版的chromedp已经抛弃了chromedp.Keys方法。这是因为chromedp.Keys方法在某些情况下会出现一问题，例如：
	
		在某些网页中，输入框可能是动态生成的，而`chromedp.Keys方法只能输入到已经存在的输入框中。
		在某些网页中，输入框可能是隐藏的，而chromedp.Keys方法无法输入到隐藏的输入框中。
		为了解决这些问题，chromedp推荐使用chromedp.SendKeys方法来模拟用户的输入操作，以便更好地处理各种复杂的情况。
	*/
	/*
		使用Click和SendKeys方法来模拟用户操作下拉列表，可以分为以下几个步骤：
	
		1. 点击下拉列表，打开下拉选项。[可以省略]
		2. 输入选项值，以便匹配选项。
	*/
	time.Sleep(time.Second * 1)
	// // 1.点击下拉列表，打开下拉项
	// err = chromedp.Run(ctx,
	// 	chromedp.Click("#select1", chromedp.ByID),
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// 2.输入选项值，以便匹选项
	// v等于select 中 option 展示的值，而不是 value
	err = chromedp.Run(ctx,
		chromedp.SendKeys("#select1", "2", chromedp.ByID), // v="2"而不是"two"
	)
	if err != nil {
		log.Fatal(err)
	}
	err = chromedp.Run(ctx,
		chromedp.Value(`#select1`, &val5, chromedp.ByID),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("#select1 value val5: %s", val5)
	
	time.Sleep(time.Second * 600)
}
