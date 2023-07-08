package ginutil

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/6 10:19
 * @Desc:
 */

type BaseResponse struct {
	BaseCommonResponse
	Data any `json:"data"` // 响应数据
}

type BaseCommonResponse struct {
	Code    int32  `json:"code"`    // 状态码，0：成功；非 0 失败
	Message string `json:"message"` // 状态码说明
	LogID   string `json:"log_id"`  // 日志 ID/链路追踪 ID
}
