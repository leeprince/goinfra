package model

import (
	"gitlab.yewifi.com/golden-cloud/transaction-risk-control/pkg/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/10/20 16:05
 * @Desc:
 */

// 查询模型：包含默认主键_id; 用于查询时能够返回默认主键_id
type QueryRiskControl struct {
	Id          primitive.ObjectID `bson:"_id" json:"_id"` // 租户ID
	RiskControl `bson:",inline"`
}

// 插入模型：不能包含默认主键_id（mongDB自动插入）,否则不传值的情况会为空导致第二条记录无法插入
type RiskControl struct {
	OpenOrgId           string                    `bson:"open_org_id" json:"open_org_id"`                       // 开放租户ID（付款方）
	TransactionType     constants.TransactionType `bson:"transaction_type" json:"transaction_type"`             // 交易渠道（01：银企；02：平安网银；03：微信支付）
	TransactionSn       string                    `bson:"transaction_sn" json:"transaction_sn"`                 // 交易流水号
	LogId               string                    `bson:"log_id" json:"log_id"`                                 // 交易流水号
	RiskResult          constants.RiskResult      `bson:"risk_result" json:"risk_result"`                       // 触发风控后的决策，默认都是终止交易（STOP：终止交易）
	RiskRuleCode        constants.RiskRuleCode    `bson:"risk_rule_code" json:"risk_rule_code"`                 // 触发的风控规则编码（A01000：单日出款总额（所有通道汇总金额）；A01001日出款总额度（按通道汇总）；A02001：日出款总笔数（按通道汇总））
	RiskRuleCodeMessage string                    `bson:"risk_rule_code_message" json:"risk_rule_code_message"` // 触发的风控规则编码说明
	IsNotice            bool                      `bson:"is_notice" json:"is_notice"`                           // 是否发送告警通知
	NoticeFailReason    string                    `bson:"notice_fail_reason" json:"notice_fail_reason"`         // 发送告警通知失败的原因
	UpdatedAt           int64                     `bson:"updated_at" json:"updated_at"`                         // 更新时间
	CreatedAt           int64                     `bson:"created_at" json:"created_at"`                         // 日志创建时间
}

var RiskControlField = struct {
	OpenOrgId           string
	TransactionType     string
	TransactionSn       string
	LogId               string
	RiskResult          string
	RiskRuleCode        string
	RiskRuleCodeMessage string
	IsNotice            string
	NoticeFailReason    string
	UpdatedAt           string
	CreatedAt           string
}{
	OpenOrgId:           "open_org_id",
	TransactionType:     "transaction_type",
	TransactionSn:       "transaction_sn",
	LogId:               "log_id",
	RiskResult:          "risk_result",
	RiskRuleCode:        "risk_rule_code",
	RiskRuleCodeMessage: "risk_rule_code_message",
	IsNotice:            "is_notice",
	NoticeFailReason:    "notice_fail_reason",
	UpdatedAt:           "updated_at",
	CreatedAt:           "created_at",
}

func (RiskControl) DataBaseName() string {
	return "gdc_transaction_risk_control"
}

func (RiskControl) CollectionName() string {
	return "risk_control"
}
