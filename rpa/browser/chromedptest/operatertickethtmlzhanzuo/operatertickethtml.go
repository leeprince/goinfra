package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/leeprince/goinfra/utils/fileutil"
	"github.com/leeprince/goinfra/utils/stringutil"
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
	ticketHtmlFileDir  = "/Users/leeprince/www/go/goinfra/rpa/browser/chromedptest/operatertickethtmlzhanzuo"
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
		
		fileBytes, err := fileutil.ReadFile(ticketHtmlFileDir, "ticket.html")
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
	
	// 开始 rpa操作
	/*
		{
		    "orderID": "订单ID-string(32)",
		    "resultType": "订单任务处理结果类型：购票成功(Success)、无满足车票(NoTicket)、任务暂停(Suspend)、-string(18)",
		    "message": "resultType的简要说明",
		    "data":"订单任务处理结果不同，data对应不同结构体-object"
		}
	
		---
		# 购票成功(Success)的data结构体
		{
		    "ticketNumber": "取票号-string(12)",
		    "passengerList": [
		        {
		            "passengerId": "乘客ID-string(32)",
		            "creditType": "证件类型:ED(居民身份证)；LS(临时身份证)；WJ(警官证)；JG(军官证)；YW(义务兵证)；SG(士官证)；WG(文职干部证)；WY(文职人员证)；WH(外国人护照，需选择国家)；HZ(中国护照)；GN(港澳居民来往内地通行证)；QT(其他)。暂仅支持ED(居民身份证)",
		            "creditTypeName": "证件类型名称",
		            "creditNo": "证件号",
		            "fullName": "乘客姓名-string(32)",
		            "ticketType": "票种:Adult(成人票)、Child(小孩票)",
		            "seatType": "座位类型",
		            "carriage": "车厢号",
		            "seatNumber": "座位号",
		            "sleeper": "确定的卧铺位置：None(无)；Up(上)；Mid(中)；Down(下)"
		            "seatPrice": "该乘客一程票价格（单位角）",
		        }
		    ]
		}
	
	
		# 无满足车票(NoTicket)的data结构体
		{
		    "noTicketType": "无满足车票的占座失败类型：Other(其他)、TrainNoTicket(车次无票)、SeatNo(坐席无法满足)、UserNameNoMatch(姓名不匹配)、TrainNoExist(车次不存在)、TrainShutdown(列车停运)、TrainStopped(已停止售票)、-string(20)",
		    "otherTypeContext":"'其他'占座失败类型的原因说明-string(255)"
		}
	
		# 任务暂停(Suspend)的data结构体
		{
		    "reasonType":"暂停原因类型: WaitUserPay(占座票等待用户付款，后web端继续发起继续扣款任务)、NoOneTicket(等待单人单程有要求甩票后无满足条件订单，需操作员手动操作)-string(18)"
		}
	*/
	data := `{
	    "orderID": "MT1686715100250001",
	    "resultType": "Success",
	    "message": "订单任务处理结果类型：购票成功(Success)、无满足车票(NoTicket)、任务暂停(Suspend)",
	    "data":
	    {
	        "ticketNumber": "ticketNumber=prince001",
	        "passengerList":
	        [
	            {
	                "passengerId": "22268659_94601872",
	                "creditType": "证件类型:ED(居民身份证)；LS(临时身份证)；WJ(警官证)；JG(军官证)；YW(义务兵证)；SG(士官证)；WG(文职干部证)；WY(文职人员证)；WH(外国人护照，需选择国家)；HZ(中国护照)；GN(港澳居民来往内地通行证)；QT(其他)。暂仅支持ED(居民身份证)",
	                "creditTypeName": "证件类型名称",
	                "creditNo": "640302XXXXXXXX0362",
	                "fullName": "user1",
	                "ticketType": "票种:Adult(成人票)、Child(小孩票)",
	                "seatType": "座位类型",
	                "carriage": "03",
	                "seatNumber": "02A",
	                "sleeper": "Down",
	                "seatPrice": 2795
	            },
	            {
	                "passengerId": "22268659_94601873",
	                "creditType": "证件类型:ED(居民身份证)；LS(临时身份证)；WJ(警官证)；JG(军官证)；YW(义务兵证)；SG(士官证)；WG(文职干部证)；WY(文职人员证)；WH(外国人护照，需选择国家)；HZ(中国护照)；GN(港澳居民来往内地通行证)；QT(其他)。暂仅支持ED(居民身份证)",
	                "creditTypeName": "证件类型名称",
	                "creditNo": "640381XXXXXXXX0343",
	                "fullName": "user2",
	                "ticketType": "票种:Adult(成人票)、Child(小孩票)",
	                "seatType": "座位类型",
	                "carriage": "03",
	                "seatNumber": "03A",
	                "sleeper": "Mid",
	                "seatPrice": 2795
	            }
	        ]
	    }
	}`
	
	// 解析数据
	var result CallbackOrderTaskResult
	err = json.Unmarshal([]byte(data), &result)
	if err != nil {
		log.Fatal("解析data为CallbackOrderTaskResult错误:", err)
	}
	fmt.Println("result:", result)
	
	orderId := result.OrderID
	fmt.Println("orderId:", orderId)
	// 开始输入
	if result.ResultType == string(ResultTypeSuccess) {
		resultTypeSuccessData := &ResultTypeSuccessData{}
		result.Data = resultTypeSuccessData
		err = json.Unmarshal([]byte(data), &result)
		if err != nil {
			log.Fatal("解析data为CallbackOrderTaskResult.resultTypeSuccessData错误:", err)
		}
		
		fmt.Println("resultTypeSuccessData:", resultTypeSuccessData)
		
		resultTypeSuccessData, resultTypeSuccessDataOk := result.Data.(*ResultTypeSuccessData)
		if !resultTypeSuccessDataOk {
			log.Fatal("断言ResultTypeSuccessData错误:", err)
		}
		
		// 设置取票号：占座订单支付后才有取票号，暂注释
		// ticketNumber := resultTypeSuccessData.TicketNumber
		// ticketNumberId := "#EOrderNumberInput" + orderId
		// // 等待 ID 出现
		// chromedp.WaitVisible(ticketNumberId)
		// fmt.Println("ticketNumberId:", ticketNumberId)
		// err = chromedp.Run(ctx,
		// 	chromedp.SetValue(ticketNumberId, ticketNumber, chromedp.ByID),
		// )
		// if err != nil {
		// 	log.Fatal("设置取票号失败：", err)
		// }
		
		for _, passenger := range resultTypeSuccessData.PassengerList {
			fmt.Println("passenger:", passenger)
			
			passengerId := passenger.PassengerId
			creditNo := passenger.CreditNo
			carriage := passenger.Carriage
			seatNumber := passenger.SeatNumber
			sleeper := passenger.Sleeper
			// seatPrice := passenger.SeatPrice
			
			fmt.Println("passengerId:", passengerId)
			// err = chromedp.Run(ctx,
			// 	chromedp.WaitVisible(fmt.Sprintf("#%s", passengerId), chromedp.ByID),
			// )
			// if err != nil {
			// 	log.Fatal("WaitVisible err:", fmt.Sprintf("#%s", passengerId), err)
			// }
			// fmt.Println("WaitVisible end:", fmt.Sprintf("#%s", passengerId))
			//
			
			var (
				haveCreditNoText string
				dom              string
			)
			
			// 校验乘客身份证
			// `//*[@id="22268659_94601872"]/td[3]`
			dom = fmt.Sprintf(`//*[@id='%s']/td[3]`, passengerId)
			fmt.Println("dom:", dom)
			err = chromedp.Run(ctx, chromedp.Text(dom, &haveCreditNoText, chromedp.BySearch))
			if err != nil {
				log.Fatal("EvaluateAsDevTools haveCreditNoText err:", err)
			}
			fmt.Println("haveCreditNoText:", haveCreditNoText)
			if haveCreditNoText == "" {
				log.Fatal("身份证号-找不到")
			}
			haveCreditNo := GetCreditNo(haveCreditNoText)
			fmt.Println("haveCreditNo:", haveCreditNo)
			if haveCreditNo != creditNo {
				log.Fatal("身份证号-不匹配")
			}
			log.Println("身份证号-匹配")
			
			// 设置车厢号：成功
			/*
				firefox xpath:/html/body/span/span/table/tbody/tr[3]/td[7]/input
				chrome xpath://*[@id="21998005_1938898"]/td[6]/input
			*/
			fmt.Println("carriage:", carriage)
			dom = fmt.Sprintf(`//*[@id="%s"]/td[6]/input`, passengerId)
			fmt.Println("dom:", dom)
			err = chromedp.Run(ctx,
				chromedp.SetValue(dom, carriage, chromedp.BySearch),
			)
			if err != nil {
				log.Fatal("SetValue carriage err:", err)
			}
			fmt.Println("carriage")
			
			// 设置座位号：成功
			// 复制的 xpath
			/*
				firefox xpath:/html/body/span/span/table/tbody/tr[3]/td[7]/input
				chrome xpath://*[@id="21998005_1938898"]/td[7]/input
			*/
			fmt.Println("seatNumber:", seatNumber)
			dom = fmt.Sprintf(`//*[@id="%s"]/td[7]/input`, passengerId)
			fmt.Println("dom:", dom)
			err = chromedp.Run(ctx,
				chromedp.SetValue(dom, seatNumber, chromedp.BySearch),
			)
			if err != nil {
				log.Fatal("SetValue seatNumber err:", err)
			}
			fmt.Println("seatNumber")
			
			fmt.Println("sleeper:", sleeper)
			if sleeper != "" && sleeper != "None" {
				// 确定的卧铺位置：None(无)；Up(上)；Mid(中)；Down(下)
				var sleeperValue string
				switch sleeper {
				case "Up":
					sleeperValue = "上铺"
				case "Mid":
					sleeperValue = "中铺"
				case "Down":
					sleeperValue = "下铺"
				default:
					log.Fatal("sleeper 无匹配")
				}
				dom = fmt.Sprintf(`//*[@id="%s"]/td[7]/select`, passengerId)
				err = chromedp.Run(ctx,
					// v等于select 中 option 展示的值，而不是 value
					chromedp.SendKeys(dom, sleeperValue, chromedp.BySearch),
				)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		
	}
	
	select {}
}

// 未处理的身份证号信息: 441521 20010116 XXXX  复制 => 44152120010116XXXX
func GetCreditNo(s string) string {
	s = stringutil.ReplaceWhitespaceChar(s)
	orderInfoRune := []rune(s)
	return string(orderInfoRune[:len(orderInfoRune)-2])
}
