package opentelemetry_client

import (
    "github.com/leeprince/goinfra/plog"
    "go.opentelemetry.io/otel/exporters/jaeger"
    "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
    sdktrace "go.opentelemetry.io/otel/sdk/trace"
    "io"
    "os"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/15 下午6:18
 * @Desc:   Exporter 导出器
 */

type exporterType int32
const (
    IOWriterExporterStdout exporterType = iota // os.Stdout 输出
    IOWriterExporterCreate                     // 创建 traces.txt 输出。启动后覆盖
    IOWriterExporterPlog
)
// --- 导出器：io.Writer 作为导出器（exporter）
func NewIOWriterWExporter(t exporterType) (spanExporter sdktrace.SpanExporter) {
    // Your preferred exporter: console, jaeger, zipkin, OTLP, etc.
    
    var (
        f   io.Writer
        err error
    )
    switch t {
    case IOWriterExporterStdout:
        f = os.Stdout
    case IOWriterExporterCreate:
        f, err = os.Create("traces.txt")
        if err != nil {
            plog.Fatal("NewIOWriterWExporter switch t IOWriterExporterCreate os.Create err:", err)
        }
    case IOWriterExporterPlog:
        dir, fileName := plog.GetLogger().GetOutFileInfo()
        if dir == "" || fileName == "" {
            plog.Fatal("NewIOWriterWExporter switch t IOWriterExporterPlog dir == '' || fileName == ''")
        }
        // dir || fileName 都不为空的情况下，file 肯定已创建
        file := dir + fileName
        f, err = os.OpenFile(file, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
        if err != nil {
            plog.Fatal("NewIOWriterWExporter os.OpenFile err:", err)
        }
    default:
        plog.Fatal("NewIOWriterWExporter switch t default error")
    }
    
    spanExporter, err = stdouttrace.New(
        stdouttrace.WithWriter(f),
        // Use human-readable output.
        stdouttrace.WithPrettyPrint(),
        // Do not print timestamps for the demo.
        // stdouttrace.WithoutTimestamps(),
    )
    if err != nil {
        plog.Fatal("NewIOWriterWExporter stdouttrace.New err:", err)
    }
    
    return
}
// --- 导出器：io.Writer 作为导出器（exporter） -end

// --- 导出器：Jaeger HTTP Thrift collector 作为导出器（exporter）
func NewJaegerExporter(url string, options ...jaeger.CollectorEndpointOption) sdktrace.SpanExporter {
    // Your preferred exporter: console, jaeger, zipkin, OTLP, etc.
    
    // 、默认 `http://localhost:14268/api/traces`
    if url != "" {
        options = append(options, jaeger.WithEndpoint(url))
    }
    
    spanExporter, err := jaeger.New(
        jaeger.WithCollectorEndpoint(
            options...,
        ),
    )
    if err != nil {
        plog.Fatal("NewJaegerExporter jaeger.New err:", err)
    }
    
    return spanExporter
}
// --- 导出器：Jaeger HTTP Thrift collector 作为导出器（exporter） -end
