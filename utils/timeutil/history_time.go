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
