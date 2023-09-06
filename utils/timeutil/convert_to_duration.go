package timeutil

import "time"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/9/3 02:11
 * @Desc:
 */

// ToMill 毫秒
func ToMill(t int) time.Duration {
	return time.Millisecond * time.Duration(t)
}

// ToSecond 秒
func ToSecond(t int) time.Duration {
	return time.Second * time.Duration(t)
}

// ToMinute 分钟
func ToMinute(t int) time.Duration {
	return time.Minute * time.Duration(t)
}

// ToHour 小时
func ToHour(t int) time.Duration {
	return time.Hour * time.Duration(t)
}
