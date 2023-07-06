package test

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/2 10:40
 * @Desc:
 */

func TestJsonDataInterface(t *testing.T) {
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
	    "orderID": "T202306100954030604914284",
	    "resultType": "Success",
	    "message": "订单任务处理结果类型：购票成功(Success)、无满足车票(NoTicket)、任务暂停(Suspend)",
	    "data":
	    {
	        "ticketNumber": "prince001",
	        "passengerList":
	        [
	            {
	                "passengerId": "21998005_1938898",
	                "creditType": "证件类型:ED(居民身份证)；LS(临时身份证)；WJ(警官证)；JG(军官证)；YW(义务兵证)；SG(士官证)；WG(文职干部证)；WY(文职人员证)；WH(外国人护照，需选择国家)；HZ(中国护照)；GN(港澳居民来往内地通行证)；QT(其他)。暂仅支持ED(居民身份证)",
	                "creditTypeName": "证件类型名称",
	                "creditNo": "证件号",
	                "fullName": "余余余",
	                "ticketType": "票种:Adult(成人票)、Child(小孩票)",
	                "seatType": "座位类型",
	                "carriage": "03",
	                "seatNumber": "02A",
	                "sleeper": "Down",
	                "seatPrice": "该乘客一程票价格（单位角）"
	            }
	        ]
	    }
	}`
	
	// 解析数据
	var result CallbackOrderTaskResult
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		log.Fatal("解析data为CallbackOrderTaskResult错误:", err)
	}
	
	fmt.Println("result:", result)
	
	// 开始输入
	if result.ResultType == string(ResultTypeSuccess) {
		// 注意 result.Data 是接口类型，所以必须用指针类型进行赋值才可以断言类型得到正确结果即：&ResultTypeSuccessData{}
		result.Data = &ResultTypeSuccessData{}
		err = json.Unmarshal([]byte(data), &result)
		if err != nil {
			log.Fatal("解析data为CallbackOrderTaskResult.resultTypeSuccessData错误:", err)
		}
		
		resultTypeSuccessData, resultTypeSuccessDataOk := result.Data.(*ResultTypeSuccessData)
		if !resultTypeSuccessDataOk {
			log.Fatal("断言ResultTypeSuccessData错误:", err)
		}
		
		fmt.Println("resultTypeSuccessData:", resultTypeSuccessData)
	}
	
	time.Sleep(time.Second * 20)
}
