package opentracing

import (
    "context"
    "fmt"
    "github.com/leeprince/goinfra/plog"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/attribute"
    "go.opentelemetry.io/otel/exporters/jaeger"
    "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
    "go.opentelemetry.io/otel/sdk/resource"
    sdktrace "go.opentelemetry.io/otel/sdk/trace"
    semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
    "os"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/9 上午10:27
 * @Desc:
 */

// 初始化 Telemetry 客户端
func InitTelemetryTrace(ctx context.Context, serverName, env string) (*sdktrace.TracerProvider, error) {
    // Exporter
    /*exp, err := newExporter()
      if err != nil {
          return nil, fmt.Errorf("InitTrace newExporter err:%v", err)
      }*/
    exp, err := newJaegerExporter("http://localhost:14268/api/traces")
    if err != nil {
        return nil, fmt.Errorf("InitTrace newExporter err:%v", err)
    }
    
    tp := sdktrace.NewTracerProvider(
        sdktrace.WithBatcher(exp),
        sdktrace.WithResource(newResource(serverName, env)),
        // sdktrace.WithSampler(sdktrace.NeverSample()),
        sdktrace.WithSampler(sdktrace.AlwaysSample()),
    )
    
    otel.SetTracerProvider(tp)
    
    return tp, nil
}

// 资源：描述应用的资源
func newResource(serverName, env string) *resource.Resource {
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

// newExporter returns a console exporter.
func newExporter() (sdktrace.SpanExporter, error) {
    // Your preferred exporter: console, jaeger, zipkin, OTLP, etc.
    
    // Your preferred exporter: console, jaeger, zipkin, OTLP, etc.
    
    // console
    // io.Writer: os.Stdout
    // Write telemetry data to os.Stdout
    /*f := os.Stdout
      exp, err := newExporter(f)
      if err != nil {
          l.Fatal(err)
      }*/
    // --- io.Writer: file
    /*// Write telemetry data to a file.
      f, err := os.Create("traces.txt")
      if err != nil {
          return nil, fmt.Errorf("InitTrace os.Create err:%v", err)
      }*/
    // --- io.Writer: plog // TODO:  - prince@todo 2022/4/9 下午12:12
    // 获取 plog 已经设置的日志文件及路径
    dir, fileName := plog.GetLogger().GetOutFileInfo()
    file := dir + fileName
    f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
    if err != nil {
        return nil, fmt.Errorf("InitTrace plog.SetOutputFile err:%v", err)
    }
    // console
    
    return stdouttrace.New(
        stdouttrace.WithWriter(f),
        // Use human-readable output.
        stdouttrace.WithPrettyPrint(),
        // Do not print timestamps for the demo.
        // stdouttrace.WithoutTimestamps(),
    )
}

// newExporter returns a console exporter.
func newJaegerExporter(url string) (sdktrace.SpanExporter, error) {
    // Your preferred exporter: console, jaeger, zipkin, OTLP, etc.
    
    // Create the Jaeger exporter
    return jaeger.New(
        jaeger.WithCollectorEndpoint(
            jaeger.WithEndpoint(url),
        ),
    )
}
