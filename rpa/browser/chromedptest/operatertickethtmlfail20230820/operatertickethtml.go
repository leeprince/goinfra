package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/leeprince/goinfra/perror"
	"github.com/leeprince/goinfra/plog"
	"github.com/leeprince/goinfra/test/constants"
	"github.com/leeprince/goinfra/utils/idutil"
	"log"
	"strings"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/2 10:40
 * @Desc:
 */

const (
	// 要访问的 html页面地址
	navigateRPAHtmlUrl = "http://127.0.0.1:19999/ticketHtmlReadFile20230820"
)

var (
	waitKeyEventTime = time.Millisecond * 400
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
		orderId = "HT20230819184809OH63YOGY1426"
		// 是否是占座订单
		IsOccupySeat = true
		
		waitSeletorSecond = time.Second * 8
	)
	
	logID := idutil.UniqIDV3()
	plogEntry := plog.LogID(logID).
		WithField("orderId", orderId).
		WithField("method", "RPAOrderTaskService.HandlerCallbackResultNoTicket")
	plogEntry.Info("request")
	
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
	if IsOccupySeat {
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
	
	chromeCtx := ctx
	// 检查选择器
	selctx, _ := context.WithTimeout(chromeCtx, waitSeletorSecond)
	err = chromedp.Run(selctx, chromedp.WaitVisible(selector, chromedp.ByJSPath))
	if err != nil {
		plogEntry.WithError(err).Error("非占座票-点击出票失败-选择器不存在")
		log.Fatal(constants.CallBackErrSelect.Key(), constants.CallBackErrSelect.Value())
	}
	err = chromedp.Run(chromeCtx,
		chromedp.Click(selector, chromedp.ByJSPath),
	)
	if err != nil {
		plogEntry.WithError(err).Error("点击出票失败 失败")
		log.Fatal(constants.CallBackErrSelect.Key(), constants.CallBackErrSelect.Value())
	}
	
	// 检查选择器
	selector = "#FailResonGroup"
	fmt.Println("选择失败的原因 selector:", selector)
	selctx, _ = context.WithTimeout(chromeCtx, waitSeletorSecond)
	err = chromedp.Run(selctx, chromedp.WaitVisible(selector, chromedp.ByID))
	if err != nil {
		plogEntry.WithError(err).Error("选择失败的原因 selector 不存在或者超时")
		log.Fatal(constants.CallBackErrSelect.Key(), constants.CallBackErrSelect.Value())
	}
	plogEntry.Info("选择失败的原因 selector 已存在")
	
	var (
		reason                 = ""
		seatPositionTypeString = ""
		transferNo             = int32(1)
	)
	// 匹配原因并设置按键
	if strings.Contains(reason, "已售完") ||
		strings.Contains(reason, "已无票") ||
		strings.Contains(reason, "要否？") ||
		strings.Contains(reason, "异常") ||
		strings.Contains(reason, "满足") {
		// 坐席无法满足
		if strings.Contains(seatPositionTypeString, "上铺") {
			// 坐席无法满足的具体单选框：无上铺
			err = chromedp.Run(chromeCtx,
				chromedp.Sleep(waitKeyEventTime),
				chromedp.KeyEvent("2"),
				chromedp.Sleep(waitKeyEventTime),
				chromedp.KeyEvent("2"),
				chromedp.Sleep(waitKeyEventTime),
				chromedp.Click("//*[@id='seatTypeNotMatch']/div/div/span[1]/label/input", chromedp.BySearch),
			)
			if err != nil {
				plogEntry.WithError(err).Error("坐席无法满足的具体单选框：无上铺 err")
				err = perror.BizErrOpreate.SetMessage("操作失败：坐席无法满足的具体单选框：无上铺")
				return
			}
			
		} else if strings.Contains(seatPositionTypeString, "中铺") {
			// 坐席无法满足的具体单选框：无中铺
			err = chromedp.Run(chromeCtx,
				chromedp.Sleep(waitKeyEventTime),
				chromedp.KeyEvent("2"),
				chromedp.Sleep(waitKeyEventTime),
				chromedp.KeyEvent("2"),
				chromedp.Sleep(waitKeyEventTime),
				chromedp.Click("//*[@id='seatTypeNotMatch']/div/div/span[2]/label/input", chromedp.BySearch),
			)
			if err != nil {
				plogEntry.WithError(err).Error("坐席无法满足的具体单选框：无中铺 err")
				err = perror.BizErrOpreate.SetMessage("操作失败：坐席无法满足的具体单选框：无中铺")
				return
			}
			
		} else if strings.Contains(seatPositionTypeString, "下铺") {
			// 坐席无法满足的具体单选框：无下铺
			err = chromedp.Run(chromeCtx,
				chromedp.Sleep(waitKeyEventTime),
				chromedp.KeyEvent("2"),
				chromedp.Sleep(waitKeyEventTime),
				chromedp.KeyEvent("2"),
				chromedp.Sleep(waitKeyEventTime),
				chromedp.Click("//*[@id='seatTypeNotMatch']/div/div/span[3]/label/input", chromedp.BySearch),
			)
			if err != nil {
				plogEntry.WithError(err).Error("坐席无法满足的具体单选框：无下铺 err")
				err = perror.BizErrOpreate.SetMessage("操作失败：坐席无法满足的具体单选框：无下铺")
				return
			}
			
		} else if strings.Contains(seatPositionTypeString, "窗") {
			// 坐席无法满足的具体单选框：无靠窗
			err = chromedp.Run(chromeCtx,
				chromedp.Sleep(waitKeyEventTime),
				chromedp.KeyEvent("2"),
				chromedp.Sleep(waitKeyEventTime),
				chromedp.KeyEvent("2"),
				chromedp.Sleep(waitKeyEventTime),
				chromedp.Click("//*[@id='seatTypeNotMatch']/div/div/span[7]/label/input", chromedp.BySearch),
			)
			if err != nil {
				plogEntry.WithError(err).Error("坐席无法满足的具体单选框：无靠窗 err")
				err = perror.BizErrOpreate.SetMessage("操作失败：坐席无法满足的具体单选框：无靠窗")
				return
			}
			
		} else if strings.Contains(seatPositionTypeString, "过道") {
			// 坐席无法满足的具体单选框：无过道
			err = chromedp.Run(chromeCtx,
				chromedp.Sleep(waitKeyEventTime),
				chromedp.KeyEvent("2"),
				chromedp.Sleep(waitKeyEventTime),
				chromedp.KeyEvent("2"),
				chromedp.Sleep(waitKeyEventTime),
				chromedp.Click("//*[@id='seatTypeNotMatch']/div/div/span[8]/label/input", chromedp.BySearch),
			)
			if err != nil {
				plogEntry.WithError(err).Error("坐席无法满足的具体单选框：无过道 err")
				err = perror.BizErrOpreate.SetMessage("操作失败：坐席无法满足的具体单选框：无过道")
				return
			}
			
		} else if strings.Contains(seatPositionTypeString, "DF") {
			// 坐席无法满足的具体单选框：无DF
			err = chromedp.Run(chromeCtx,
				chromedp.Sleep(waitKeyEventTime),
				chromedp.KeyEvent("2"),
				chromedp.Sleep(waitKeyEventTime),
				chromedp.KeyEvent("2"),
				chromedp.Sleep(waitKeyEventTime),
				chromedp.Click("//*[@id='seatTypeNotMatch']/div/div/span[10]/label/input", chromedp.BySearch),
			)
			if err != nil {
				plogEntry.WithError(err).Error("坐席无法满足的具体单选框：无DF err")
				err = perror.BizErrOpreate.SetMessage("操作失败：坐席无法满足的具体单选框：无DF")
				return
			}
			
		} else if strings.Contains(seatPositionTypeString, "F") {
			// 坐席无法满足的具体单选框：无F
			err = chromedp.Run(chromeCtx,
				chromedp.Sleep(waitKeyEventTime),
				chromedp.KeyEvent("2"),
				chromedp.Sleep(waitKeyEventTime),
				chromedp.KeyEvent("2"),
				chromedp.Sleep(waitKeyEventTime),
				chromedp.Click("//*[@id='seatTypeNotMatch']/div/div/span[9]/label/input", chromedp.BySearch),
			)
			if err != nil {
				plogEntry.WithError(err).Error("坐席无法满足的具体单选框：无F err")
				err = perror.BizErrOpreate.SetMessage("操作失败：坐席无法满足的具体单选框：无F")
				return
			}
		} else if strings.Contains(seatPositionTypeString, "连坐") {
			// 坐席无法满足的具体单选框：无连坐
			err = chromedp.Run(chromeCtx,
				chromedp.Sleep(waitKeyEventTime),
				chromedp.KeyEvent("2"),
				chromedp.Sleep(waitKeyEventTime),
				chromedp.KeyEvent("2"),
				chromedp.Sleep(waitKeyEventTime),
				chromedp.Click("//*[@id='seatTypeNotMatch']/div/div/span[11]/label/input", chromedp.BySearch),
			)
			if err != nil {
				plogEntry.WithError(err).Error("坐席无法满足的具体单选框：无连坐 err")
				err = perror.BizErrOpreate.SetMessage("操作失败：坐席无法满足的具体单选框：无连坐")
				return
			}
			
		} else if strings.Contains(seatPositionTypeString, "同车厢") {
			// 坐席无法满足的具体单选框：无同车厢
			err = chromedp.Run(chromeCtx,
				chromedp.Sleep(waitKeyEventTime),
				chromedp.KeyEvent("2"),
				chromedp.Sleep(waitKeyEventTime),
				chromedp.KeyEvent("2"),
				chromedp.Sleep(waitKeyEventTime),
				chromedp.Click("//*[@id='seatTypeNotMatch']/div/div/span[12]/label/input", chromedp.BySearch),
			)
			if err != nil {
				plogEntry.WithError(err).Error("坐席无法满足的具体单选框：无同车厢 err")
				err = perror.BizErrOpreate.SetMessage("操作失败：坐席无法满足的具体单选框：无同车厢")
				return
			}
			
		} else if strings.Contains(seatPositionTypeString, "同包厢") {
			// 坐席无法满足的具体单选框：无同包厢
			err = chromedp.Run(chromeCtx,
				chromedp.Sleep(waitKeyEventTime),
				chromedp.KeyEvent("2"),
				chromedp.Sleep(waitKeyEventTime),
				chromedp.KeyEvent("2"),
				chromedp.Sleep(waitKeyEventTime),
				chromedp.Click("//*[@id='seatTypeNotMatch']/div/div/span[13]/label/input", chromedp.BySearch),
			)
			if err != nil {
				plogEntry.WithError(err).Error("坐席无法满足的具体单选框：无同包厢 err")
				err = perror.BizErrOpreate.SetMessage("操作失败：坐席无法满足的具体单选框：无同包厢")
				return
			}
			
		} else if strings.Contains(seatPositionTypeString, "不接受无座") {
			// 坐席无法满足的具体单选框：只有无座，用户不接受无座
			err = chromedp.Run(chromeCtx,
				chromedp.Sleep(waitKeyEventTime),
				chromedp.KeyEvent("2"),
				chromedp.Sleep(waitKeyEventTime),
				chromedp.KeyEvent("2"),
				chromedp.Sleep(waitKeyEventTime),
				chromedp.Click("//*[@id='seatTypeNotMatch']/div/div/span[17]/label/input", chromedp.BySearch),
			)
			if err != nil {
				plogEntry.WithError(err).Error("坐席无法满足的具体单选框：只有无座，用户不接受无座 err")
				err = perror.BizErrOpreate.SetMessage("操作失败：坐席无法满足的具体单选框：只有无座，用户不接受无座")
				return
			}
			
		} else {
			err = HandlerCallbackResultNoTicketOfNoTicket(chromeCtx, transferNo)
			if err != nil {
				plogEntry.WithError(err).Error("车次已无票 err")
				err = perror.BizErrOpreate.SetMessage("操作失败：车次已无票")
				return
			}
			
		}
	} else if strings.Contains(reason, "冲突") {
		// 行程冲突
		err = chromedp.Run(chromeCtx,
			chromedp.Sleep(waitKeyEventTime),
			chromedp.KeyEvent("3"),
			// 选择具体乘客：默认选择第一个即可
			chromedp.Sleep(waitKeyEventTime),
			chromedp.Click("//*[@id='SubFailReasonForPassengerForm']/span/input", chromedp.BySearch),
		
		)
		if err != nil {
			plogEntry.WithError(err).Error("行程冲突 err")
			err = perror.BizErrOpreate.SetMessage("操作失败：行程冲突")
			return
		}
	} else if strings.Contains(reason, "未支付") {
		// 存在未支付订单：需选择具体乘客
		err = chromedp.Run(chromeCtx,
			chromedp.Sleep(waitKeyEventTime),
			chromedp.KeyEvent("4"),
			// 选择具体乘客：默认选择第一个即可
			chromedp.Sleep(waitKeyEventTime),
			chromedp.Click("//*[@id='SubFailReasonForPassengerForm']/span/input", chromedp.BySearch),
		)
		if err != nil {
			plogEntry.WithError(err).Error("存在未支付订单 err")
			err = perror.BizErrOpreate.SetMessage("操作失败：存在未支付订单")
			return
		}
	} else if strings.Contains(reason, "消费") {
		// 限制高消费：需选择具体乘客
		err = chromedp.Run(chromeCtx,
			chromedp.Sleep(waitKeyEventTime),
			chromedp.KeyEvent("5"),
			// 选择具体乘客：默认选择第一个即可
			chromedp.Sleep(waitKeyEventTime),
			chromedp.Click("//*[@id='SubFailReasonForPassengerForm']/span/input", chromedp.BySearch),
		
		)
		if err != nil {
			plogEntry.WithError(err).Error("限制高消费 err")
			err = perror.BizErrOpreate.SetMessage("操作失败：限制高消费")
			return
		}
	} else if strings.Contains(reason, "采集") ||
		strings.Contains(reason, "联系") {
		// 未采集联系方式：需选择具体乘客
		err = chromedp.Run(chromeCtx,
			chromedp.Sleep(waitKeyEventTime),
			chromedp.KeyEvent("6"),
			// 选择具体乘客：默认选择第一个即可
			chromedp.Sleep(waitKeyEventTime),
			chromedp.Click("//*[@id='SubFailReasonForPassengerForm']/span/input", chromedp.BySearch),
		)
		if err != nil {
			plogEntry.WithError(err).Error("未采集联系方式 err")
			err = perror.BizErrOpreate.SetMessage("操作失败：未采集联系方式")
			return
		}
	} else if strings.Contains(reason, "核验") {
		// 身份信息未核验：需选择具体乘客
		err = chromedp.Run(chromeCtx,
			chromedp.Sleep(waitKeyEventTime),
			chromedp.KeyEvent("7"),
			// 选择具体乘客：默认选择第一个即可
			chromedp.Sleep(waitKeyEventTime),
			chromedp.Click("//*[@id='SubFailReasonForPassengerForm']/span/input", chromedp.BySearch),
		)
		if err != nil {
			plogEntry.WithError(err).Error("身份信息未核验 err")
			err = perror.BizErrOpreate.SetMessage("操作失败：身份信息未核验")
			return
		}
	} else if strings.Contains(reason, "姓名") {
		// 姓名不匹配：需选择具体乘客
		err = chromedp.Run(chromeCtx,
			chromedp.Sleep(waitKeyEventTime),
			chromedp.KeyEvent("8"),
			// 选择具体乘客：默认选择第一个即可
			chromedp.Sleep(waitKeyEventTime),
			chromedp.Click("//*[@id='SubFailReasonForPassengerForm']/span/input", chromedp.BySearch),
		)
		if err != nil {
			plogEntry.WithError(err).Error("姓名不匹配 err")
			err = perror.BizErrOpreate.SetMessage("操作失败：姓名不匹配")
			return
		}
	} else if strings.Contains(reason, "证件号码") {
		// 姓名不匹配：需选择具体乘客
		err = chromedp.Run(chromeCtx,
			chromedp.Sleep(waitKeyEventTime),
			chromedp.KeyEvent("9"),
			// 选择具体乘客：默认选择第一个即可
			chromedp.Sleep(waitKeyEventTime),
			chromedp.Click("//*[@id='SubFailReasonForPassengerForm']/span/input", chromedp.BySearch),
		)
		if err != nil {
			plogEntry.WithError(err).Error("姓名不匹配 err")
			err = perror.BizErrOpreate.SetMessage("操作失败：姓名不匹配")
			return
		}
	} else if strings.Contains(reason, "停运") {
		// 列车停运
		err = chromedp.Run(chromeCtx,
			chromedp.Sleep(waitKeyEventTime),
			chromedp.Click("//*[@id='FailResonGroup']/span[14]/label/input", chromedp.BySearch),
		)
		if err != nil {
			plogEntry.WithError(err).Error("列车停运 err")
			err = perror.BizErrOpreate.SetMessage("操作失败：列车停运")
			return
		}
	} else if strings.Contains(reason, "停止售票") ||
		strings.Contains(reason, "止售") {
		// 已停止售票
		err = chromedp.Run(chromeCtx,
			chromedp.Sleep(waitKeyEventTime),
			chromedp.Click("//*[@id='FailResonGroup']/span[15]/label/input", chromedp.BySearch),
		)
		if err != nil {
			plogEntry.WithError(err).Error("已停止售票 err")
			err = perror.BizErrOpreate.SetMessage("操作失败：已停止售票")
			return
		}
	} else if strings.Contains(reason, "价格不符") {
		// 价格不符
		err = chromedp.Run(chromeCtx,
			chromedp.Sleep(waitKeyEventTime),
			chromedp.Click("//*[@id='FailResonGroup']/span[16]/label/input", chromedp.BySearch),
		)
		if err != nil {
			plogEntry.WithError(err).Error("价格不符 err")
			err = perror.BizErrOpreate.SetMessage("操作失败：价格不符")
			return
		}
	} else if strings.Contains(reason, "进京") {
		// 不可进京
		err = chromedp.Run(chromeCtx,
			chromedp.Sleep(waitKeyEventTime),
			chromedp.Click("//*[@id='FailResonGroup']/span[12]/label/input", chromedp.BySearch),
		)
		if err != nil {
			plogEntry.WithError(err).Error("不可进京 err")
			err = perror.BizErrOpreate.SetMessage("操作失败：不可进京")
			return
		}
	} else if strings.Contains(reason, "车次不存在") {
		// 车次不存在
		err = chromedp.Run(chromeCtx,
			chromedp.Sleep(waitKeyEventTime),
			chromedp.Click("//*[@id='cc']", chromedp.BySearch),
		)
		if err != nil {
			plogEntry.WithError(err).Error("车次不存在 err")
			err = perror.BizErrOpreate.SetMessage("操作失败：车次不存在")
			return
		}
	} else if strings.Contains(reason, "预售期") {
		// 未到预售期
		err = chromedp.Run(chromeCtx,
			chromedp.Sleep(waitKeyEventTime),
			chromedp.Click("//*[@id='xx']", chromedp.BySearch),
		)
		if err != nil {
			plogEntry.WithError(err).Error("未到预售期 err")
			err = perror.BizErrOpreate.SetMessage("操作失败：未到预售期")
			return
		}
	} else if strings.Contains(reason, "封站") {
		// 出站已封站
		err = chromedp.Run(chromeCtx,
			chromedp.Sleep(waitKeyEventTime),
			chromedp.Click("//*[@id='x']", chromedp.BySearch),
		)
		if err != nil {
			plogEntry.WithError(err).Error("出站已封站 err")
			err = perror.BizErrOpreate.SetMessage("操作失败：出站已封站")
			return
		}
	} else {
		// 车次已无票
		err = HandlerCallbackResultNoTicketOfNoTicket(chromeCtx, transferNo)
		if err != nil {
			plogEntry.WithError(err).Error("车次已无票 err")
			err = perror.BizErrOpreate.SetMessage("操作失败：车次已无票")
			return
		}
	}
	
	// 设置占座失败-点击确定
	fmt.Println("设置占座失败-点击确定 selector:", selector)
	err = chromedp.Run(chromeCtx,
		chromedp.Click("//*[@id='SetBookFailPanel']/div[1]/div/div[3]/button[2]", chromedp.BySearch),
	)
	if err != nil {
		plogEntry.WithError(err).Error("设置占座失败-点击确定 err")
		log.Fatal(constants.CallBackErrConfirmFail.Key(), constants.CallBackErrConfirmFail.Value())
	}
	
	select {}
}

// 车次已无票
func HandlerCallbackResultNoTicketOfNoTicket(chromeCtx context.Context, transferNo int32) (err error) {
	logID := idutil.UniqIDV3()
	plogEntry := plog.LogID(logID).
		WithField("method", "RPAOrderTaskService.HandlerCallbackResultNoTicket")
	plogEntry.Info("request")
	
	if transferNo <= 1 {
		// 车次已无票
		// 选择原因并按键: 按下，并且松开
		err = chromedp.Run(chromeCtx,
			chromedp.Sleep(waitKeyEventTime),
			chromedp.KeyEvent("1"),
		)
		if err != nil {
			plogEntry.WithError(err).Error("车次已无票 err")
			err = perror.BizErrOpreate.SetMessage("操作失败：车次已无票")
			return
		}
	} else {
		// 第二程无票：需选择具体乘客
		err = chromedp.Run(chromeCtx,
			chromedp.Sleep(waitKeyEventTime),
			chromedp.Click("//*[@id='B']", chromedp.BySearch),
		)
		if err != nil {
			plogEntry.WithError(err).Error("第二程无票 err")
			err = perror.BizErrOpreate.SetMessage("操作失败：第二程无票")
			return
		}
	}
	return
}
