package testdata

import (
	"encoding/json"
	"fmt"
	"github.com/leeprince/goinfra/test/message"
	"github.com/leeprince/goinfra/utils/jsonutil"
	"github.com/spf13/cast"
	"log"
	"testing"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/2 10:40
 * @Desc:
 */

// 当结构体中定义 Any 类型时，需要根据不同的场景，解析称不同的结构体
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
	                "seatPrice": 1000
	            }
	        ]
	    }
	}`

	// 解析数据
	var result message.CallbackOrderTaskResponse
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		log.Fatal("解析data为CallbackOrderTaskResult错误:", err)
	}

	fmt.Println("result:", result)

	// 开始输入
	if result.ResultType == message.ResultTypeSuccess {
		// must be a pointer to an interface or to a type implementing the interface
		// 注意 result.Data 是接口类型，所以必须用指针类型进行赋值才可以断言类型得到正确结果即：&SuccessData{},且 &message.CallbackOrderTaskResponse
		// 正确
		result.Data = &message.SuccessData{}
		err = json.Unmarshal([]byte(data), &result)
		if err != nil {
			log.Fatal("解析data为 CallbackOrderTaskResponse.SuccessData 错误:", err)
		}
		resultTypeSuccessData, resultTypeSuccessDataOk := result.Data.(*message.SuccessData)
		if !resultTypeSuccessDataOk {
			log.Fatal("断言 SuccessData 错误:", resultTypeSuccessDataOk)
		}
		fmt.Println("resultTypeSuccessData:", resultTypeSuccessData)

		/*// 错误：原因 result.Data 是接口类型，所以必须用指针类型进行赋值才可以断言类型得到正确结果即：&SuccessData{},且 &message.CallbackOrderTaskResponse
		result.Data = &message.SuccessData{}
		err = json.Unmarshal([]byte(data), result)
		if err != nil {
			log.Fatal("解析 data CallbackOrderTaskResponse.SuccessData 1 错误:", err)
		}
		resultTypeSuccessData1, resultTypeSuccessDataOk := result.Data.(message.SuccessData)
		if !resultTypeSuccessDataOk {
			log.Println("断言 SuccessData 1 result:", result)
			log.Fatal("断言 SuccessData 1 错误:", resultTypeSuccessDataOk)
		}
		fmt.Println("SuccessData 1:", resultTypeSuccessData1)*/

		// 错误：原因 result.Data 是接口类型，所以必须用指针类型进行赋值才可以断言类型得到正确结果即：&SuccessData{},且 &message.CallbackOrderTaskResponse
		result.Data = message.SuccessData{}
		err = json.Unmarshal([]byte(data), &result)
		if err != nil {
			log.Fatal("解析 data CallbackOrderTaskResponse.SuccessData 2 错误:", err)
		}
		resultTypeSuccessData2, resultTypeSuccessDataOk := result.Data.(message.SuccessData)
		if !resultTypeSuccessDataOk {
			log.Println("断言 SuccessData 2 result:", result)
			log.Fatal("断言 SuccessData 2 错误:", resultTypeSuccessDataOk)
		}
		fmt.Println("SuccessData 2:", resultTypeSuccessData2)
	}

	time.Sleep(time.Second * 20)
}

func TestToMapList(t *testing.T) {
	/*jsonData := `
		[
	      {
	        "audit_status": 1,
	        "corp_id": "",
	        "created_at": 1636387643,
	        "del_status": 0,
	        "exclusive_time": 0,
	        "invoice_id": "6863507238099358784"
	      }
	    ]
		`*/
	jsonData := `[{"audit_status":1,"corp_id":"","created_at":1636387643,"del_status":0,"exclusive_time":0,"invoice_id":"6863507238099358784"}]`

	// 如果数组元素的字段类型不是string, 则会报错
	resp, err := convertToMap1([]byte(jsonData))
	fmt.Println("convertToMap1:", err)
	fmt.Println("convertToMap1:", resp)

	fmt.Println("-------------:")
	// > 成功
	resp, err = convertToMap2([]byte(jsonData))
	fmt.Println("convertToMap2:", err)
	fmt.Println("convertToMap2:", resp)
}

// 将 json 数组数据 转为 []map[string]string
// > 如果数组元素的字段类型不是string, 则会报错
func convertToMap1(data []byte) ([]map[string]string, error) {
	var invoiceMapList []map[string]string

	err := json.Unmarshal(data, &invoiceMapList)
	if err != nil {
		return nil, err
	}
	return invoiceMapList, nil
}

// 将 json 数组数据 转为 []map[string]string
// > 成功
func convertToMap2(data []byte) ([]map[string]string, error) {
	var invoiceMapListValueAny []map[string]interface{}
	err := jsonutil.JsoniterCompatible.Unmarshal(data, &invoiceMapListValueAny)
	if err != nil {
		return nil, err
	}

	var invoiceMapList []map[string]string
	for _, mapList := range invoiceMapListValueAny {
		item := make(map[string]string)
		for key, value := range mapList {
			item[key] = cast.ToString(value)
		}
		invoiceMapList = append(invoiceMapList, item)
	}

	return invoiceMapList, nil
}

func TestJson(t *testing.T) {
	dataJson := `{
	    "code":0,
	    "message":"说明-string",
	    "log_id":"日志ID",
	    "data":null
	}`

	data := message.VerifyAuthResp{}
	err := json.Unmarshal([]byte(dataJson), &data)
	fmt.Println("err", err)
	fmt.Println("data", data)
}
