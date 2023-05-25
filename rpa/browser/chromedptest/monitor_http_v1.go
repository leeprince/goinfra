package main

import (
	"context"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
	"log"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/4/22 15:09
 * @Desc:	监控指定URL的http请求：成功获取响应内容
 * 				外部HTTP服务：goinfra/http/httpservertest/sample/main.go
 */

func MonitorHttpV1() {
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

	// 监听指定的URL的HTTP请求
	url := "http://localhost:8090/prince/post"
	log.Println("监听指定的URL的HTTP请求 url:", url)

	// 这将用于捕获匹配网络事件的请求id
	var requestID network.RequestID
	// 设置一个缓冲区为1的channel，用于异步接收等待监听指定url事件的结果通知
	listenChan := make(chan struct{}, 1)
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch ev := ev.(type) {
		case *network.EventRequestWillBeSent:
			log.Println("--- *network.EventRequestWillBeSent ---")
			if ev.Request.URL == url {
				log.Printf("Request URL: %s\nRequest Headers: %v\nRequest Body: %s\n", ev.Request.URL, ev.Request.Headers, ev.Request.PostData)
				requestID = ev.RequestID
			}
		case *network.EventResponseReceived:
			log.Println("--- *network.EventResponseReceived ---")
			if ev.RequestID == requestID {
				listenChan <- struct{}{}
			}
		}
	})

	// 打开目标网页
	log.Println("打开目标网页")
	err := chromedp.Run(ctx,
		chromedp.Navigate("http://localhost:8090"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// 获取多行输入框内容
	log.Println("获取多行输入框内容")
	var textareaValue string
	err = chromedp.Run(ctx,
		chromedp.SendKeys(`#textarea`, kb.End+"\nprince hello world!", chromedp.ByID),
		chromedp.Value(`#textarea`, &textareaValue, chromedp.ByID),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("#textarea1 value: %s", textareaValue)

	// 点击触发post Ajax请求
	log.Println("点击触发post Ajax请求")
	err = chromedp.Run(ctx,
		chromedp.Click("#sendPost", chromedp.ByID),
	)
	if err != nil {
		log.Fatal(err)
	}

	// 这将被阻止，直到chromedp监听器通知接收到响应的通知到channel中
	for {
		select {
		case <-listenChan:
			// 根据已接收到响应的请求ID获取响应的内容
			if err := chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
				byteBody, err := network.GetResponseBody(requestID).Do(ctx)
				if err != nil {
					log.Println("GetResponseBody", err)
					return err
				}

				// 响应内容
				log.Printf("Response body:%+v\n", string(byteBody))

				return nil
			})); err != nil {
				log.Fatal(err)
			}
		}
	}
}
