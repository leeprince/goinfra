package message

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/29 22:55
 * @Desc:
 */

type VerifyAuthReq struct {
	VerifyCode string `json:"verify_code"`
}

type VerifyAuthResp struct {
	Code    int32              `json:"code"`    // 状态码，0：成功；非 0 失败
	Message string             `json:"message"` // 状态码说明
	LogID   string             `json:"log_id"`  // 日志 ID/链路追踪 ID
	Data    VerifyAuthRespData `json:"data"`    // 数据
}
type VerifyAuthRespData struct {
	ConfirmVerifyCode string `json:"confirm_verify_code"`
}
