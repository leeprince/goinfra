package stringutil

import (
	"bytes"
	"regexp"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/6/13 23:27
 * @Desc:
 */

// 不支持去除 格式的空格： 代表的是非常规的空格字符，也称为“不间断空格”或“硬空格”。它与普通空格字符（ASCII码为32）不同，它的ASCII码为160。在某些情况下，这种空格字符可能会导致问题，因为它不会被所有的文处理工具和编程语言解释为普通的空格字符。因此，在处理文本时，最好将其替换为普通的空字符。
func ReplaceSpace(s string) string {
	return string(bytes.Replace([]byte(s), []byte(" "), []byte(""), -1))
}

// 匹配所有的空格字符和制表符使用ReplaceAllString函数将其替换为空字符串
var re = regexp.MustCompile(`[\p{Zs}\t]+`)

// 支持去除 格式的空格
func ReplaceWhitespaceChar(s string) string {
	return re.ReplaceAllString(s, "")
}
