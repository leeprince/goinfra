package main

import (
	"context"
	"encoding/base64"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
	"log"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/4/22 15:09
 * @Desc:	监控指定URL的http请求：获取请求体内容时出错，成功获取的示例在`monitor_http_v1.go`中
 * 				外部HTTP服务：goinfra/http/httpservertest/sample/main.go
 */

func MonitorHttp() {
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

	// 监听指定的URL的HTTP请求
	url := "http://localhost:8090/prince/post"
	log.Println("监听指定的URL的HTTP请求 url:", url)
	listenCtx, cancelListen := context.WithCancel(ctx)
	defer cancelListen()
	chromedp.ListenTarget(listenCtx, func(ev interface{}) {
		switch ev := ev.(type) {
		case *network.EventRequestWillBeSent:
			log.Println("--- *network.EventRequestWillBeSent ---")
			go func() {
				if ev.Request.URL == url {
					log.Printf("Request URL: %s\nRequest Headers: %v\nRequest Body: %s\n", ev.Request.URL, ev.Request.Headers, ev.Request.PostData)
				}
			}()
		case *network.EventResponseReceived:
			log.Println("--- *network.EventResponseReceived ---")
			go func() {
				if ev.Response.URL == url {
					log.Printf("Response:%+v\n URL: %s\nResponse; Headers: %v\nResponse;", ev.Response, ev.Response.URL, ev.Response.Headers)

					// 获取响应体的实际内容
					/*
						在监听到的响应结果*network.EventResponseReceived中，ev.Response确实不存在Body属性。这是因为Body属性响应体的实际内容，而在Chrome DevTools协议中，响应体的实际内容是以base64编码的方式传输的，因此在*network.EventResponse中，ev.Response只包含了响应头和响应状态码等信息，而没有响应体的实际内容。

						如果你需要获取响应体的实际内容，可以使用chromedp提供的network.GetResponseBody`方法
					*/
					body, err := network.GetResponseBody(ev.RequestID).Do(ctx) // 总是会报错：invalid context
					if err != nil {
						log.Println("GetResponseBody", err)
						return
					}

					// 解码响应体的实际内容
					decodedBody, err := base64.StdEncoding.DecodeString(string(body))
					if err != nil {
						log.Println("DecodeString", err)
						return
					}
					log.Printf("Response decodedBody:%+v\n", decodedBody)
				}
			}()
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
