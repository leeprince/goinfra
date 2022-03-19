package plog_test

import (
    "fmt"
    "github.com/leeprince/goinfra/plog"
    "testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/5 下午10:29
 * @Desc:
 */

func TestSetReportCallerLogger(t *testing.T) {
    fmt.Println(">>>>> 001")
    plog.NewDefaultLogger()
    plog.Debug("prince log Debug SetReportCaller")
    plog.SetReportCaller(true)
    plog.Debug("prince log Debug SetReportCaller 01")
    plog.WithField("WithField01", "WithFieldValue01").Debug("prince log Debug WithField")
    
    fmt.Println(">>>>> 002")
    plog.NewDefaultLogger()
    plog.Debug("prince log Debug SetReportCaller")
    plog.AddHookReportCaller()
    plog.Debug("prince log Debug SetReportCaller 01")
    plog.WithField("WithField01", "WithFieldValue01").Debug("prince log Debug WithField")
}