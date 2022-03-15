package plog

import (
    jsoniter "github.com/json-iterator/go"
    "github.com/sirupsen/logrus"
    "goinfra/plog/formatters"
    "goinfra/plog/hooks"
    "goinfra/resource/file"
    "io"
    "os"
    "path/filepath"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/5 上午8:40
 * @Desc:
 */

// --- 设置 logger 参数
// jsoniter 更快的json序列化方式
// jsoniterAPI =
// 	jsoniter.ConfigDefault |
// 	jsoniter.ConfigCompatibleWithStandardLibrary() (推荐：100% 兼容json标准库:"encoding/json") |
// 	jsoniter.ConfigFastest
func SetFormatterJsonInter(jsoniterAPI jsoniter.API) {
    formatter := &formatters.JSONFormatter{
        TimestampFormat:  "2006-01-02 15:04:05.000",
        DisableTimestamp: false,
        DataKey:          "data",
        FieldMap: formatters.FieldMap{
            logrus.FieldKeyTime: "logTime",
        },
        CallerPrettyfier: nil,
        PrettyPrint:      false,
        JSON:             jsoniterAPI,
    }
    SetFormatter(formatter)
}

// dirPath 支持多层级目录结构
// isBothStdout：是否支持同时通过 os.Stdout 记录日志
func SetOutputFile(dirPath, filename string, isBothStdout bool) error {
    var writer io.Writer
    var err error
    
    filePath := filepath.Join(dirPath, filename)
    if _, ok := file.CheckFileExist(filePath); ok {
        writer, err = os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
    } else {
        // 创建目录
        err = os.MkdirAll(dirPath, os.ModePerm)
        if ok = os.IsNotExist(err); ok {
            return err
        }
        writer, err = os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
    }
    if err != nil {
        return err
    }
    
    if logger.Out == nil {
        SetOutput(writer)
    }
    
    if isBothStdout {
        // 支持 logger.Out 和 *io.file(可用于io) 的记录日志方式
        // MultiWriter同时接受了3种数据类型，分别是io.Writer、*os.File、io.WriteCloser
        SetOutput(io.MultiWriter(logger.Out, writer))
    } else {
        SetOutput(writer)
    }
    
    return nil
}

// SetReportCaller
//  是否记录日志调用者的标记(位置信息)。
//  实现：github.com/sirupsen/logrus@v1.8.1/entry.go@getCaller 方法
//  说明：retrieves the name of the first non-logrus calling function（译：检索第一个非 logrus 包调用函数的名称）
// AddHookOfReportCaller
//  目标：检索第一个非 plog 包调用函数的名称
func AddHookOfReportCaller(reportCaller bool) {
    // 开启 ReportCaller
	SetReportCaller(true)
	
    AddHook(hooks.NewReportCallerHook(logrus.DebugLevel))
}

// --- 设置 logger 参数 - end
