package test

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/2 10:52
 * @Desc:
 */

type CallbackOrderTaskResult struct {
	OrderID    string      `json:"orderID"`    // 订单ID
	ResultType string      `json:"resultType"` // 订单任务处理结果类型：购票成功(Success)、无满足车票(NoTicket)、任务暂停(Suspend)
	Message    string      `json:"message"`    // resultType的简要说明
	Data       interface{} `json:"data"`       // 订单任务处理结果不同，data对应不同结构体-object
}

// 订单任务处理结果类型：购票成功(Success)、无满足车票(NoTicket)、任务暂停(Suspend)
type ResultType string

const (
	ResultTypeSuccess  ResultType = "Success"
	ResultTypeNoTicket ResultType = "NoTicket"
	ResultTypeSuspend  ResultType = "Suspend"
)

// resultType=购票成功(Success)的data结构体
type ResultTypeSuccessData struct {
	TicketNumber  string      `json:"ticketNumber"` // 取票号-string(12)
	PassengerList []Passenger `json:"passengerList"`
}

type Passenger struct {
	PassengerId    string `json:"passengerId"`    // 乘客ID
	CreditType     string `json:"creditType"`     // 证件类型:ED(居民身份证)；LS(临时身份证)；WJ(警官证)；JG(军官证)；YW(义务兵证)；SG(士官证)；WG(文职干部证)；WY(文职人员证)；WH(外国人护照，需选择国家)；HZ(中国护照)；GN(港澳居民来往内地通行证)；QT(其他)。暂仅支持ED(居民身份证)
	CreditTypeName string `json:"creditTypeName"` // 证件类型名称
	CreditNo       string `json:"creditNo"`       // 证件号
	FullName       string `json:"fullName"`       // 乘客姓名-string(32)
	TicketType     string `json:"ticketType"`     // 票种:Adult(成人票)、Child(小孩票)
	SeatType       string `json:"seatType"`       // 座位类型
	Carriage       string `json:"carriage"`       // 车厢号
	SeatNumber     string `json:"seatNumber"`     // 座位号
	Sleeper        string `json:"sleeper"`        // 确定的卧铺位置：None(无)；Up(上)；Mid(中)；Down(下)
	SeatPrice      string `json:"seatPrice"`      // 该乘客一程票价格（单位角）
}

// resultType=无满足车票(NoTicket)的data结构体
type ResultTypeNoTicketData struct {
	NoTicketType     string `json:"noTicketType"`     // 无满足车票的占座失败类型：Other(其他)、TrainNoTicket(车次无票)、SeatNo(坐席无法满足)、UserNameNoMatch(姓名不匹配)、TrainNoExist(车次不存在)、TrainShutdown(列车停运)、TrainStopped(已停止售票)、-string(20)
	OtherTypeContext string `json:"otherTypeContext"` // '其他'占座失败类型的原因说明-string(255)
}

// resultType=任务暂停(Suspend)的data结构体
type ResultTypeSuspendData struct {
	ReasonType string `json:"reasonType"` // 暂停原因类型: WaitUserPay(占座票等待用户付款，后端web继续发起继续扣款任务)、NoOneTicket(等待单人单程有要求甩票后无满足条件订单，需操作员手动操作)-string(18)
}
