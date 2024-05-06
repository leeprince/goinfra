package constants

import "github.com/leeprince/goinfra/consts/constval"

/**
 * @Author: prince.lee
 * @Date:   2022/6/15 18:03F
 * @Desc:   响应错误码
 */

var (
	// 处理成功:0
	Success                    = constval.NewInt(0, "成功响应")
	BaseErrParamsInvalid       = constval.NewInt(210001, "无效参数!")
	BaseErrParamsRequired      = constval.NewInt(210002, "参数必填!")
	BaseErrDataNull            = constval.NewInt(210003, "数据为空!")
	BaseErrInsert              = constval.NewInt(210004, "新增错误!")
	BaseErrDelete              = constval.NewInt(210005, "删除错误!")
	BaseErrUpdate              = constval.NewInt(210006, "更新错误!")
	BaseErrSign                = constval.NewInt(210007, "签名错误!")
	BaseErrFind                = constval.NewInt(210008, "查询错误!")
	BaseErrGetToken            = constval.NewInt(210009, "获取access_token失败!")
	BaseErrCheckToken          = constval.NewInt(210010, "检查access_token失败!")
	BaseErrConfig              = constval.NewInt(210011, "配置错误!")
	BaseErrDataParse           = constval.NewInt(210012, "数据解析错误!")
	BaseErrEventPublish        = constval.NewInt(210013, "事件发布失败!")
	BaseErrEventSubscribe      = constval.NewInt(210014, "事件订阅失败!")
	BaseErrLimit               = constval.NewInt(210015, "限流中!")
	BaseErrFieldType           = constval.NewInt(210016, "字段类型转换错误!")
	BaseErrInstance            = constval.NewInt(210017, "获取实例失败!")
	BaseErrNoticeSendFail      = constval.NewInt(210018, "通知发送失败!")
	BaseErrNoticeSmsSendFail   = constval.NewInt(210019, "短信通知发送失败!")
	BaseErrNoticeEmailSendFail = constval.NewInt(210020, "邮件通知发送失败!")
	
	// --- 业务逻辑错误
	// 通用业务逻辑错误
	BusinessLogicErr = constval.NewInt(221000, "处理失败，请重试!")
	
	// --- 业务逻辑错误-end
	BusinessLogicErrRiskControl = constval.NewInt(221001, "触发了风控规则")
	
	// 异常错误:25xxxx
	PanicErrInternalService = constval.NewInt(250000, "服务繁忙,请稍后重试")
)
