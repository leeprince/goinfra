package hooks

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func NewFormatLogIdHook() *FormatLogIdHook {
	return &FormatLogIdHook{}
}

type FormatLogIdHook struct {
}

func (hook *FormatLogIdHook) Fire(entry *logrus.Entry) error {
	entry.Data["logId"] = "未获取到logId"
	if len(entry.Message) >= 13 {
		defer func() {
			if e := recover(); e != nil {
				entry.Message = "分割logId发生了错误:" + fmt.Sprintf("%s %s", e, entry.Message)
				entry.Data["logId"] = "获取logId错误"
			}
		}()
		logId := entry.Message[1:14]
		var e, d, o int
		for i := o; i < len(logId); i++ {
			switch {
			case 96 < logId[i] && logId[i] < 123:
				e += 1
			case 47 < logId[i] && logId[i] < 58:
				d += 1
			}
		}
		if e+d == 13 {
			entry.Data["logId"] = logId
			entry.Message = entry.Message[15 : len(entry.Message)-1]
		} else {
			entry.Data["logId"] = "获取logId失败"
		}
	}
	return nil
}

func (hook *FormatLogIdHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
