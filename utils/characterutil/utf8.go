package characterutil

import "unicode/utf8"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/6/8 17:26
 * @Desc:
 */

func IsUTF8(v string) bool {
	return utf8.ValidString(v)
}
