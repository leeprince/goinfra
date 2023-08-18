package test

import (
	"fmt"
	"reflect"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/18 17:24
 * @Desc:
 */

type AccountMonthBalance struct {
	YearMonth_01 string `json:"year_month_01"` // 01月份（余额），单位元
	YearMonth_02 string `json:"year_month_02"` // 02月份（余额），单位元
	YearMonth_03 string `json:"year_month_03"` // 03月份（余额），单位元
	YearMonth_04 string `json:"year_month_04"` // 04月份（余额），单位元
	YearMonth_05 string `json:"year_month_05"` // 05月份（余额），单位元
	YearMonth_06 string `json:"year_month_06"` // 06月份（余额），单位元
	YearMonth_07 string `json:"year_month_07"` // 07月份（余额），单位元
	YearMonth_08 string `json:"year_month_08"` // 08月份（余额），单位元
	YearMonth_09 string `json:"year_month_09"` // 09月份（余额），单位元
	YearMonth_10 string `json:"year_month_10"` // 10月份（余额），单位元
	YearMonth_11 string `json:"year_month_11"` // 11月份（余额），单位元
	YearMonth_12 string `json:"year_month_12"` // 12月份（余额），单位元
}

func TestDynamicFieldSetValue(t *testing.T) {
	month := 2
	balance := "100.00"
	monthDynamicField := fmt.Sprintf("YearMonth_%02d", month)

	accountMonthBalance := AccountMonthBalance{}
	// 拆解
	valueOf := reflect.ValueOf(&accountMonthBalance) // 注意这里必须使用变量的指针地址
	fmt.Printf("--- valueof:%+v \n", valueOf)
	valueOfElem := valueOf.Elem()
	fmt.Printf("--- valueOfElem:%+v \n", valueOfElem)
	valueOfElemFieldByName := valueOfElem.FieldByName(monthDynamicField)
	fmt.Printf("--- valueOfElemFieldByName:%+v \n", valueOfElemFieldByName)
	valueOfElemFieldByName.SetString(balance)
	/*
		//valueOfElemFieldByName := reflect.ValueOf(&accountMonthBalance).Elem().FieldByName(monthDynamicField)
		//valueOfElemFieldByName.SetString(balance)
	*/

	fmt.Printf("--- valueOfElemFieldByName:%+v \n", accountMonthBalance)

}
