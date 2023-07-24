package consts

import "github.com/leeprince/goinfra/consts/constval"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/20 23:38
 * @Desc:
 */

const (
	KeyEsc = "esc"
	KeyF1  = "f1"
	KeyF2  = "f2"
	KeyF3  = "f3"
	KeyF4  = "f4"
	KeyF5  = "f5"
	KeyF6  = "f6"
	KeyF7  = "f7"
	KeyF8  = "f8"
	KeyF9  = "f9"
	KeyF10 = "f10"
	KeyF11 = "f11"
	KeyF12 = "f12"
)

var DarwinOSKeyButtonRawcode = constval.NewStringUint16Group(
	constval.NewStringUint16(KeyEsc, 53),
	constval.NewStringUint16(KeyF1, 122),
	constval.NewStringUint16(KeyF2, 120),
	constval.NewStringUint16(KeyF3, 99),
	constval.NewStringUint16(KeyF4, 118),
	constval.NewStringUint16(KeyF5, 96),
	constval.NewStringUint16(KeyF6, 97),
	constval.NewStringUint16(KeyF7, 98),
	constval.NewStringUint16(KeyF8, 100),
	constval.NewStringUint16(KeyF9, 101),
	constval.NewStringUint16(KeyF10, 109),
	constval.NewStringUint16(KeyF11, 103),
	constval.NewStringUint16(KeyF12, 111),
)

var WindowsOSKeyButtonRawcode = constval.NewStringUint16Group(
	constval.NewStringUint16(KeyEsc, 27),
	constval.NewStringUint16(KeyF1, 112),
	constval.NewStringUint16(KeyF2, 113),
	constval.NewStringUint16(KeyF3, 114),
	constval.NewStringUint16(KeyF4, 115),
	constval.NewStringUint16(KeyF5, 116),
	constval.NewStringUint16(KeyF6, 117),
	constval.NewStringUint16(KeyF7, 118),
	constval.NewStringUint16(KeyF8, 119),
	constval.NewStringUint16(KeyF9, 120),
	constval.NewStringUint16(KeyF10, 121),
	constval.NewStringUint16(KeyF11, 122),
	constval.NewStringUint16(KeyF12, 123),
)
