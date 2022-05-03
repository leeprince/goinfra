package main

import (
    "context"
    "fmt"
    "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/attribute"
    "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
    "go.opentelemetry.io/otel/sdk/resource"
    "go.opentelemetry.io/otel/sdk/trace"
    semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
    "log"
    "net/http"
    "os"
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/7 下午2:39
 * @Desc:
 */

// Package-level tracer.
// This should be configured in your code setup instead of here.

const serviceName = "opentelemetry_client-libraries"

var tracer = otel.Tracer(serviceName)

func main() {
    tracerProvider, err := InitTrace()
    if err != nil {
        log.Fatal(err)
    }
    // Handle shutdown properly so nothing leaks.
    defer func() {
        if err := tracerProvider.Shutdown(context.Background()); err != nil {
            log.Fatal(err)
        }
    }()
    
    // Wrap your httpHandler function.
    handler := http.HandlerFunc(librariesHttpHandler)
    wrappedHandler := otelhttp.NewHandler(handler, "main:NewHandler")
    http.Handle("/hello-instrumented", wrappedHandler)
    
    // And start the HTTP serve.
    log.Fatal(http.ListenAndServe(":8200", nil))
}

// httpHandler is an HTTP handler function that is going to be instrumented.
func librariesHttpHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World! I am instrumented automatically!")
    ctx := r.Context()
    sleepy(ctx)
}

// sleepy mocks work that your application does.
func sleepy(ctx context.Context) {
    _, span := tracer.Start(ctx, "sleep")
    defer span.End()
    
    sleepTime := 1 * time.Second
    fmt.Println("time.Sleep...")
    time.Sleep(sleepTime)
    fmt.Println("time.Sleep -end")
    
    span.SetAttributes(attribute.Int("sleep.duration", int(sleepTime)))
}

func InitTrace() (*trace.TracerProvider, error) {
    exp, err := newExporter()
    if err != nil {
        return nil, fmt.Errorf("InitTrace newExporter err:%v", err)
    }
    
    tp := trace.NewTracerProvider(
        trace.WithBatcher(exp),
        trace.WithResource(newResource()),
    )
    
    otel.SetTracerProvider(tp)
    return tp, nil
}

// newResource returns a resource describing this application.
func newResource() *resource.Resource {
    r, _ := resource.Merge(
        resource.Default(),
        resource.NewWithAttributes(
            semconv.SchemaURL,
            semconv.ServiceNameKey.String(serviceName),
            semconv.ServiceVersionKey.String("v0.1.0"),
            attribute.String("environment", "demo"),
        ),
    )
    return r
}

// newExporter returns a console exporter.
func newExporter() (trace.SpanExporter, error) {
    // Your preferred exporter: console, jaeger, zipkin, OTLP, etc.
    
    // console
    // --- io.Writer: os.Stdout
    f := os.Stdout
    // --- io.Writer: file
    /*
       // Write telemetry data to a file.
       f, err := os.Create("traces.txt")
       if err != nil {
           return nil, fmt.Errorf("InitTrace os.Create err:%v", err)
       }
    */
    // --- io.Writer: plog
    /*
       // 获取 plog 已经设置的日志文件及路径
       dir, fileName := plog.GetLogger().GetOutFileInfo()
       file := dir + fileName
       f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
       if err != nil {
           return nil, fmt.Errorf("InitTrace plog.SetOutputFile err:%v", err)
       }
    */
    // console -end
    
    return stdouttrace.New(
        stdouttrace.WithWriter(f),
        // Use human-readable output.
        stdouttrace.WithPrettyPrint(),
        // Do not print timestamps for the demo.
        // stdouttrace.WithoutTimestamps(),
    )
}
