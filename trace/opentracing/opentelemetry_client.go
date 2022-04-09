package opentracing

import (
    "github.com/leeprince/goinfra/plog"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/attribute"
    "go.opentelemetry.io/otel/exporters/jaeger"
    "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
    "go.opentelemetry.io/otel/sdk/resource"
    sdktrace "go.opentelemetry.io/otel/sdk/trace"
    semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
    "io"
    "os"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/9 上午10:27
 * @Desc:
 */

// 初始化 Telemetry 客户端
func InitTelemetryTrace(serverName, env string, exporter sdktrace.SpanExporter) (*sdktrace.TracerProvider, error) {
    tp := sdktrace.NewTracerProvider(
        sdktrace.WithBatcher(exporter),
        sdktrace.WithResource(NewResource(serverName, env)),
        // sdktrace.WithSampler(sdktrace.NeverSample()),
        sdktrace.WithSampler(sdktrace.AlwaysSample()),
    )
    
    otel.SetTracerProvider(tp)
    
    return tp, nil
}

// 资源：描述应用的资源
func NewResource(serverName, env string) *resource.Resource {
    r, _ := resource.Merge(
        resource.Default(),
        resource.NewWithAttributes(
            semconv.SchemaURL,
            semconv.ServiceNameKey.String(serverName),
            attribute.String("environment", env),
        ),
    )
    
    return r
}

// --- 导出器
type writerExporterType int32
const (
    WriterExporterStdout writerExporterType = iota // os.Stdout 输出
    WriterExporterCreate    // 创建 traces.txt 输出。启动后覆盖
    WriterExporterPlog
)
// --- 导出器：io.Writer 作为导出器（exporter）
func NewWriterExporter(t writerExporterType) (spanExporter sdktrace.SpanExporter) {
    // Your preferred exporter: console, jaeger, zipkin, OTLP, etc.
    
    var (
        f   io.Writer
        err error
    )
    switch t {
    case WriterExporterStdout:
        f = os.Stdout
    case WriterExporterCreate:
        f, err = os.Create("traces.txt")
        if err != nil {
            plog.Fatal("NewWriterExporter switch t WriterExporterCreate os.Create err:", err)
        }
    case WriterExporterPlog:
        dir, fileName := plog.GetLogger().GetOutFileInfo()
        if dir == "" || fileName == "" {
            plog.Fatal("NewWriterExporter switch t WriterExporterPlog dir == '' || fileName == ''")
        }
        file := dir + fileName
        // file 存在则肯定已创建
        f, err = os.OpenFile(file, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
        if err != nil {
            plog.Fatal("NewWriterExporter os.OpenFile err:", err)
        }
    default:
        plog.Fatal("NewWriterExporter switch t default error")
    }
    
    spanExporter, err = stdouttrace.New(
        stdouttrace.WithWriter(f),
        // Use human-readable output.
        stdouttrace.WithPrettyPrint(),
        // Do not print timestamps for the demo.
        // stdouttrace.WithoutTimestamps(),
    )
    if err != nil {
        plog.Fatal("NewWriterExporter stdouttrace.New err:", err)
    }
    
    return
}
// --- 导出器：io.Writer 作为导出器（exporter） -end


// --- 导出器：Jaeger HTTP Thrift collector 作为导出器（exporter）
func NewJaegerExporter(url string) (sdktrace.SpanExporter, error) {
    // Your preferred exporter: console, jaeger, zipkin, OTLP, etc.
    
    if url == "" {
        plog.Fatal("NewJaegerExporter url == ''")
    }
    return jaeger.New(
        jaeger.WithCollectorEndpoint(
            jaeger.WithEndpoint(url),
        ),
    )
}
// --- 导出器：Jaeger HTTP Thrift collector 作为导出器（exporter） -end
// --- 导出器 -end
