package dumputil

import (
	"fmt"
	"strings"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/6/10 17:27
 * @Desc:	存在递归遍历时更清晰的输入格式
 */

func ForIndentSPrintf(forNum int64, indentChar string, format string, a ...any) string {
	if indentChar == "" {
		indentChar = " "
	}
	indent := strings.Repeat(indentChar, int(forNum)*4)
	format = indent + format
	return fmt.Sprintf(format, a...)
}

func ForIndentPrintf(forNum int64, indentChar string, format string, a ...any) {
	if indentChar == "" {
		indentChar = " "
	}
	indent := strings.Repeat(indentChar, int(forNum)*4)
	format = indent + format
	fmt.Printf(format, a...)
}

func ForIndentPrintln(forNum int64, indentChar, dataStr string) {
	if indentChar == "" {
		indentChar = " "
	}
	indent := strings.Repeat(indentChar, int(forNum)*4)
	fmt.Println(indent, dataStr)
}
