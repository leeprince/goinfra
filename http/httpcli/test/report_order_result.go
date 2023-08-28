package test

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/26 15:08
 * @Desc:
 */

type ReportOrderResultReq struct {
	OrderId        string             `json:"orderId"`        // 订单ID;size:32;
	TicketNumber   string             `json:"ticketNumber"`   // 订单取票号;size:12;
	ContactPhone   string             `json:"contactPhone"`   // 订单联系电话;size:11;
	IsTransfer     RORRIsTransfer     `json:"isTransfer"`     // 是否换乘.0:false;1:true;
	IsOccupySeat   RORRIsOccupySeat   `json:"isOccupySeat"`   // 是否占座.0:false;1:true;
	CompleteStatus RORRCompleteStatus `json:"completeStatus"` // 完成状态.0:未完成;1:出票成功;2:出票失败;
	FailReason     string             `json:"failReason"`     // 出票失败原因;size:255;
	MachineName    string             `json:"machineName"`    // 设备名称;size:128;
}

// 是否换乘.0:false;1:true
type RORRIsTransfer int

var (
	RORRIsTransferFalse RORRIsTransfer = 0
	RORRIsTransferTrue  RORRIsTransfer = 1
)

// 是否占座.0:false;1:true
type RORRIsOccupySeat int

var (
	RORRIsOccupySeatFalse RORRIsOccupySeat = 0
	RORRIsOccupySeatTrue  RORRIsOccupySeat = 1
)

// 完成状态.0:未完成;1:出票成功;2:出票失败
type RORRCompleteStatus int

var (
	RORRCompleteStatusPre  RORRCompleteStatus = 0
	RORRCompleteStatusSucc RORRCompleteStatus = 1
	RORRCompleteStatusFail RORRCompleteStatus = 2
)
