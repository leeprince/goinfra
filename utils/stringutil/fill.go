package stringutil

import (
	"fmt"
	"strings"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/15 13:30
 * @Desc:
 */

// IsFillLeft: 是否填充在左边
func FillChar(sourceStr string, fillChar rune, fillLen int, IsFillLeft bool) string {
	sourceStrLen := len(sourceStr)
	if fillLen <= sourceStrLen {
		return sourceStr
	}
	
	// 需要填充的个数
	fillCharNum := fillLen - sourceStrLen
	
	// 按填充个数填充字符
	fillChars := strings.Repeat(string(fillChar), fillCharNum)
	
	if IsFillLeft {
		return fmt.Sprintf("%s%s", fillChars, sourceStr)
	}
	return fmt.Sprintf("%s%s", sourceStr, fillChars)
}

func FillCharLeft(sourceStr string, fillChar rune, fillLen int) string {
	return FillChar(sourceStr, fillChar, fillLen, true)
}

func FillCharRight(sourceStr string, fillChar rune, fillLen int) string {
	return FillChar(sourceStr, fillChar, fillLen, false)
}
