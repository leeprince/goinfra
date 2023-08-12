package consts

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/12 12:49
 * @Desc:
 */

type StringFormat string

const (
	StringFormatNone   StringFormat = "None"
	StringFormatBase64 StringFormat = "Base64"
	StringFormatHex    StringFormat = "Hex"
)
