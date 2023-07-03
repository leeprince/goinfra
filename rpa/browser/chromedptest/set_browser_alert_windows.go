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

// 显示浏览器执行自动化操作,并弹窗提示
func SetBrowserAlertWindows() {
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
		chromedp.Navigate("http://www.baidu.com"),
	)
	if err != nil {
		log.Fatal(err)
	}
	
	// 弹窗
	err = chromedp.Run(ctx,
		chromedp.Sleep(2*time.Second),
		chromedp.EvaluateAsDevTools(`alert("Hello, world!");`, nil),
	)
	if err != nil {
		fmt.Println(err)
	}
	
	// 弹窗
	err = chromedp.Run(ctx,
		chromedp.Sleep(2*time.Second),
		chromedp.EvaluateAsDevTools(`alert("Please close this dialog box manually.");`, nil),
	)
	if err != nil {
		fmt.Println(err)
	}
	
	// 弹窗
	err = chromedp.Run(ctx,
		chromedp.Sleep(2*time.Second),
		chromedp.EvaluateAsDevTools(`
            // 创建弹窗元素
            var popup = document.createElement('div');
            popup.innerHTML = '这是一个自动关闭的弹窗';
            
            // 设置弹窗的位置
            var right = "-300px"
            popup.style.position = 'fixed';
            popup.style.top = '60px';
            popup.style.right = right; // 将弹窗放置在窗口的右侧
            popup.style.padding = '10px';
            
            // 设置背景及边框
            popup.style.backgroundColor = '#fff';
            popup.style.border = '1px solid #ccc';
            popup.style.boxShadow = '0 2px 6px rgba(0, 0, 0, 0.3)';
            
            // 设置
            popup.style.width = '200px';
            popup.style.height = '20px';
            
            // 设置字体大小为18像素
            popup.style.fontSize = '18px';
            
            document.body.appendChild(popup);
            
            // 定义弹窗自动关闭的时间（单位：毫秒）
            var timeout = 5000;
            
            // 定义一个计时器，当时间到达时关闭弹窗
            var timer = setTimeout(function() {
              popup.style.right = right; // 将弹窗滑动回窗口的右侧
              setTimeout(function() {
                popup.style.display = 'none';
              }, 3000); // 等待0.5秒后隐藏弹窗
            }, timeout);
            
            // 将弹窗从右向左滑动到最终位置
            var slideIn = function() {
              var pos = -300;
            
              // 加快弹窗移动的速度：setInterval(frame, 5)中的5越小，速度越快。请注意，如果您将速度设置得太快，可能会导致弹窗的动画效果不太平滑因此，您需要根据自己的需要和喜好来调整速度。
              var id = setInterval(frame, 5);
              function frame() {
                if (pos == 20) {
                  clearInterval(id);
                } else {
                  pos++;
                  popup.style.right = pos + 'px';
                }
              }
            }
            slideIn();
            
            // 当鼠标移动到弹窗上时，清除计时器，防止弹窗关闭
            popup.addEventListener('mouseover', function() {
              clearTimeout(timer);
            });
            
            // 当鼠标离开弹窗时，重新开始计时
            popup.addEventListener('mouseout', function() {
              timer = setTimeout(function() {
                popup.style.right = '-300px'; // 将弹窗滑动回窗口的右侧
                setTimeout(function() {
                  popup.style.display = 'none';
                }, 5000); // 等待0.5秒后隐藏弹窗
              }, timeout);
            });
        `, nil),
	)
	if err != nil {
		fmt.Println(err)
	}
	
	time.Sleep(time.Second * 20)
}
