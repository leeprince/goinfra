package hooks

import (
    "fmt"
    "github.com/sirupsen/logrus"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/6 下午11:52
 * @Desc:
 */

type ReportCallerHook struct {
    dataKey  string
   	fieldKey string
   	skip     int
   	levels   []logrus.Level
}

func NewReportCallerHook(dataKey string, fieldKey string, skip int, levels ...logrus.Level) *ReportCallerHook {
	if levels == nil {
		levels = logrus.AllLevels
	}
	return &ReportCallerHook{
		dataKey:  dataKey,
		fieldKey: fieldKey,
		skip:     skip,
		levels:   levels,
	}
}

func (h *ReportCallerHook) Levels() []logrus.Level {
	return h.levels
}

func (h *ReportCallerHook) Fire(entry *logrus.Entry) error {
	obj, err := maputil.GetOrCreatePath(entry.Data, h.dataKey)
	if err != nil {
		fmt.Printf("[ERROR] gclog/caller_hook.go: GetOrCreatePath failed, dataKey = %v \n", h.dataKey)
		return err
	}

	callerName, lineNo := runtimeutil.GetCaller(h.skip)
	obj[h.fieldKey] = fmt.Sprintf("%s:%d", callerName, lineNo)
	return nil
}