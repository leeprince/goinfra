package moneyutil

import "github.com/shopspring/decimal"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/13 23:46
 * @Desc:
 */

func JiaoToYuan(jiao int64) (yuan string) {
	jiaoD := decimal.NewFromInt(jiao)
	yuanD := jiaoD.Div(decimal.NewFromInt(10))
	
	return yuanD.String()
}

func YuanToJiao(yuan int64) (jiao string) {
	yuanD := decimal.NewFromInt(yuan)
	jiaoD := yuanD.Mul(decimal.NewFromInt(10))
	
	return jiaoD.String()
}