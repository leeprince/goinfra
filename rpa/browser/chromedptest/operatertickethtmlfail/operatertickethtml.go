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
	navigateRPAHtmlUrl = "http://127.0.0.1:19999/ticketHtmlWaitReadFile"
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
		orderId = "HT20230724175444A6YDXJ3W3408"
		
		// 是否是占座订单
		isOccupy = true
	)
	
	var selector string
	/*
		// 非占座票-出票成功
		document.querySelector("#BookSucTBody{订单ID} input.btn.btn-primary.btn-lg")
		// 非占座票-出票失败
		document.querySelector("#BookSucTBody{订单ID} input.btn.btn-default.btn-lg")
	
		// 占座票-占座成功
		document.querySelector("#BookSucTBody{订单ID} input.btn.btn-success.btn-lg")
		// 占座票-占座失败
		document.querySelector("#BookSucTBody{订单ID} input.btn.btn-danger.btn-lg")")
	*/
	if isOccupy {
		selector = fmt.Sprintf(`document.querySelector("#BookSucTBody%s input.btn.btn-danger.btn-lg")`, orderId)
		fmt.Println("占座票-点击出票失败 selector:", selector)
	} else {
		selector = fmt.Sprintf(`document.querySelector("#BookSucTBody%s input.btn.btn-default.btn-lg")`, orderId)
		fmt.Println("非占座票-点击出票失败 selector:", selector)
	}
	err = chromedp.Run(ctx,
		chromedp.Click(selector, chromedp.ByJSPath),
	)
	if err != nil {
		log.Fatal(err)
	}
	
	// 选择失败的原因
	// 需要稍微等待一下，否则弹窗可能还没出来；或者等待元素可见
	waitSeletorTime := time.Second * 5
	// 检查选择器
	selector = "#FailResonGroup"
	fmt.Println("选择失败的原因 selector:", selector)
	selctx, _ := context.WithTimeout(ctx, waitSeletorTime)
	err = chromedp.Run(selctx, chromedp.WaitVisible(selector, chromedp.ByID))
	if err != nil {
		fmt.Println("选择失败的原因 selector 不存在或者超时")
		log.Fatal("WaitVisible ", err)
	}
	fmt.Println("选择失败的原因 selector 已存在")
	err = chromedp.Run(ctx,
		chromedp.Sleep(time.Millisecond*300),
		
		// 按下 1，并且松开
		// 车次已无票
		chromedp.KeyEvent("1"), // 成功
		
		// 坐席无法满足
		/*chromedp.Sleep(time.Second*1),
		chromedp.KeyEvent("2"),
		// 坐席无法满足的具体单选框：无上铺
		chromedp.Sleep(time.Second*1),
		chromedp.Click("//*[@id='seatTypeNotMatch']/div/div/span[1]/input", chromedp.BySearch),
		// 坐席无法满足的具体单选框：无中铺
		chromedp.Sleep(time.Second*1),
		chromedp.Click("//*[@id='seatTypeNotMatch']/div/div/span[2]/input", chromedp.BySearch),
		// 坐席无法满足的具体单选框：无下铺
		chromedp.Sleep(time.Second*1),
		chromedp.Click("//*[@id='seatTypeNotMatch']/div/div/span[3]/input", chromedp.BySearch),
		// 坐席无法满足的具体单选框：无靠窗
		chromedp.Sleep(time.Second*1),
		chromedp.Click("//*[@id='seatTypeNotMatch']/div/div/span[7]/input", chromedp.BySearch),
		// 坐席无法满足的具体单选框：无过道
		chromedp.Sleep(time.Second*1),
		chromedp.Click("//*[@id='seatTypeNotMatch']/div/div/span[8]/input", chromedp.BySearch),
		// 坐席无法满足的具体单选框：无F
		chromedp.Sleep(time.Second*1),
		chromedp.Click("//*[@id='seatTypeNotMatch']/div/div/span[9]/input", chromedp.BySearch),
		// 坐席无法满足的具体单选框：无DF
		chromedp.Sleep(time.Second*1),
		chromedp.Click("//*[@id='seatTypeNotMatch']/div/div/span[10]/input", chromedp.BySearch),
		// 坐席无法满足的具体单选框：无连坐
		chromedp.Sleep(time.Second*1),
		chromedp.Click("//*[@id='seatTypeNotMatch']/div/div/span[11]/input", chromedp.BySearch),
		// 坐席无法满足的具体单选框：无同车厢
		chromedp.Sleep(time.Second*1),
		chromedp.Click("//*[@id='seatTypeNotMatch']/div/div/span[12]/input", chromedp.BySearch),
		// 坐席无法满足的具体单选框：无同包厢
		chromedp.Sleep(time.Second*1),
		chromedp.Click("//*[@id='seatTypeNotMatch']/div/div/span[13]/input", chromedp.BySearch),
		// 坐席无法满足的具体单选框：无指定车厢
		chromedp.Sleep(time.Second*1),
		chromedp.Click("//*[@id='seatTypeNotMatch']/div/div/span[14]/input", chromedp.BySearch),
		// 坐席无法满足的具体单选框：只有无座，用户不接受无座
		chromedp.Sleep(time.Second*1),
		chromedp.Click("//*[@id='seatTypeNotMatch']/div/div/span[17]/input", chromedp.BySearch),
		*/
		
		// 行程冲突：需选择具体乘客
		/*chromedp.Sleep(time.Second*1),
		chromedp.KeyEvent("3"),
		// 选择具体乘客：默认选择第一个即可
		chromedp.Sleep(time.Second*1),
		chromedp.Click("//*[@id='SubFailReasonForPassengerForm']/span/input", chromedp.BySearch),
		*/
		
		// 存在未支付订单：需选择具体乘客
		chromedp.Sleep(time.Second*1),
		chromedp.KeyEvent("4"),
		
		// 限制高消费：需选择具体乘客
		chromedp.Sleep(time.Second*1),
		chromedp.KeyEvent("5"),
		
		// 未采集联系方式：需选择具体乘客
		chromedp.Sleep(time.Second*1),
		chromedp.KeyEvent("6"),
		
		// 身份信息未核验：需选择具体乘客
		chromedp.Sleep(time.Second*1),
		chromedp.KeyEvent("7"),
		
		// 姓名不匹配：需选择具体乘客
		chromedp.Sleep(time.Second*1),
		chromedp.KeyEvent("8"),
		
		// 证件号重复
		chromedp.Sleep(time.Second*1),
		chromedp.KeyEvent("9"),
	
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
