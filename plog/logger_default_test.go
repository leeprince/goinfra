package plog

import (
    "errors"
    "github.com/sirupsen/logrus"
    "testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/5 上午8:57
 * @Desc:
 */

func TestDefaultLogger(t *testing.T) {
    // 测试方法在当前包中，不会执行init() 方法，手动设置默认的logger==init()
    NewDefaultLogger()
    
    Debug("prince log Debug")
    Info("prince log Info")
    Warn("prince log Warn")
    Warning("prince log Warning")
    Error("prince log Error")
    // Fatal("prince log Fatal") // 记录并结束程序允许
    // Panic("prince log Panic") // 记录并抛出异常
    
}

func TestDefaultLoggerWithField(t *testing.T) {
    Debug("prince log Debug html <br> 001")
    Info("prince log Info html <br> 001")
    
    WithField("WithField01", "WithFieldValue01").Debug("prince log Debug WithField")
    WithField("WithField02", "WithFieldValue02")
    Debug("prince log Debug WithField") // Debug logger 记录日志的方法，每次都是获取一个新的 entry。所以不会记录：WithField02,WithFieldValue02
    WithFields(map[string]interface{}{
        "WithFields001": "WithFieldsValue001V",
        "WithFields002": "WithFieldsValue002V",
    }).Debug("prince log Debug WithFields")
    WithError(errors.New("WithError01")).Debug("prince log Debug WithError")
}

func TestDefaultLoggerSetReportCaller(t *testing.T) {
    SetReportCaller(true, false)
    Debug("prince log Debug SetReportCaller")
    Debugf("prince log Debugf SetReportCaller")
    Debugln("prince log Debugln SetReportCaller")
    WithField("WithField01", "WithFieldValue01").Debug("prince log Debug WithField")
}

func TestDefaultLoggerSetFormatter(t *testing.T) {
    NewDefaultLogger()
    Debug("prince log Debug SetFormatter before")
    WithField("WithField01", "WithFieldValue01").Debug("prince log Debug WithField")
    SetFormatter(&logrus.JSONFormatter{
        TimestampFormat:   "2006-01-02 15:04:05.000000",
        DisableTimestamp:  false,
        DisableHTMLEscape: true,
        DataKey:           "data", // 允许将用户通过WithXXX设置的所有参数，放入该字段中，并且支持嵌套。不设置则平铺所有参数
        FieldMap: logrus.FieldMap{
            logrus.FieldKeyTime: "logTime",
        },
        CallerPrettyfier: nil,
        PrettyPrint:      false,
    })
    Debug("prince log Debug SetFormatter before")
    WithField("WithField01", "WithFieldValue01").Debug("prince log Debug WithField")
}

type MyHook struct{}
func (h *MyHook) Levels() []logrus.Level {
    return logrus.AllLevels
}
func (h *MyHook) Fire(entry *logrus.Entry) error {
    entry.Data["message"] =  "[MyHook] prince message001"
    return nil
}
func TestDefaultLoggerAddHook(t *testing.T) {
    AddHook(&MyHook{})
    Info("AddHook After log Info")
    NewDefaultLogger()
    AddHook(&MyHook{})
    Info("AddHook After log Info")
}
