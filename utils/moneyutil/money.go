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

func FenFloat64ToYuan(fen float64) (yuan string) {
	from := decimal.NewFromFloat(fen)
	toInt := from.Div(decimal.NewFromInt(100))

	return toInt.String()
}

func FenFloat64ToCeilYuan(fen float64) (yuan string) {
	from := decimal.NewFromFloat(fen)
	toInt := from.Div(decimal.NewFromInt(100))

	return toInt.Ceil().String()
}

func FenFloat64ToFloorYuan(fen float64) (yuan string) {
	from := decimal.NewFromFloat(fen)
	toInt := from.Div(decimal.NewFromInt(100))

	return toInt.Floor().String()
}

func FenFloat64ToRoundYuan(fen float64) (yuan string) {
	from := decimal.NewFromFloat(fen)
	// 保留两位小数点，并且四舍五入
	toInt := from.DivRound(decimal.NewFromInt(100), 2)

	return toInt.String()
}

func YuanToJiao(yuan int64) (jiao int64) {
	from := decimal.NewFromInt(yuan)
	toInt := from.Mul(decimal.NewFromInt(10))

	return toInt.IntPart()
}

func YuanToFen(yuan string) (fen int64, err error) {
	from, err := decimal.NewFromString(yuan)
	if err != nil {
		return
	}
	toInt := from.Mul(decimal.NewFromInt(100))

	return toInt.IntPart(), nil
}
