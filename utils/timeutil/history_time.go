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
func GetHistoryData(day int) time.Time {
	// 获取当前时间
	currentTime := time.Now()

	beforDay := day * -1
	return currentTime.AddDate(0, 0, beforDay)
}

// 计算指定时间过去的时间:日
func GetHistoryDataByTime(useTime time.Time, day int) time.Time {
	beforDay := day * -1
	return useTime.AddDate(0, 0, beforDay)
}

// 计算过去的时间:月
func GetHistoryMonth(month int) time.Time {
	// 获取当前时间
	currentTime := time.Now()

	beforMonth := month * -1
	return currentTime.AddDate(0, beforMonth, 0)
}

// 计算指定时间过去的时间:月
func GetHistoryMonthByTime(useTime time.Time, month int) time.Time {
	beforMonth := month * -1
	return useTime.AddDate(0, beforMonth, 0)
}

// 包含请求参数的date的月份
// 	- date:202308 的时间格式
func GetHistoryMonthListByMonth(date string, oldMonth int) (monthList []int64, err error) {
	t, err := ToTime(date, consts.TimeLayoutV71)
	if err != nil {
		return
	}

	for i := 0; i <= oldMonth; i++ {
		historyMonthTime := GetHistoryMonthByTime(t, i)
		month := cast.ToInt64(historyMonthTime.Format("200601"))
		monthList = append(monthList, month)
	}

	return
}

// 计算指定时间过去的时间:月
func GetHistoryMonthByUseMonth(useYear, useMonth, month int) time.Time {
	useTime := time.Date(useYear, time.Month(useMonth), 1, 0, 0, 0, 0, time.Local)
	// 需要填充月份为2位
	//useTime, _ := ToTime(fmt.Sprintf("%d%02d", useYear, useMonth), consts.TimeLayoutV71)

	return GetHistoryMonthByTime(useTime, month)
}

// 计算过去的时间:年
func GetHistoryYear(year int) time.Time {
	// 获取当前时间
	currentTime := time.Now()

	beforeYear := year * -1
	return currentTime.AddDate(beforeYear, 0, 0)
}

// 计算指定时间过去的时间:月
func GetHistoryYearByTime(useTime time.Time, year int) time.Time {
	beforeYear := year * -1
	return useTime.AddDate(beforeYear, 0, 0)
}
