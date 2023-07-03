package characterutil

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/6/20 00:47
 * @Desc:
 */
/*
在 ASCII 编码中，
- 字符 '0' 到 '9' 的十进制 ASCII 码值分别是 48 到 57
- 字符 'A' 到 'Z' 的十进制 ASCII 码值分别是 65 到 90字符
- 'a' 到 'z' 的十进制 ASCII 码值分别是 97 到 122

而中文字符，并不属于 ASCII 字符集。如字符 '日' 和 '月' 并不属于 ASCII 字符集，它们的编码方式取决于具体的字符编码标准。
在 Unicode 编码中，字符 '日' 和 '月' 的编码分别是 U+65E5 和 U+6708。在十进制表示中，它们的值分别是 26085 和 26376。需要注意的是，这里的十进制表示并不是 ASCII 码值，而是 Unicode 码值。
*/
/*
字符 '日' 和 '月' 并不属于 ASCII 字符集，因此无法使用 ASCII 编码表示。下面是字符 '日' 和 '月' 在其他编码方式下的表示：
- Unicode 编码：字符 '日' 和 '月' Unicode 编码分别是 U+65E5 和 U+6708。在十六进制表示中，它们的值分别是 0x65E5 和 0x6708。在十进制表示中，它们的值分别是 26085 和 26376。
- UTF-8 编码：UTF-8 是一种变长编码方式，可以将 Unicode 字符编码为 1 到 4 个字节。字符 '日' 和 '月' 在 UTF-8 编码中的表示分别是 0x06 0x97 0xA5 和 0xE6 0x9C 0x88。这里的每个字节都是十六进制表示，对应的十进制值分别是 230、151、165 和 230、156、136。
- GBK 编码：GBK 是一种中文编码方式，它将汉字和 ASCII 字符编码为 1 或 2 个字节。字符 '日' 和 '月' 在 GBK 编码中的表示分别是 0xD0 0xA1 和 0xB0 0xE0。这里的每个字节都是十六进制表示，对应的十进制值分别是 208、161 和 176、224。

需要注意的是，不同的编码方式使用不同的编码规则，因此同一个字符在不同的编码方式下可能会有不同的表示方式。
*/

// 将 ASCII 字符转为 ASCII 十进制和 ASCII 十六进制
func ASCIICharToASCIIIntAndASCIIHex(char int8) (intStr, hexStr string) {
	// ASCII 十进制
	intStr = fmt.Sprintf("%d", char)
	// ASCII 十六进制
	hexStr = fmt.Sprintf("%x", char)
	return
}

// 将 ASCII 字符的字符串转为 ASCII 十进制的字符串和 ASCII 十六进制的字符串
// 注意：如果传入的charStr包含非`ASCII 字符`，如：中文字符，则可以转成相应的 hexStr，但是该 hexStr 通过 `ASCIIHexStrToASCIIChar` 并不能转成相应的中文字符的字符串
func ASCIICharStrToASCIIIntAndASCIIHex(charStr string) (intStr, hexStr string, err error) {
	var intStrBuf bytes.Buffer
	var hexStrBuf bytes.Buffer
	for _, char := range charStr {
		// ASCII 十进制
		_, err = intStrBuf.WriteString(fmt.Sprintf("%d", char))
		if err != nil {
			err = errors.New("ASCII 十进制发生错误：" + err.Error())
			return
		}
		// ASCII 十六进制
		_, err = hexStrBuf.WriteString(fmt.Sprintf("%x", char))
		if err != nil {
			err = errors.New("ASCII 十六进制发生错误：" + err.Error())
			return
		}
	}
	intStr = intStrBuf.String()
	hexStr = hexStrBuf.String()
	return
}

// 将ASCII十进制整数（这是十进制整数是ASCII 码值的十进制表示；当然还有十六进制表示方式）转换为对应的 ASCII 字符
// fmt.Sprintf 函数的第一个参数是格式化字符串，其中 %c 表示要将参数转换为对应的 ASCII 字符。函数的值是一个字符串，表示转换后的字符。
// 需要注意的是: ASCII 字符集只包含 128 个字符，对应的整数范围是 0 到 127。如果要将大于 127 的整数转换为对应的字符，可能会得到一个非 ASCII 字符。
func ASCIIIntToASCIIChar(i int8) string {
	return fmt.Sprintf("%c", i)
}

// 将ASCII十进制整数的数组转成 ASCII 字符的字符串
func ASCIIIntArrToASCIIChars(ii []int8) string {
	var s bytes.Buffer
	for _, i := range ii {
		s.WriteString(fmt.Sprintf("%c", i))
	}
	return s.String()
}

// 将ASCII十进制整数的字符串转成 ASCII 字符的字符串
// 	ASCII十进制整数的字符串必须是由ASCII十进制整数拼接在一起的字符串
func ASCIIIntStrToASCIIChars(ii string) (string, error) {
	if len(ii)%2 != 0 {
		return "", errors.New("ASCII十进制整数的字符串的长度必须是偶数")
	}
	
	var s bytes.Buffer
	for i := 0; i < len(ii); i += 2 {
		numStr := ii[i : i+2]
		asciiInt, err := strconv.Atoi(numStr)
		if err != nil {
			return "", errors.New("拆分ASCII十进制整数的字符串发生错误：" + err.Error())
		}
		// 检查十进制整数；不检查则让`fmt.Sprintf("%c", asciiInt)`去检查
		// if num > 128 {
		// 	return "", errors.New("拆分ASCII十进制整数大于 127")
		// }
		
		s.WriteString(fmt.Sprintf("%c", asciiInt))
	}
	
	return s.String(), nil
}

// 将ASCII十六进制的字符串转成 ASCII 字符的字符串
func ASCIIHexStrToASCIIChar(s string) (string, error) {
	charByte, err := hex.DecodeString(s)
	if err != nil {
		return "", errors.New("十六进制转ASCII字符的字符串错误：" + err.Error())
	}
	return string(charByte), nil
}
