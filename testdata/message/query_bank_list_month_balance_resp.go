package message

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/26 23:38
 * @Desc:
 */

type QueryBankListMonthBalanceResp struct {
	Code    int64                              `json:"code,omitempty"`    // 状态码。0：成功；非0：失败
	Message string                             `json:"message,omitempty"` // 状态码说明
	LogId   string                             `json:"log_id,omitempty"`  // 日志id
	Data    *QueryBankListMonthBalanceRespData `json:"data,omitempty"`    // 数据
}

type QueryBankListMonthBalanceRespData struct {
	AccountBalanceList []*AccountBalance `json:"account_balance_list,omitempty"` // 账户列表月度余额
}

type AccountBalance struct {
	BankInfo     *BankInfo `json:"bank_info,omitempty"` // 银行信息
	YearMonth_01 string    `json:"year_month_01,omitempty"`
	YearMonth_02 string    `json:"year_month_02,omitempty"`
	YearMonth_03 string    `json:"year_month_03,omitempty"`
}

type BankInfo struct {
	BankAccountName  string `json:"bank_account_name,omitempty"`  // 银行账户名称
	BankName         string `json:"bank_name,omitempty"`          // 银行名称
	BankNumberSuffix string `json:"bank_number_suffix,omitempty"` // 银行卡号(后四位)
}
