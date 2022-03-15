package plog

import (
    "errors"
    "fmt"
    jsoniter "github.com/json-iterator/go"
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
    err = SetOutputFile("./log/l/", "application.log", false)
    if err != nil {
        fmt.Println("SetOutputFile err:", err)
        return
    }
    Debug("prince log Debug 01")
    Info("prince log Info 01")
    WithField("WithField01", "WithFieldValue01 中国，我爱你 01").Debug("prince log Debug WithField")
}

func TestSetReportCallerLogger(t *testing.T) {
    NewDefaultLogger()
    
    Debug("prince log Debug SetReportCaller")
    SetReportCaller(true)
    Debug("prince log Debug SetReportCaller 01")
}