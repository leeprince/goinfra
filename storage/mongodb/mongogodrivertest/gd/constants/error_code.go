package constants

import "gitlab.yewifi.com/golden-cloud/common/constval2"

/**
 * @Author: prince.lee
 * @Date:   2022/6/15 18:03
 * @Desc:   响应错误码
 */

var (
	// 处理成功:0
	Success                    = constval2.NewInt(0, "成功响应")
	BaseErrParamsInvalid       = constval2.NewInt(210001, "无效参数!")
	BaseErrParamsRequired      = constval2.NewInt(210002, "参数必填!")
	BaseErrDataNull            = constval2.NewInt(210003, "数据为空!")
	BaseErrInsert              = constval2.NewInt(210004, "新增错误!")
	BaseErrDelete              = constval2.NewInt(210005, "删除错误!")
	BaseErrUpdate              = constval2.NewInt(210006, "更新错误!")
	BaseErrSign                = constval2.NewInt(210007, "签名错误!")
	BaseErrFind                = constval2.NewInt(210008, "查询错误!")
	BaseErrGetToken            = constval2.NewInt(210009, "获取access_token失败!")
	BaseErrCheckToken          = constval2.NewInt(210010, "检查access_token失败!")
	BaseErrConfig              = constval2.NewInt(210011, "配置错误!")
	BaseErrDataParse           = constval2.NewInt(210012, "数据解析错误!")
	BaseErrEventPublish        = constval2.NewInt(210013, "事件发布失败!")
	BaseErrEventSubscribe      = constval2.NewInt(210014, "事件订阅失败!")
	BaseErrLimit               = constval2.NewInt(210015, "限流中!")
	BaseErrFieldType           = constval2.NewInt(210016, "字段类型转换错误!")
	BaseErrInstance            = constval2.NewInt(210017, "获取实例失败!")
	BaseErrNoticeSendFail      = constval2.NewInt(210018, "通知发送失败!")
	BaseErrNoticeSmsSendFail   = constval2.NewInt(210019, "短信通知发送失败!")
	BaseErrNoticeEmailSendFail = constval2.NewInt(210020, "邮件通知发送失败!")

	// --- 业务逻辑错误
	// 通用业务逻辑错误
	BusinessLogicErr = constval2.NewInt(221000, "处理失败，请重试!")

	// --- 业务逻辑错误-end
	BusinessLogicErrRiskControl = constval2.NewInt(221001, "触发了风控规则")

	// 异常错误:25xxxx
	PanicErrInternalService = constval2.NewInt(250000, "服务繁忙,请稍后重试")
)
