package plog

import (
    "errors"
    "fmt"
    jsoniter "github.com/json-iterator/go"
    "github.com/leeprince/goinfra/utils"
    "testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/5 下午10:29
 * @Desc:
 */

func TestSetJsoniterFormatter(t *testing.T) {
    NewDefaultLogger()
    WithField("WithField01", "WithFieldValue01 中国，我爱你").Debug("prince log Debug WithField")
    WithField("WithField02", "WithFieldValue02 中国，我爱你")
    Debug("prince log Debug WithField") // Debug logger 记录日志的方法，每次都是获取一个新的 entry。所以不会记录：WithField02,WithFieldValue02 中国，我爱你
    WithFields(map[string]interface{}{
        "WithFields001": "WithFieldsValue001V",
        "WithFields002": "WithFieldsValue002V",
    }).Debug("prince log Debug WithFields")
    WithError(errors.New("WithError01")).Debug("prince log Debug WithError")
    
    fmt.Println("--- ...")
    NewDefaultLogger()
    SetFormatterJsonInter(jsoniter.ConfigCompatibleWithStandardLibrary)
    WithField("WithField01", "WithFieldValue01 中国，我爱你").Debug("prince log Debug WithField")
    WithField("WithField02", "WithFieldValue02 中国，我爱你")
    Debug("prince log Debug WithField") // Debug logger 记录日志的方法，每次都是获取一个新的 entry。所以不会记录：WithField02,WithFieldValue02 中国，我爱你
    WithFields(map[string]interface{}{
        "WithFields001": "WithFieldsValue001 中国，我爱你",
        "WithFields002": "WithFieldsValue002 中国，我爱你",
    }).Debug("prince log Debug WithFields")
    WithError(errors.New("WithError01")).Debug("prince log Debug WithError")
}

func TestSetOutputFile(t *testing.T) {
    NewDefaultLogger()
    Debug("prince log Debug")
    Info("prince log Info")
    WithField("WithField01", "WithFieldValue01 中国，我爱你").Debug("prince log Debug WithField")
    err := SetOutputFile("./", "application.log", true)
    if err != nil {
        fmt.Println("SetOutputFile err:", err)
        return
    }
    fmt.Println("--- ...")
    Debug("prince log Debug")
    Info("prince log Info")
    WithField("WithField01", "WithFieldValue01 中国，我爱你").Debug("prince log Debug WithField")
    
    fmt.Println("--- ... 00")
    NewDefaultLogger()
    err = SetOutputFile("./", "application.log", false)
    if err != nil {
        fmt.Println("SetOutputFile err:", err)
        return
    }
    Debug("prince log Debug 00")
    Info("prince log Info 00")
    WithField("WithField01", "WithFieldValue01 中国，我爱你 00").Debug("prince log Debug WithField")
    
    fmt.Println("--- ... 01")
    NewDefaultLogger()
    err = SetOutputFile("./logs/", "application.log", false)
    // err = SetOutputFile("./logs/l/", "application.log", false)
    if err != nil {
        fmt.Println("SetOutputFile err:", err)
        return
    }
    Debug("prince log Debug 01")
    Info("prince log Info 01")
    WithField("WithField01", "WithFieldValue01 中国，我爱你 01").Debug("prince log Debug WithField")
}


func TestSetOutputRotateFile(t *testing.T) {
    NewDefaultLogger()
    WithField("WithField01", "WithFieldValue01 中国，我爱你 0000").Debug("prince log Debug WithField")
    
    err := SetOutputRotateFile("./logs/", "application.log", true, nil)
    if err != nil {
        fmt.Println("SetOutputFile err:", err)
        return
    }
    fmt.Println("--- ...")
    WithField("WithField01", "WithFieldValue01 中国，我爱你 0001").Debug("prince log Debug WithField")
}

func TestSetReportCallerLogger(t *testing.T) {
    NewDefaultLogger()
    Debug("prince log Debug SetReportCaller")
    SetReportCaller(true)
    Debug("prince log Debug SetReportCaller 01")
    WithField("WithField01", "WithFieldValue01").Debug("prince log Debug WithField")
    
    fmt.Println(">>>>> 002")
    NewDefaultLogger()
    Debug("prince log Debug SetReportCaller")
    // 因为 AddHookReportCaller 目标是`检索第一个非 plog 包调用函数的名称`, 所以在当前包中测试不准确
    // 测试 AddHookReportCaller 的方法应在与 plog 包同目录的 plog_test 的 logger_option_test.go 中测试
    SetReportCaller(true)
    Debug("prince log Debug SetReportCaller 01")
    WithField("WithField01", "WithFieldValue01").Debug("prince log Debug WithField")
}

func TestWithFiledLogID(t *testing.T) {
    Debug("prince log Debug SetReportCaller 01")
    WithField("WithField01", "WithFieldValue01").Debug("prince log Debug WithField")
    
    logID := utils.UniqID()
    WithFiledLogID(logID).Debug("prince log Debug")
    Debug("prince log Debug")
    WithFiledLogID(logID).WithField("WithField01", "WithFieldValue01").Debug("prince log Debug WithField")
    Debug("prince log WithFiled")
}

func TestAddHookSentry(t *testing.T) {
    dsn := "http://58be04091efa42feb1aa18390230bf2f@127.0.0.1:9100/1"
    
    Info("TestAddHookSentry Info 001")
    
    AddHookSentry(dsn)
    Info("TestAddHookSentry Info 002")
}