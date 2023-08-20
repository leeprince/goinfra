package message

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/19 18:21
 * @Desc:
 */

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/6/13 09:22
 * @Desc:
 */

type OrderTask struct {
	OrderID      string       `json:"orderID"`      // 订单ID
	TaskType     TaskType     `json:"taskType"`     // 任务类型:(NewOrder(新订单)、ContinueTask(继续任务)
	ContinueType ContinueType `json:"continueType"` // 继续任务的类型，当taskType=ContinueTask时有效。继续任务的类型：选票(SelectTicket)-输入乘客列表信息(InputPassengerList)-扣款(Deduction)
	IsTransfer   bool         `json:"isTransfer"`   // 是否换乘,是换乘orderItem则有多条数据
	IsOccupySeat bool         `json:"isOccupySeat"` // 是否是占座票,是则输入完乘客信息后等待用户付款
	ContactPhone string       `json:"contactPhone"` // 联系电话
	TotalPrice   int32        `json:"totalPrice"`   // 订单总金额（单位角）
	CreateAt     string       `json:"createAt"`     // 创建时间
	OrderList    []OrderItem  `json:"orderItem"`    // 订单项
}

// 任务类型:(NewOrder(新订单)、ContinueTask(继续任务)
type TaskType string

const (
	TaskTypeNewOrder TaskType = "NewOrder"
	TaskTypeContinue TaskType = "ContinueTask"
)

// 继续任务的类型，当taskType=ContinueTask时有效。继续任务的类型：选票(SelectTicket)-输入乘客列表信息(InputPassengerList)-扣款(Deduction)
type ContinueType string

const (
	ContinueTypeNone               ContinueType = ""
	ContinueTypeNewOrder           ContinueType = "SelectTicket"
	ContinueTypeInputPassengerList ContinueType = "InputPassengerList"
	ContinueTypeDeduction          ContinueType = "Deduction"
)

type OrderItem struct {
	DepartureDate          string             `json:"departureDate"`          // 出发时间(2023-06-08)
	DepartureDateKey       string             `json:"departureDateKey"`       // 出发时间F1输入:08
	TrainCode              string             `json:"trainCode"`              // 车次
	DepartureStation       string             `json:"departureStation"`       // 出发站点名称
	DepartureStationKey    string             `json:"departureStationKey"`    // 出发站点名称-F3电报码:-SZQ)
	DestinationStation     string             `json:"destinationStation"`     // 目的站点名称
	DestinationStationKey  string             `json:"destinationStationKey"`  // 目的站点名称-F4电报码:-GGQ
	PersonNumber           int32              `json:"personNumber"`           // 人数(该程票数)
	ChildNum               int32              `json:"childNum"`               // 小孩数量
	AdultNum               int32              `json:"adultNum"`               // 成人数量,小孩数量+成人数量=len(passengersList)
	TotalPrice             int64              `json:"totalPrice"`             // 总金额(单位角)
	SeatType               SeatType           `json:"seatType"`               // 座位类型-None(无类型,不存在无类型的)、Business(商务座)、Superlative(特等座) 、AdvancedDynamicSleeper(高级动卧)、DynamicSleeper(动卧)、First(一等座)、Second(二等座)、SoftSleeper(软卧)、HardSleeper(硬卧)、Soft(软座)、Hard(硬座)、No(无座)-string
	SeatPositionTypeString string             `json:"seatPositionTypeString"` // 座位位置要求说明
	SeatPositionTypes      []SeatPositionType `json:"seatPositionType"`       // 座位位置要求类型
	IsAcceptNoSeat         bool               `json:"isAcceptNoSeat"`         // 是否接受无座
	SwingTicketNum         int32              `json:"swingTicketNum"`         // 甩票数，正常大于orderItem数组下的personNumber人数进行出票。甩票数>0时代表不依赖personNumber，而是按照该字段的值进行甩票，出票后按要求找到符合条件的票，删除不满足条件的票；如果甩出的票较多导致铁路局系统弹窗提示没有票则减 1 继续甩，直到等于personNumber。场景：一个人有要求时;指定车厢和座位号的时;
	PassengerList          []Passenger        `json:"passengerList"`          // 乘客列表
	OperatingStepsList     []OperatingSteps   `json:"operatingStepsList"`     // 操作的步骤。是一个二维数组，第一维数组用于重试；第二维数组用于一组选票操作且最后一项操作后需要检查是否有票。详细请看示例-[]operateSteps
}

// 座位类型-None(无类型,不存在无类型的)、Business(商务座)、Superlative(特等座) 、AdvancedDynamicSleeper(高级动卧)、DynamicSleeper(动卧)、First(一等座)、Second(二等座)、SoftSleeper(软卧)、Soft(软座)、HardSleeper(硬卧)、Hard(硬座)、No(无座)-string
type SeatType string

const (
	SeatTypeNone                   SeatType = "None"
	SeatTypeBusiness               SeatType = "Business"
	SeatTypeSuperlative            SeatType = "Superlative"
	SeatTypeAdvancedDynamicSleeper SeatType = "AdvancedDynamicSleeper"
	SeatTypeDynamicSleeper         SeatType = "DynamicSleeper"
	SeatTypeFirst                  SeatType = "First"
	SeatTypeSecond                 SeatType = "Second"
	SeatTypeSoftSleeper            SeatType = "SoftSleeper"
	SeatTypeSoft                   SeatType = "Soft"
	SeatTypeHardSleeper            SeatType = "HardSleeper"
	SeatTypeHard                   SeatType = "Hard"
	SeatTypeNo                     SeatType = "No"
)

// 座位位置要求类型
/*
   None(无要求，无要求时该数组为空)、
   Window(靠窗。对外使用SpecifySeat代替)、
   Aisle(靠过道。对外使用SpecifySeat代替)、
   SpecifySeat(匹配指定座位位置后缀，只需匹配其中一个后缀，且按给定的数组优先级升序，用于对售票系统靠窗、靠过道等后缀匹配要求统一字段。如：靠窗(火车["4","9","0","5"];高铁["A","F"])---靠过道(火车["8","3","7","2"];高铁["C","D"]))、
   SpecifySeatRequired(必须匹配座位位置后缀，必须匹配出所有后缀。如：一张D，一张F["D","F"])、
   Carriage(指定车厢)、
   Seat(指定座位号)、
   Sleeper(指定卧铺上中下)、
   CarriageSeat(指定车厢和座位号)、
   CarriageSeatSleeper(指定车厢、座位号和卧铺上中下)、
   AdjacentSeat(邻座)、
   SameCarriage(同车厢)、
   SameCarriageSleeper(同车厢，指定上下铺)、
   SameRoom(同包厢包间)、
*/

type SeatPositionType string

const (
	SeatPositionTypeNone                SeatPositionType = "None"
	SeatPositionTypeWindow              SeatPositionType = "Window"
	SeatPositionTypeAisle               SeatPositionType = "Aisle"
	SeatPositionTypeSpecifySeat         SeatPositionType = "SpecifySeat"
	SeatPositionTypeSpecifySeatRequired SeatPositionType = "SpecifySeatRequired"
	SeatPositionTypeCarriage            SeatPositionType = "Carriage"
	SeatPositionTypeSeat                SeatPositionType = "Seat"
	SeatPositionTypeSleeper             SeatPositionType = "Sleeper"
	SeatPositionTypeCarriageSeat        SeatPositionType = "CarriageSeat"
	SeatPositionTypeCarriageSeatSleeper SeatPositionType = "CarriageSeatSleeper"
	SeatPositionTypeAdjacentSeat        SeatPositionType = "AdjacentSeat"
	SeatPositionTypeSameCarriage        SeatPositionType = "SameCarriage"
	SeatPositionTypeSameCarriageSleeper SeatPositionType = "SameCarriageSleeper"
	SeatPositionTypeSameRoom            SeatPositionType = "SameRoom"
)

// 火车
var (
	// 靠窗
	SeatPositionTypeWindowNormal = []string{"4", "9", "0", "5"}
	// 靠过道
	SeatPositionTypeAisleNormal = []string{"8", "3", "7", "2"}
)

// 高级列车：动车、高铁
var (
	// 靠窗
	SeatPositionTypeWindowAdvanced = []string{"A", "F"}
	// 靠过道
	SeatPositionTypeAisleAdvanced = []string{"C", "D"}
	// 靠过道:商务座
	SeatPositionTypeAisleAdvancedSeatTypeBusiness = []string{"A", "C", "F"}
)

type Passenger struct {
	PassengerId    string     `json:"passengerId"`    // 乘客 Id
	CreditType     CreditType `json:"creditType"`     // 证件类型:ED(居民身份证)；LS(临时身份证)；WJ(警官证)；JG(军官证)；YW(义务兵证)；SG(士官证)；WG(文职干部证)；WY(文职人员证)；WH(外国人护照，需选择国家)；HZ(中国护照)；GN(港澳居民来往内地通行证)；QT(其他)。暂仅支持ED(居民身份证)
	CreditTypeName string     `json:"creditTypeName"` // 证件类型名称
	CreditNo       string     `json:"creditNo"`       // 证件号
	FullName       string     `json:"fullName"`       // 乘客姓名
	TicketType     TicketType `json:"ticketType"`     // 票种:Adult(成人票)、Child(小孩票)
	SeatType       SeatType   `json:"seatType"`       // 座位类型
	Carriage       string     `json:"carriage"`       // 车厢号
	SeatNumber     string     `json:"seatNumber"`     // 座位号
	Sleeper        Sleeper    `json:"sleeper"`        // 确定的卧铺位置：None(无)；Up(上)；Mid(中)；Down(下)
	SeatPrice      int32      `json:"seatPrice"`      // 该乘客一程票价格（单位角）
}

// 确定的卧铺位置：None(无)；Up(上)；Mid(中)；Down(下)
type Sleeper string

const (
	SleeperNone = "None"
	SleeperUp   = "Up"
	SleeperMid  = "Mid"
	SleeperDown = "Down"
)

var SleeperMap map[string]Sleeper = map[string]Sleeper{
	"上": SleeperUp,
	"中": SleeperMid,
	"下": SleeperDown,
}

// 票种:Adult(成人票)、Child(小孩票)
type TicketType string

const (
	TicketTypeAdult = "Adult"
	TicketTypeChild = "Child"
)

type OperatingSteps struct {
	OprateList              []Oprate         `json:"oprateList"`              //
	SeatType                SeatType         `json:"seatType"`                // 座位类型，同上面的seatType-string
	SeatPositionType        SeatPositionType `json:"seatPositionType"`        // 座位位置要求类型，同上面的seatPositionType。存在None(无要求)的情况，而其他要求需要配合其他字段匹配：input、matchPositionSuffixList、positionList
	MatchPositionSuffixList []string         `json:"matchPositionSuffixList"` // 按要求匹配座位号后缀,如要求是seatPositionType=Window(靠窗)["0","4","5","9"];高铁["A","F"]，seatPositionType=SpecifySeat(指定座位位置后缀)高铁D\F["D","F"]
	PositionList            []Position       `json:"positionList"`            // 按要求精准匹配座位,如要求是Carriage(指定⻋厢)、seatPositionType=Seat(指定座位号)、seatPositionType=CarriageSeat(指定⻋厢和座位号)、CarriageSeatSleeper(指定车厢、座位号和卧铺上中下)。如果是只是Carriage(指定⻋厢)且是高铁票则可以直接通过操作检查，而不需要匹配，因为匹配目前依赖 OCR
}

type Position struct {
	Carriage   string  `json:"carriage"`   // 车厢
	SeatNumber string  `json:"seatNumber"` // 座位号
	Sleeper    Sleeper `json:"Sleeper"`    // 确定的卧铺位置：None(无)；Up(上)；Mid(中)；Down(下)
}

type Oprate struct {
	SeatTypeKey string `json:"seatTypeKey"` // 按键，包含组合键，且都为大写，组合键的的连接符是加号+ -string
	Input       string `json:"input"`       // 输入值，如按键之后有些操作会自动到达输入框，这是就需要输入值。
}

// 证件类型:ED(居民身份证)；LS(临时身份证)；WJ(警官证)；JG(军官证)；YW(义务兵证)；SG(士官证)；WG(文职干部证)；WY(文职人员证)；WH(外国人护照，需选择国家)；HZ(中国护照)；GN(港澳居民来往内地通行证)；QT(其他)。暂仅支持ED(居民身份证)
type CreditType string

const (
	CreditTypeNone CreditType = "NONE"
	CreditTypeED   CreditType = "ED"
	CreditTypeLS   CreditType = "LS"
	CreditTypeWJ   CreditType = "WJ"
	CreditTypeJG   CreditType = "JG"
	CreditTypeYW   CreditType = "YW"
	CreditTypeSG   CreditType = "SG"
	CreditTypeWG   CreditType = "WG"
	CreditTypeWY   CreditType = "WY"
	CreditTypeWH   CreditType = "WH"
	CreditTypeHZ   CreditType = "HZ"
	CreditTypeGN   CreditType = "GN"
	CreditTypeQT   CreditType = "QT"
)

// 售票系统窗口类型;普通(Normal,硬座、软座、硬卧、软卧、硬座无座、指定铺别);高级(Advanced，有商务座、特等座、一等座、二等座、指定属性、指定席位)
// 与车次有关，判断`售票系统窗口类型`的方法是检查车次的前缀。不同的车次的前缀，会包含不用的座位类型(SeatType)。
type WindowType string

const (
	WindowTypeNormal   WindowType = "Normal"
	WindowTypeAdvanced WindowType = "Advanced"
)
