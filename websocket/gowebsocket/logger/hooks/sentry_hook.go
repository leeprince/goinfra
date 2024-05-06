package hooks

import (
	"github.com/getsentry/raven-go"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"strconv"
)

var EnableSentry = false

func NewSentryHook() *SentryHook {
	return &SentryHook{}
}

type SentryHook struct {
}

func (hook *SentryHook) Fire(entry *logrus.Entry) error {
	//if entry.Level == logrus.PanicLevel || entry.Level == logrus.FatalLevel || entry.Level == logrus.ErrorLevel || entry.Level == logrus.WarnLevel{
	context := make(map[string]string)
	var str = ""
	for k, v := range entry.Data {
		switch v.(type) {
		case string:
			str = v.(string)
		case int:
			str = strconv.Itoa(v.(int))
		case bool:
			if v.(bool) {
				str = "true"
			} else {
				str = "false"
			}
		default:
			str = "未支持数据类型"
		}
		context[k] = str
	}
	raven.CaptureError(errors.New(entry.Message), context)
	return nil
}

func (hook *SentryHook) Levels() []logrus.Level {
	levels := []logrus.Level{logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel, logrus.WarnLevel}
	return levels
}

func SetSentryDSN(dsn string) {
	raven.SetDSN(dsn)
}

func SetEnableSentry() {
	EnableSentry = true
}
