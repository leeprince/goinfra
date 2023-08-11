package panicutil

import (
	"github.com/leeprince/goinfra/plog"
	"github.com/sirupsen/logrus"
	"runtime"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/13 01:04
 * @Desc:
 */

func PanicRecover(recover interface{}, logIDs ...string) error {
	var plogEntry *logrus.Entry
	if len(logIDs) > 0 && logIDs[0] != "" {
		plogEntry = plog.LogID(logIDs[0]).WithField("method", "PanicRecover")
	} else {
		plogEntry = plog.WithField("method", "PanicRecover")
	}
	
	// 断言错误类型
	reconverErr, isError := recover.(error)
	if isError {
		plogEntry.Error("isError reconverErr:", reconverErr)
	}
	
	// 获取panic发生的位置
	_, file, line, ok := runtime.Caller(3)
	if ok {
		plogEntry.Errorf("Panic occurred at %s:%d\n", file, line)
	}
	
	return reconverErr
}
