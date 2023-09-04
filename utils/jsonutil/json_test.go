package jsonutil

import (
	"encoding/json"
	"fmt"
	"github.com/leeprince/goinfra/test/message"
	"github.com/leeprince/goinfra/utils/dumputil"
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
	
	var err error
	// 报错：json: Unmarshal(nil *message.QueryBankListMonthBalanceResp)，因为传值到 json.Unmarshal 的第二个参数时是空指针
	var respError *message.QueryBankListMonthBalanceResp
	fmt.Printf("respError-------------------\n")
	{
		err = json.Unmarshal([]byte(jsonStr), respError)
		if err != nil {
			fmt.Println("json.Unmarshal respError Error:", err)
		}
		fmt.Println(">>>")
		
		err = JsoniterCompatible.Unmarshal([]byte(jsonStr), respError)
		if err != nil {
			fmt.Println("JsoniterCompatible.Unmarshal respError Error:", err)
		}
	}
	
	// 正确，因为传值到 json.Unmarshal 的第二个参数时不会是空指针
	resp1 := &message.QueryBankListMonthBalanceResp{}
	fmt.Printf("resp1-------------------\n")
	{
		err = json.Unmarshal([]byte(jsonStr), resp1)
		if err != nil {
			fmt.Println("json.Unmarshal resp1 Error:", err)
			return
		}
		fmt.Printf("json.Unmarshal resp1: %+v\n", resp1)
		fmt.Println(">>>")
		
		err = JsoniterCompatible.Unmarshal([]byte(jsonStr), resp1)
		if err != nil {
			fmt.Println("JsoniterCompatible.Unmarshal resp1 Error:", err)
			return
		}
		fmt.Printf("JsoniterCompatible.Unmarshal resp1: %+v\n", resp1)
		
		// Now you can work with the populated QueryBankListMonthBalanceResp structure
		fmt.Printf("Code: %d\n", resp1.Code)
		fmt.Printf("Message: %s\n", resp1.Message)
		fmt.Printf("LogId: %s\n", resp1.LogId)
		for _, balance := range resp1.Data.AccountBalanceList {
			fmt.Println("Bank Account Name:", balance.BankInfo.BankAccountName)
			fmt.Println("Bank Name:", balance.BankInfo.BankName)
			fmt.Println("Bank Number Suffix:", balance.BankInfo.BankNumberSuffix)
			fmt.Println("Year Month 01:", balance.YearMonth_01)
			fmt.Println("Year Month 02:", balance.YearMonth_02)
			fmt.Println("Year Month 03:", balance.YearMonth_03)
			fmt.Println("------------")
		}
		
		fmt.Println("+++++++++++++++++++++++++++++++")
		dumputil.Println(resp1)
	}
	
	// 正确，因为传值到 json.Unmarshal 的第二个参数时不会是空指针
	var resp2 message.QueryBankListMonthBalanceResp
	fmt.Printf("resp2-------------------\n")
	{
		err = json.Unmarshal([]byte(jsonStr), &resp2)
		if err != nil {
			fmt.Println("json.Unmarshal resp2 Error:", err)
			return
		}
		fmt.Printf("json.Unmarshal resp2: %+v\n", resp2)
		fmt.Println(">>>")
		
		err = JsoniterCompatible.Unmarshal([]byte(jsonStr), &resp2)
		if err != nil {
			fmt.Println("JsoniterCompatible.Unmarshal resp2 Error:", err)
			return
		}
		fmt.Printf("JsoniterCompatible.Unmarshal resp2: %+v\n", resp2)
		
		// Now you can work with the populated QueryBankListMonthBalanceResp structure
		fmt.Printf("Code: %d\n", resp2.Code)
		fmt.Printf("Message: %s\n", resp2.Message)
		fmt.Printf("LogId: %s\n", resp2.LogId)
		for _, balance := range resp2.Data.AccountBalanceList {
			fmt.Println("Bank Account Name:", balance.BankInfo.BankAccountName)
			fmt.Println("Bank Name:", balance.BankInfo.BankName)
			fmt.Println("Bank Number Suffix:", balance.BankInfo.BankNumberSuffix)
			fmt.Println("Year Month 01:", balance.YearMonth_01)
			fmt.Println("Year Month 02:", balance.YearMonth_02)
			fmt.Println("Year Month 03:", balance.YearMonth_03)
			fmt.Println("------------")
		}
		
		fmt.Println("+++++++++++++++++++++++++++++++")
		dumputil.Println(resp2)
	}
	
}
