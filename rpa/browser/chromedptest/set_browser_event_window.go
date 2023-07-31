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
 * @Date:   2023/7/23 01:06
 * @Desc:	在网页上设置弹窗，并在弹窗中支持按钮
 */

/*
需要通过"github.com/chromedp/chromedp"打开www.baidu.com网页后，左下角添加一个悬浮框，悬浮框中包含五部分内容。
    标题：按键说明
    全局自动化模式：自动模式、分步骤模式
    f1:根据订单选票
    f2:输入乘客信息
    f3:开始扣款

同时，需要通过 golang 控制悬浮框中第 2 点的内容显示：是自动模式还是分步骤模式
*/
func SetBrowserEventWindow() {
	
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
	
	autoModeElID := "PrinceSetAutoMode"
	
	tableContent := `
		<thead>
			<tr>
			  <th>按键</th>
			  <th>仅在"半自动分段"模式有效</th>
			</tr>
		</thead>
		<tbody>
		  <tr>
			<td>f1</td>
			<td>根据订单选票</td>
		  </tr>
		  <tr>
			<td>f2</td>
			<td>输入乘客信息</td>
		  </tr>
		  <tr>
			<td>f3</td>
			<td>开始扣款</td>
		  </tr>
		</tbody>
	`
	err = chromedp.Run(ctx,
		chromedp.EvaluateAsDevTools(fmt.Sprintf(`
			var floatBox = document.createElement("div");
			floatBox.style.position = "fixed";
			floatBox.style.bottom = "0";
			floatBox.style.left = "0";
			floatBox.style.width = "200px";
			floatBox.style.height = "150px";
			floatBox.style.background = "#fff";
			floatBox.style.border = "1px solid #ccc";
			floatBox.style.padding = "10px";
			floatBox.style.zIndex = "9999";
			document.body.appendChild(floatBox);
			
			var title = document.createElement("h3");
			title.innerText = "辅助窗口";
			floatBox.appendChild(title);

			var autoModeEl = document.createElement("h5");
			autoModeEl.id = "%s"
			autoModeEl.innerText = "当前自动化模式：全自动";
			autoModeEl.onclick = function() {
			  if (autoModeEl.innerText == "当前自动化模式：全自动") {
			    autoModeEl.innerText = "当前自动化模式：半自动分段";
			  } else {
			    autoModeEl.innerText = "当前自动化模式：全自动";
			  }
			};
			floatBox.appendChild(autoModeEl);

			var tipsEl = document.createElement("span");
			tipsEl.innerText = "1.自动化模式(按Esc键切换)\n2.自动化模式:全自动、半自动分段";
			floatBox.appendChild(tipsEl);
			
			var lineEl = document.createElement("p");
			lineEl.innerHTML = "<hr />";
			floatBox.appendChild(lineEl);

			var tableEl = document.createElement("table");
			tableEl.innerHTML = %s%s%s;
			floatBox.appendChild(tableEl);
		`, autoModeElID, "`", tableContent, "`"), nil),
	)
	if err != nil {
		log.Println(">>>>>>windows err:", err)
	}
	
	// 等待后，触发点击切换
	time.Sleep(time.Second * 3)
	autoModelSelector := fmt.Sprintf("#%s", autoModeElID)
	err = chromedp.Run(ctx, chromedp.Click(autoModelSelector, chromedp.ByID))
	if err != nil {
		log.Fatal(">>>>>>Click err:", err)
	}
	
	// 等待后，检查该弹窗是否存在，用于判断弹窗不存在时重新设置浮动弹窗
	fmt.Println("等待后，检查该弹窗是否存在，用于判断弹窗不存在时重新设置浮动弹窗")
	time.Sleep(time.Second * 5)
	
	// 执行 JavaScript 代码来检查指定 ID 是否存在
	var exists bool
	err = chromedp.Run(ctx, chromedp.Evaluate(fmt.Sprintf(`document.getElementById("%s") !== null`, autoModeElID), &exists))
	if err != nil {
		panic(err)
	}
	if exists {
		fmt.Println("ID exists")
	} else {
		fmt.Println("ID does not exist")
	}
	
	// 等待后，关闭浏览器
	time.Sleep(time.Second * 120)
	err = chromedp.Cancel(ctx)
	if err != nil {
		log.Fatal(">>>>>>Cancel err:", err)
	}
}
