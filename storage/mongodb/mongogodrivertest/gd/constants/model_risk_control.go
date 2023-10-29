package constants

import "gitlab.yewifi.com/golden-cloud/common/constval2"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/10/24 11:55
 * @Desc:
 */

// RiskRuleCode 触发的风控规则编码（A01000：单日出款总额（所有通道汇总金额）；A01001日出款总额度（按通道汇总）；A02001：日出款总笔数（按通道汇总））
/*
编码规则：特征（3位）+维度（3位）
特征：
	A01：出款总额
	A02：出款笔数
维度：
	000：所有通道日维度
	001：指定通道日维度
*/
type RiskRuleCode string

var (
	RiskRuleCodeSingleDayAllChannelTotalAmount = constval2.NewString("A01000", "单日出款总额（所有通道汇总金额）")
	RiskRuleCodeSingleDayOneChannelTotalAmount = constval2.NewString("A01001", "日出款总额度（按通道汇总）")
	RiskRuleCodeSingleDayTotalCount            = constval2.NewString("A02001", "日出款总笔数（按通道汇总）")
)

// 触发风控后的决策，默认都是终止交易（STOP：终止交易）
type RiskResult string

const (
	RiskResultStop RiskResult = "STOP"
)

// 交易渠道（01：银企；02：平安网银；03：微信支付）
type TransactionType string

const (
	TransactionTypeYQ TransactionType = "01"
	TransactionTypePA TransactionType = "02"
	TransactionTypeWX TransactionType = "03"
)

var AllTransactionType = []TransactionType{
	TransactionTypeYQ,
	TransactionTypePA,
	TransactionTypeWX,
}

var TransactionTypeName = map[TransactionType]string{
	TransactionTypeYQ: "银企",
	TransactionTypePA: "平安网银",
	TransactionTypeWX: "微信支付",
}
