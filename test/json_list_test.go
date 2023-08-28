package test

import (
	"encoding/json"
	"fmt"
	"github.com/leeprince/goinfra/test/message"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/18 15:02
 * @Desc:
 */

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
	
	var resp message.QueryBankListMonthBalanceResp
	err := json.Unmarshal([]byte(jsonStr), &resp)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	
	// Now you can work with the populated QueryBankListMonthBalanceResp structure
	fmt.Printf("Code: %d\n", resp.Code)
	fmt.Printf("Message: %s\n", resp.Message)
	fmt.Printf("LogId: %s\n", resp.LogId)
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
