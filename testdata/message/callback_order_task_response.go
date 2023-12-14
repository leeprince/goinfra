package message

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/6/18 14:27
 * @Desc:
 */

type CallbackOrderTaskResponse struct {
	OrderID    string      `json:"orderID"`    // 订单ID
	ResultType ResultType  `json:"resultType"` // 订单任务处理结果类型:购票成功(Success)、无满足车票(NoTicket)、任务暂停(Suspend)
	Message    string      `json:"message"`    // resultType的简要说明
	Data       interface{} `json:"data"`       // 订单任务处理结果不同，data对应不同结构体
}

// 订单任务处理结果类型:购票成功(Success)、无满足车票(NoTicket)、任务暂停(Suspend)
type ResultType string

const (
	ResultTypeSuccess  ResultType = "Success"
	ResultTypeNoTicket ResultType = "NoTicket"
	ResultTypeSuspend  ResultType = "Suspend"
)

// ResultType=ResultTypeSuccess 的Data
type SuccessData struct {
	TicketNumber  string      `json:"ticketNumber"`  // 取票号
	PassengerList []Passenger `json:"passengerList"` // 乘客列表
}

// ResultType=ResultTypeNoTicket 的Data
type NoTicketData struct {
	TransferNo int32  `json:"transferNo"` // 第几程无满足车票，最大2程；1(第1程);2(第2程)-int32
	Reason     string `json:"reason"`     // 无满足车票的原因（实际上是弹窗内容）
}

// ResultType=ResultTypeSuspend 的Data
type SuspendData struct {
	ReasonType    SuspendReasonType `json:"reasonType"`    // 暂停原因类型: WaitUserPay(占座票等待用户付款，返回乘客占座信息，后端web继续发起继续扣款任务)、NoOneTicket(等待单人单程有要求甩票后无满足条件订单，需操作员手动操作)
	PassengerList []Passenger       `json:"passengerList"` // 乘客列表
}

type SuspendReasonType string

// 暂停原因类型: WaitUserPay(占座票等待用户付款，返回乘客占座信息，后端web继续发起继续扣款任务)
const (
	SuspendReasonTypeWaitUserPay    SuspendReasonType = "WaitUserPay"
	SuspendReasonTypeNoSameCarriage SuspendReasonType = "NoSameCarriage"
)
