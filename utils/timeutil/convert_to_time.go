package timeutil

import (
	"github.com/leeprince/goinfra/consts"
	"github.com/spf13/cast"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/15 17:44
 * @Desc:
 */

// timeLayout参考：consts的time.go
func ToTime(timeStr, timeLayout string) (t time.Time, err error) {
	loc := time.Local
	/*
		//loc, err := time.LoadLocation("Local")
		//loc, err := time.LoadLocation("Asia/Shanghai")
		if err != nil {
			return
		}
	*/

	t, err = time.ParseInLocation(timeLayout, timeStr, loc)
	if err != nil {
		return
	}

	return
}

func ToTimeUnix(timeStr, timeLayout string) (timeUnix int64, err error) {
	t, err := ToTime(timeStr, timeLayout)
	if err != nil {
		return
	}

	timeUnix = t.Unix()
	return
}

// 将时间戳转换为 time.Time 类型
func ToTimeByUnix(unixTime int64) (t time.Time) {
	return time.Unix(unixTime, 0)
}

// 计算当前月1日的时间
func GetMonthFirthDayTime() time.Time {
	// 获取当前时间
	currentTime := time.Now()

	return time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0, currentTime.Location())
}

// 计算当前年1月的时间
func GetYearFirthMonthTime() time.Time {
	// 获取当前时间
	currentTime := time.Now()

	return time.Date(currentTime.Year(), 1, 1, 0, 0, 0, 0, currentTime.Location())
}

// 计算当前年1月~当前月
func GetCurrentYearMonthList() []int64 {
	// 当前年1月
	firstMonth := cast.ToInt64(GetYearFirthMonthTime().Format(consts.TimeLayoutV71))
	// 当前月
	endMonth := cast.ToInt64(GetMonthFirthDayTime().Format(consts.TimeLayoutV71))

	var monthList []int64
	for firstMonth <= endMonth {
		monthList = append(monthList, firstMonth)
		firstMonth++
	}
	return monthList
}

// 获取当前时间的上一个月开始时间&结束时间
func GetPreMonthUnixTimeRange() (startTime, endTime time.Time) {
	// 获取当前时间
	currentTime := time.Now()

	startTime = time.Date(currentTime.Year(), currentTime.Month()-1, 1, 0, 0, 0, 0, currentTime.Location())

	currentMonthTime := time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0, currentTime.Location())
	endTime = currentMonthTime.Add(-time.Second * 1)
	return
}
