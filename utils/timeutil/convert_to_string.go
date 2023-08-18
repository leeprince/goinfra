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
	TimeLayoutV1  = "2006-01-02 15:04:05"
	TimeLayoutV11 = "2006-01-02 15:04:05.999999999 -0700 MST"
	TimeLayoutV12 = "2006-01-02 15:04:05.999999999" // 机器原因，基本值都等于"2006-01-02 15:04:05.999999"
	TimeLayoutV13 = "2006-01-02 15:04:05.999999"
	TimeLayoutV14 = "2006-01-02 15:04:05.999"
	TimeLayoutV15 = "06-01-02 15:04:05"
	TimeLayoutV2  = "2006-01-02 15:04"
	TimeLayoutV3  = "2006-01-02 15"
	TimeLayoutV4  = "2006-01-02"
	TimeLayoutV41 = "2006/01/02"
	TimeLayoutV5  = "01-02-2006"
	TimeLayoutV6  = "01-02-06"
	TimeLayoutV7  = "2006-01"
	TimeLayoutV71 = "200601"
	TimeLayoutV8  = "2006"
)
*/

func DataTime() string {
	return time.Now().Format(consts.TimeLayoutV1)
}
func DataTimeT(t time.Time) string {
	return t.Format(consts.TimeLayoutV1)
}
func DataTimeNanosecond() string {
	return time.Now().Format(consts.TimeLayoutV12)
}
func DataTimeMicrosecond() string {
	return time.Now().Format(consts.TimeLayoutV13)
}
func DataTimeMillisecond() string {
	return time.Now().Format(consts.TimeLayoutV14)
}
func DataTimeDataSecond() string {
	return time.Now().Format(consts.TimeLayoutV15)
}
func Year() string {
	return time.Now().Format(consts.TimeLayoutV8)
}
func Month() string {
	return time.Now().Format(consts.TimeLayoutV7)
}
func MonthT(t time.Time) string {
	return t.Format(consts.TimeLayoutV7)
}
func MonthNum() string {
	return time.Now().Format(consts.TimeLayoutV71)
}
func MonthNumT(t time.Time) string {
	return t.Format(consts.TimeLayoutV71)
}
func MonthNumUnixTime(unixTime int64) string {
	return MonthNumT(ToTimeByUnix(unixTime))
}
func Data() string {
	return time.Now().Format(consts.TimeLayoutV4)
}
