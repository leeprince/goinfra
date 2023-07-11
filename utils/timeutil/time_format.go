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

func ToLocalUnix(timeStr, timeLayout string) (timeUnix int64, err error) {
	loc, err := time.LoadLocation("Local")
	// loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return
	}
	t, err := time.ParseInLocation(timeLayout, timeStr, loc)
	if err != nil {
		return
	}
	timeUnix = t.Unix()
	return
}

func DataTime() string {
	return time.Now().Format(consts.TimeLayoutV1)
}
func DataTimeF(t time.Time) string {
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
func Year() string {
	return time.Now().Format(consts.TimeLayoutV8)
}
func Month() string {
	return time.Now().Format(consts.TimeLayoutV7)
}
func Data() string {
	return time.Now().Format(consts.TimeLayoutV4)
}
