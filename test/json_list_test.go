package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/18 15:02
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

func TestJsonList(t *testing.T) {
	jsonStr := `{
	  "code": 52564,
	  "message": "ok",
	  "log_id": "prince-001",
	  "data": {
		"account_balance_list": [
		  {
			"bank_info": {
			  "bank_account_name": "prince-中国银行-01",
			  "bank_name": "中国银行",
			  "bank_number_suffix": "0001"
			},
			"year_month_01": "100.01",
			"year_month_02": "101.02",
			"year_month_03": "102.03"
		  },
		  {
			"bank_info": {
			  "bank_account_name": "prince-中国银行-02",
			  "bank_name": "中国银行",
			  "bank_number_suffix": "0004"
			},
			"year_month_01": "200.01",
			"year_month_02": "201.02",
			"year_month_03": "202.03"
		  }
		]
	  }
	}`

	var resp QueryBankListMonthBalanceResp
	err := json.Unmarshal([]byte(jsonStr), &resp)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Now you can work with the populated QueryBankListMonthBalanceResp structure
	fmt.Printf("Code: %d\n", resp.Code)
	fmt.Printf("Message: %s\n", resp.Message)
	fmt.Printf("LogID: %s\n", resp.LogId)
	for _, balance := range resp.Data.AccountBalanceList {
		fmt.Println("Bank Account Name:", balance.BankInfo.BankAccountName)
		fmt.Println("Bank Name:", balance.BankInfo.BankName)
		fmt.Println("Bank Number Suffix:", balance.BankInfo.BankNumberSuffix)
		fmt.Println("Year Month 01:", balance.YearMonth_01)
		fmt.Println("Year Month 02:", balance.YearMonth_02)
		fmt.Println("Year Month 03:", balance.YearMonth_03)
		fmt.Println("------------")
	}
}
