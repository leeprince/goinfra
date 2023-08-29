package timeutil

import (
	"github.com/leeprince/goinfra/consts"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/5/20 11:49
 * @Desc:
 */

/*
const (
	TimeYYmdHis  = "2006-01-02 15:04:05"
	TimeLayoutV11 = "2006-01-02 15:04:05.999999999 -0700 MST"
	TimeFYmdHisNano = "2006-01-02 15:04:05.999999999" // 机器原因，基本值都等于"2006-01-02 15:04:05.999999"
	TimeFYmdHisMicro = "2006-01-02 15:04:05.999999"
	TimeFYmdHisMill = "2006-01-02 15:04:05.999"
	TimeYmdHis = "06-01-02 15:04:05"
	TimeYYmdHi  = "2006-01-02 15:04"
	TimeYYmdH  = "2006-01-02 15"
	TimeYYmd  = "2006-01-02"
	TimeYymd = "2006/01/02"
	TimeMdyy  = "01-02-2006"
	TimeMdy  = "01-02-06"
	TimeYYm  = "2006-01"
	TimeYYmm = "200601"
	TimeYY  = "2006"
)
*/

// DataTime "2006-01-02 15:04:05"
func DataTime() string {
	return time.Now().Format(consts.TimeYYmdHis)
}

// DataTimeT "2006-01-02 15:04:05"
func DataTimeT(t time.Time) string {
	return t.Format(consts.TimeYYmdHis)
}

// DataTimeMicrosecond "2006-01-02 15:04:05.999999"
func DataTimeMicrosecond() string {
	return time.Now().Format(consts.TimeFYmdHisMicro)
}

// DataTimeMillisecond "2006-01-02 15:04:05.999"
func DataTimeMillisecond() string {
	return time.Now().Format(consts.TimeFYmdHisMill)
}

// DataTimeDataSecond "06-01-02 15:04:05"
func DataTimeDataSecond() string {
	return time.Now().Format(consts.TimeYmdHis)
}
func Year() string {
	return time.Now().Format(consts.TimeYY)
}
func Month() string {
	return time.Now().Format(consts.TimeYYm)
}
func MonthT(t time.Time) string {
	return t.Format(consts.TimeYYm)
}
func MonthNum() string {
	return time.Now().Format(consts.TimeYYmm)
}
func MonthNumT(t time.Time) string {
	return t.Format(consts.TimeYYmm)
}
func MonthNumUnixTime(unixTime int64) string {
	return MonthNumT(ToTimeByUnix(unixTime))
}
func Data() string {
	return time.Now().Format(consts.TimeYYmd)
}
