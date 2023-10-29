package model

import (
	"gitlab.yewifi.com/golden-cloud/transaction-risk-control/pkg/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 查询模型：包含默认主键_id; 用于查询时能够返回默认主键_id
type QueryTransactionRecord struct {
	Id                primitive.ObjectID `bson:"_id" json:"_id"` // 租户ID
	TransactionRecord `bson:",inline"`
}

// 插入模型：不能包含默认主键_id（mongDB自动插入）,否则不传值的情况会为空导致第二条记录无法插入
type TransactionRecord struct {
	OpenOrgId         string                      `bson:"open_org_id" json:"open_org_id"`               // 开放租户ID（付款方）
	TransactionType   constants.TransactionType   `bson:"transaction_type" json:"transaction_type"`     // 交易渠道（01：银企；02：平安网银；03：微信支付）
	TransactionSn     string                      `bson:"transaction_sn" json:"transaction_sn"`         // 交易流水号
	TransactionAmount int64                       `bson:"transaction_amount" json:"transaction_amount"` // 交易金额（单位：分）
	TransactionStatus constants.TransactionStatus `bson:"transaction_status" json:"transaction_status"` // 交易状态（Success：付款成功；Ing：付款中；Fail：付款失败）
	PayerAccountNo    string                      `bson:"payer_account_no" json:"payer_account_no"`     // 付款银行账号
	PayerAccountName  string                      `bson:"payer_account_name" json:"payer_account_name"` // 付款账户名
	PayeeAccountNo    string                      `bson:"payee_account_no" json:"payee_account_no"`     // 收款银行账号
	PayeeAccountName  string                      `bson:"payee_account_name" json:"payee_account_name"` // 收款银行账号
	UpdatedAt         int64                       `bson:"updated_at" json:"updated_at"`                 // 更新时间
	CreatedAt         int64                       `bson:"created_at" json:"created_at"`                 // 创建时间
}

var TransactionRecordField = struct {
	OpenOrgId         string
	TransactionType   string
	TransactionSn     string
	TransactionAmount string
	TransactionStatus string
	PayerAccountNo    string
	PayerAccountName  string
	PayeeAccountNo    string
	PayeeAccountName  string
	CreatedAt         string
}{
	OpenOrgId:         "open_org_id",
	TransactionType:   "transaction_type",
	TransactionSn:     "transaction_sn",
	TransactionAmount: "transaction_amount",
	TransactionStatus: "transaction_status",
	PayerAccountNo:    "payer_account_no",
	PayerAccountName:  "payer_account_name",
	PayeeAccountNo:    "payee_account_no",
	PayeeAccountName:  "payee_account_name",
	CreatedAt:         "created_at",
}

func (TransactionRecord) DataBaseName() string {
	return "gdc_transaction_risk_control"
}

func (TransactionRecord) CollectionName() string {
	return "transaction_record"
}
