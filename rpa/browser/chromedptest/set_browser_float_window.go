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
 * @Date:   2023/5/6 00:09
 * @Desc:
 */

// 显示浏览器执行自动化操作,悬浮弹窗
func SetBrowserFloatWindow() {
	// --- 创建有头浏览器 ---
	// 设置Chrome浏览器的启动参数
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("remote-debugging-port", "9222"),
		// 设置最大窗口
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
	
	// 打开目标网页
	err := chromedp.Run(ctx,
		chromedp.Navigate("http://www.baidu.com"),
	)
	if err != nil {
		log.Fatal(err)
	}
	
	// 弹窗
	/*
		注意：如果需要通过 innerHTML 添加html内容则需要使用`或者""包含起来，但是由于 golang 不能嵌套`符，所以
		var popup = document.createElement("table");
		popup.innerHTML = %s%s%s;
		popup.appendChild(tableEl);
	
		说明：
			第一个%s：`
			第二个%s：html 内容
			第三个%s：`
	*/
	err = chromedp.Run(ctx,
		chromedp.EvaluateAsDevTools(`
            // 创建弹窗元素
			var popup = document.createElement('div');
			popup.id = "popupId";
			popup.innerHTML = '这是一个自动关闭的弹窗';
			
			// 设置弹窗的位置
			var right = "0px"
			popup.style.position = 'fixed';

			// 弹窗出现在右上角
			popup.style.top = '60px';
			popup.style.right = right;

			// 弹窗出现在右下角
			// popup.style.bottom = '20px';
			// popup.style.right = '20px';

			// 弹窗出现在左下角
			// popup.style.bottom = '50px';
			// popup.style.left = '20px';
			
			// 内填充
			popup.style.padding = '10px';

			// 设置背景及边框
			popup.style.backgroundColor = '#fff';
			popup.style.border = '1px solid #ccc';
			popup.style.boxShadow = '0 2px 6px rgba(0, 0, 0, 0.3)';
			
			// 设置边框大小
			popup.style.width = '200px';
			popup.style.height = '20px';
			
			// 设置字体大小为18像素
			popup.style.fontSize = '18px';
			
			document.body.appendChild(popup);

			// 定义弹窗自动关闭的时间（单位：毫秒）
			var timeout = 5000// 定义一个计时器，当时间到达时关闭弹窗
			var timer = setTimeout(function() {
			  popup.style.display = 'none';
			}, timeout);
			
			// 当鼠标移动到弹窗上时，清除计时器，防止弹窗关闭
			popup.addEventListener('mouseover', function() {
			  clearTimeout(timer);
			});
			
			// 当鼠标离开弹窗时，重新开始计时
			popup.addEventListener('mouseout', function() {
			  timer = setTimeout(function() {
				popup.style.display = 'none';
			  }, timeout);
			});
        `, nil),
	)
	if err != nil {
		fmt.Println(err)
	}
	
	time.Sleep(time.Second * 5)
	
	// 移除弹窗
	expresion := fmt.Sprintf(`document.getElementById("%s").remove()`, "popupId")
	err = chromedp.Run(ctx,
		chromedp.EvaluateAsDevTools(expresion, nil),
	)
	if err != nil {
		log.Fatal("")
		return
	}
	
	time.Sleep(time.Second * 2000)
}
