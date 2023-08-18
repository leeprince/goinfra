package timeutil

import (
	"regexp"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/15 16:59
 * @Desc:
 */

// 是否使用毫秒为单位
//	小于秒的单位时间则使用毫秒
func UseMillisecondUnit(dur time.Duration) bool {
	return dur < time.Second || dur%time.Second != 0
}

func CheckDateFormat(date string) bool {
	// 使用正则表达式匹配年月格式：202308
	pattern := `^\d{4}(0[1-9]|1[0-2])$`
	matched, _ := regexp.MatchString(pattern, date)
	return matched
}
