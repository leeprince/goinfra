package characterutil

import "unicode"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/6/8 17:21
 * @Desc:
 */

func IsChinese(v []rune) bool {
	for _, r := range v {
		if unicode.Is(unicode.Han, r) {
			return true
		}
	}

	return false
}
