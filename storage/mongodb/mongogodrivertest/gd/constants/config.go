package constants

/**
 * @Author: prince.lee
 * @Date:   2022/6/15 9:55
 * @Desc:   与配置文件conf.yaml 相关的配置项
 */

// ------------------ DB
const ()

// ------------------ DB - end

// ------------------ MQ
// 配置rabbitmq对应的key
const ()

// ------------------ MQ - end

// TplSmsRiskWarning 配置短信模板
const (
	TplSmsRiskWarning      = 1961460                                         // 风险告警短信模板
	TplEmailRiskWarning    = "<p>尊敬的用户您好！</p><p>%s资金池触发通道限额，已终止出款，请尽快查看！<p>" // 风险告警邮件模板
	EmailSenderRiskWarning = "notice@info.goldentec.com"                     // 风险告警邮件发送方地址
)
