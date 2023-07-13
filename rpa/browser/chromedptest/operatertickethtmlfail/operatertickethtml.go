package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/leeprince/goinfra/utils/fileutil"
	"log"
	"net/http"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/2 10:40
 * @Desc:
 */

const (
	// 启动 http 服务器后的要访问的 html页面地址
	navigateRPAHtmlUrl = "http://localhost:8090/defaultHandler"
	htmlFileDir        = "/Users/leeprince/www/go/goinfra/rpa/browser/chromedptest/operatertickethtmlfail"
)

var port *int

// 启动 http 服务器
func HttpServer() {
	port = flag.Int("port", 8090, "port")
	flag.Parse()
	
	http.HandleFunc("/defaultHandler", defaultHandler)
	
	fmt.Printf("Server listening on port:%d ...\n", *port)
	go http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		
		fileBytes, err := fileutil.ReadFile(htmlFileDir, "ticket.html")
		if err != nil {
			http.Error(w, "读取 html文件错误", http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, string(fileBytes))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	HttpServer()
	
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
		orderId = "FeiZ3424193749701983461"
	)
	
	// 非占座票-出票失败
	selector := fmt.Sprintf(`document.querySelector("#BookSucTBody%s input.btn.btn-default.btn-lg")`, orderId)
	fmt.Println("sel:", selector)
	err = chromedp.Run(ctx,
		chromedp.Click(selector, chromedp.ByJSPath),
	)
	
	select {}
}
