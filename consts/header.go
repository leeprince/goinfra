package consts

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/6/19 10:07
 * @Desc:
 */

var (
	HeaderUberTraceID = "Uber-Trace-Id" // 交给中间件做：如jaeger、opentelemetryc
	HeaderXRealIp     = "X-Real-Ip"     // 代理服务转发
	HeaderXLogID      = "X-Log-Id"      // 自定义链路追踪ID
	HeaderToken       = "Token"         // token. token 或者 access-token 选择其一统一即可
	HeaderAccessToken = "Access-Token"  // access-token. token 或者 access-token 选择其一统一即可
)
