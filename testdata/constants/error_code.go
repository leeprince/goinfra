package constants

import "github.com/leeprince/goinfra/consts/constval"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/5/20 00:49
 * @Desc:
 */

var (
	
	// 处理成功:0
	Success = constval.NewInt32(0, "成功响应")
	Fail    = constval.NewInt32(1, "系统异常，请重试！")
	
	// 基础错误:11xxxx
	BaseErrParamsInvalid    = constval.NewInt32(110001, "无效参数!")
	BaseErrParamsRequired   = constval.NewInt32(110002, "参数必填!")
	BaseErrDataNull         = constval.NewInt32(110003, "数据为空!")
	BaseErrFind             = constval.NewInt32(110004, "查询错误!")
	BaseErrCreate           = constval.NewInt32(110005, "创建错误!")
	BaseErrUpdate           = constval.NewInt32(110006, "更新错误!")
	BaseErrDelete           = constval.NewInt32(110007, "删除错误!")
	BaseErrSign             = constval.NewInt32(110008, "签名错误!")
	BaseErrGetAccessToken   = constval.NewInt32(110009, "获取access_token失败!")
	BaseErrCheckAccessToken = constval.NewInt32(110010, "检查access_token失败!")
	BaseErrConfig           = constval.NewInt32(110011, "配置错误!")
	BaseErrDataParse        = constval.NewInt32(110012, "数据解析错误!")
	BaseErrEventPublish     = constval.NewInt32(110013, "事件发布失败!")
	BaseErrLimit            = constval.NewInt32(110014, "检查限流错误!")
	BaseErrFieldType        = constval.NewInt32(110015, "字段类型转换错误!")
	BaseErrInstance         = constval.NewInt32(110016, "获取实例失败!")
	BaseErrGetToken         = constval.NewInt32(110017, "获取token失败!")
	BaseErrCheckToken       = constval.NewInt32(110018, "检查token失败!")
	BaseErrGetPassword      = constval.NewInt32(110019, "获取密码错误!")
	BaseErrSetPassword      = constval.NewInt32(110020, "设置密码错误!")
	BaseErrResponse         = constval.NewInt32(110021, "响应错误!")
	BaseErrInternal         = constval.NewInt32(119999, "内部服务器错误!")
	
	// 登录错误:12xxxx
	LoginErrOpenLogin          = constval.NewInt32(120001, "打开登录网站错误!")
	LoginErrLoginInput         = constval.NewInt32(120002, "输入账号和密码!")
	LoginErrClickLogin         = constval.NewInt32(120003, "点击登录!")
	LoginErrWaitSetFloatWindow = constval.NewInt32(120003, "等待并设置辅助弹窗失败!")
	LoginErrSetFloatWindow     = constval.NewInt32(120004, "设置辅助弹窗失败!")
	
	// 写入数据错误:13xxxx
	WriteErrNoOK = constval.NewInt32(130001, "写入数据失败!")
	WriteErrErr  = constval.NewInt32(130001, "写入数据错误!")
	
	// 订单解析处理:14xxxx
	OrderErrOrderLen = constval.NewInt32(140001, "出发时间、车次、出发地及目的地、座位类型数组的长度不一致!")
	
	// 订单发送:15xxxx
	SendOrderErrSend = constval.NewInt32(150001, "orderID不匹配!")
	
	// 订单回调处理:16xxxx
	CallBackErrOrderIDNoMatch         = constval.NewInt32(160001, "回调orderID与当前任务orderID不匹配!")
	CallBackErrResultType             = constval.NewInt32(160002, "回调类型不存在!")
	CallBackErrSelect                 = constval.NewInt32(160003, "查找选择器不存在或者超时!")
	CallBackErrNoCurrOrderTask        = constval.NewInt32(160004, "不存在当前订单任务!")
	CallBackErrPassengerListNum       = constval.NewInt32(160005, "回调乘客列表数量与当前任务订单下所有乘客列表数量不匹配!")
	CallBackErrReasonType             = constval.NewInt32(160006, "暂停原因类型不存在!")
	CallBackErrReasonTypeNoOccupySeat = constval.NewInt32(160007, "暂停原因类型为占座等待付款，但是当前任务为非占座不存在!")
	CallBackErrSetValue               = constval.NewInt32(160008, "设置值错误!")
	CallBackErrCreditNo               = constval.NewInt32(160009, "证件号不匹配!")
	CallBackErrTicketNumber           = constval.NewInt32(160010, "获取取票号错误!")
	CallBackErrOccupy                 = constval.NewInt32(160011, "点击出票成功:占座票失败!")
	CallBackErrNoOccupy               = constval.NewInt32(160012, "点击出票成功:非占座票!")
	CallBackErrFailReason             = constval.NewInt32(160013, "选择原因失败!")
	CallBackErrConfirmFail            = constval.NewInt32(160014, "设置占座失败-点击确定失败!")
	CallBackErrClickAliPayTradeNo     = constval.NewInt32(160015, "点击刷新流水号-失败!")
	CallBackErrGetAliPayTradeNo       = constval.NewInt32(160016, "获取刷新流水号-失败!")
	CallBackErrNoOrderTaskList        = constval.NewInt32(160016, "订单任务列表为空!")
	
	// 弹窗处理:17xxxx
	WinErrSeatTypeNoMatch             = constval.NewInt32(170001, "该窗口类型暂不支持该座位类型!")
	WinErrSeatTypePositionTypeNoMatch = constval.NewInt32(170002, "该窗口类型暂不支持该座位类型下的要求!")
	WinErrSeatType                    = constval.NewInt32(170003, "暂不支持该座位类型!")
	WinErrSeatTypeRequired            = constval.NewInt32(170004, "暂不支持该座位要求，请联系系统管理员!")
	WinErrSeatCarriage                = constval.NewInt32(170005, "暂不支持乘客指定车厢!")
	WinErrSeatSeatNumber              = constval.NewInt32(170006, "暂不支持乘客指定座位号!")
	WinErrSeatSleeper                 = constval.NewInt32(170007, "暂不支持乘客指定上中下铺!")
	WinErrSeatMatch                   = constval.NewInt32(170008, "暂不支持该座位要求的匹配，请联系系统管理员!")
	WinErrSeatSleeperMatch            = constval.NewInt32(170009, "上中下铺匹配失败，请联系系统管理员!")
	WinErrSeatMatchSleeper            = constval.NewInt32(1700010, "座位类型非卧铺不可以指定上中下铺，请联系系统管理员!")
	WinErrOccupySeat                  = constval.NewInt32(1700011, "暂不支持占座票，请联系系统管理员!")
)
