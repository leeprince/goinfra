package moneyutil

import "github.com/shopspring/decimal"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/13 23:46
 * @Desc:
 */

func JiaoToYuan(jiao int64) (yuan string) {
	from := decimal.NewFromInt(jiao)
	toInt := from.Div(decimal.NewFromInt(10))

	return toInt.String()
}

func FenToYuan(fen int64) (yuan string) {
	from := decimal.NewFromInt(fen)
	toInt := from.Div(decimal.NewFromInt(100))

	return toInt.String()
}

func YuanToJiao(yuan int64) (jiao string) {
	from := decimal.NewFromInt(yuan)
	toInt := from.Mul(decimal.NewFromInt(10))

	return toInt.String()
}

func YuanToFen(yuan int64) (jiao string) {
	from := decimal.NewFromInt(yuan)
	toInt := from.Mul(decimal.NewFromInt(100))

	return toInt.String()
}
