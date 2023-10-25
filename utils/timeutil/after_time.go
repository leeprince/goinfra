package timeutil

import "time"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/29 15:31
 * @Desc:
 */

// 当前时间N秒后的时间
func AfterSecond(second int64) time.Time {
	return time.Now().Add(time.Second * time.Duration(second))
}

// 当前时间N分钟后的时间
func AfterMinute(minute time.Duration) time.Time {
	return time.Now().Add(time.Minute * minute)
}

// 当前时间N小时后的时间
func AfterHours(hour time.Duration) time.Time {
	return time.Now().Add(time.Hour * hour)
}

// 当前时间N天后的时间
func AfterDay(day int) time.Time {
	return time.Now().AddDate(0, 0, day)
}

// 当前时间N月后的时间
func AfterMonth(month int) time.Time {
	return time.Now().AddDate(0, month, 0)
}

// 当前时间N年后的时间
func AfterYear(year int) time.Time {
	return time.Now().AddDate(year, 0, 0)
}
