package timeutil

import "time"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/28 18:47
 * @Desc:
 */

// 计算过去的时间
func GetBeforeData(day int) time.Time {
	// 获取当前时间
	currentTime := time.Now()

	beforDay := day * -1
	return currentTime.AddDate(0, 0, beforDay)
}

// 计算当前月1日的时间
func GetMonthFirthDay() time.Time {
	// 获取当前时间
	currentTime := time.Now()

	return time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0, currentTime.Location())
}
