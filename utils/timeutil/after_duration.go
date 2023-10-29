package timeutil

import (
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/10/23 14:50
 * @Desc:
 */

func AfterSecondD(second int64) (d time.Duration, err error) {
	// 方式1：解析字符串，不过支持的单位有限：Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	//return time.ParseDuration(fmt.Sprintf("%ds", second))

	// 方式2：直接的方式【推荐】
	return time.Second * time.Duration(second), nil
}

func AfterMinuteD(minute int64) (d time.Duration, err error) {
	// 方式1：解析字符串，不过支持的单位有限：Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	//return time.ParseDuration(fmt.Sprintf("%dm", minute))

	// 方式2：直接的方式【推荐】
	return time.Minute * time.Duration(minute), nil
}

func AfterHoursD(hours int64) (d time.Duration, err error) {
	// 方式1：解析字符串，不过支持的单位有限：Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	//return time.ParseDuration(fmt.Sprintf("%dh", hours))

	// 方式2：直接的方式【推荐】
	return time.Hour * time.Duration(hours), nil
}

func AfterDayD(day int64) (d time.Duration, err error) {
	// 方式1：解析字符串，不过支持的单位有限：Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	//return time.ParseDuration(fmt.Sprintf("%dh", day*24))

	// 方式2：直接的方式【推荐】
	return time.Hour * 24 * time.Duration(day), nil
}
