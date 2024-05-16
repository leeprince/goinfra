package timeutil

import (
	"github.com/leeprince/goinfra/consts"
	"github.com/spf13/cast"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/28 18:47
 * @Desc:
 */

// 计算过去的时间:日
func HistoryData(day int) time.Time {
	// 获取当前时间
	currentTime := time.Now()

	beforDay := day * -1
	return currentTime.AddDate(0, 0, beforDay)
}

// HistoryDataByTime 计算指定时间过去的时间:日
func HistoryDataByTime(useTime time.Time, day int) time.Time {
	beforDay := day * -1
	return useTime.AddDate(0, 0, beforDay)
}

// 计算过去的时间:月
func HistoryMonth(month int) time.Time {
	// 获取当前时间
	currentTime := time.Now()

	beforMonth := month * -1
	return currentTime.AddDate(0, beforMonth, 0)
}

// 计算指定时间过去的时间:月
func HistoryMonthByTime(useTime time.Time, month int) time.Time {
	beforMonth := month * -1
	return useTime.AddDate(0, beforMonth, 0)
}

// 包含请求参数的date的月份
// 	- date:202308 的时间格式
func HistoryMonthListByMonth(date string, oldMonth int) (monthList []int64, err error) {
	t, err := ToTime(date, consts.TimeYYmm)
	if err != nil {
		return
	}

	for i := 0; i <= oldMonth; i++ {
		historyMonthTime := HistoryMonthByTime(t, i)
		month := cast.ToInt64(historyMonthTime.Format("200601"))
		monthList = append(monthList, month)
	}

	return
}

// ConvertOneDateStringToTimestamps 将一个日期字符串（格式为"20240901"）转换为日期对应的开始时间和结束时间戳
func ConvertOneDateStringToTimestamps(dateString string) (startTime, endTime int64) {
	// 解析日期字符串为time.Time类型
	dateFormat := "20060102"
	parsedDate, _ := time.Parse(dateFormat, dateString)

	// 获取当天的开始时间戳（即00:00:00）
	startTime = parsedDate.Unix()

	// 计算结束时间戳（即第二天的00:00:00减一秒）
	endOfDay := parsedDate.AddDate(0, 0, 1)
	endTime = endOfDay.Add(-time.Second).Unix()

	return
}

// ConvertDateStringToTimestamps 计算指定日期的开始时间和结束时间戳 将开始/结束日期字符串（格式为"20240901"）转换为日期对应的开始时间和结束时间戳
func ConvertDateStringToTimestamps(dateStartString, dateEndString string) (startTime, endTime int64) {
	// 解析日期字符串为time.Time类型
	dateFormat := "20060102"
	dateStartData, _ := time.Parse(dateFormat, dateStartString)
	dateEndData, _ := time.Parse(dateFormat, dateEndString)

	// 获取当天的开始时间戳（即00:00:00）
	startTime = dateStartData.Unix()

	// 计算结束时间戳（即第二天的00:00:00减一秒）
	endOfDay := dateEndData.AddDate(0, 0, 1)
	endTime = endOfDay.Add(-time.Second).Unix()

	return
}

// 计算指定时间过去的时间:月
func HistoryMonthByUseMonth(useYear, useMonth, month int) time.Time {
	useTime := time.Date(useYear, time.Month(useMonth), 1, 0, 0, 0, 0, time.Local)
	// 需要填充月份为2位
	//useTime, _ := ToTime(fmt.Sprintf("%d%02d", useYear, useMonth), consts.TimeYYmm)

	return HistoryMonthByTime(useTime, month)
}

// 计算过去的时间:年
func HistoryYear(year int) time.Time {
	// 获取当前时间
	currentTime := time.Now()

	beforeYear := year * -1
	return currentTime.AddDate(beforeYear, 0, 0)
}

// 计算指定时间过去的时间:月
func HistoryYearByTime(useTime time.Time, year int) time.Time {
	beforeYear := year * -1
	return useTime.AddDate(beforeYear, 0, 0)
}
