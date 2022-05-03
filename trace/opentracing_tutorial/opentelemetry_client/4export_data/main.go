package main

import (
    "context"
    "fmt"
    "github.com/leeprince/goinfra/plog"
    "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/attribute"
    "go.opentelemetry.io/otel/exporters/jaeger"
    "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
    "go.opentelemetry.io/otel/propagation"
    "go.opentelemetry.io/otel/sdk/resource"
    sdktrace "go.opentelemetry.io/otel/sdk/trace"
    semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
    "go.opentelemetry.io/otel/trace"
    "net/http"
    "os"
    "sync"
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/7 下午3:50
 * @Desc:
 */

const (
    serviceName = "opentelemetry_client-export_data"
    env        = "local"
    port = ":8083"
)

var tracer trace.Tracer

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    
    err := plog.SetOutputFile("./", "traces.log", true)
    if err != nil {
        plog.Fatal("main:plog.SetOutputFile err:", err)
    }
    plog.SetReportCaller(true)
    
    tracerProvider, err := InitTrace(ctx, serviceName, env)
    if err != nil {
        plog.Fatal(err)
    }
    
    // Cleanly shutdown and flush telemetry when the application exits.
    defer func(ctx context.Context) {
        // Do not make the application hang when it is shutdown.
        ctx, cancel = context.WithTimeout(ctx, time.Second*5)
        defer cancel()
        if err := tracerProvider.Shutdown(ctx); err != nil {
            plog.Fatal(err)
        }
    }(ctx)
    
    // 二者都可以，底层方法一样
    // tracer = otel.Tracer(serviceName)
    tracer = tracerProvider.Tracer(serviceName)
    
    // Wrap your httpHandler function.
    handler := http.HandlerFunc(manualHttpHandler)
    wrappedHandler := otelhttp.NewHandler(handler, "main:NewHandler")
    http.Handle("/hello-export", wrappedHandler)
    
    // And start the HTTP serve.
    plog.Info("port:", port)
    plog.Fatal(http.ListenAndServe(port, nil))
}

// myHttpHandler is an HTTP handler function that is going to be instrumented.
func manualHttpHandler(w http.ResponseWriter, r *http.Request) {
    ctx, span := tracer.Start(r.Context(), "httpHandler")
    plog.Debug("4export_data:manualHttpHandler:start")
    
    // 从上下文中获取当前 span
    // span := trace.SpanFromContext(ctx)
    defer span.End()
    
    fmt.Fprintf(w, "Hello, World! I am instrumented automatically!")
    parentFunction(ctx)
    
    // Attribute keys can be precomputed
    var myKey = attribute.Key("httpHandler-SetAttributes")
    span.SetAttributes(myKey.String("a value"))
    
    // --- Events
    mutex := sync.Mutex{}
    span.AddEvent("Acquiring lock")
    mutex.Lock()
    span.AddEvent("Got lock, doing work...")
    // do stuff
    span.AddEvent("Unlocking")
    mutex.Unlock()
    
    // Events can also have attributes of their own
    span.AddEvent("Cancelled wait due to external signal", trace.WithAttributes(attribute.Int("pid", 4328)))
    
    // 传播上下文：Propagators and Context.将全局传播器设置为 tracecontext（默认为无操作）
    otel.SetTextMapPropagator(propagation.TraceContext{})
    
    plog.Debug("4export_data:manualHttpHandler:end")
}

func parentFunction(ctx context.Context) {
    ctx, parentSpan := tracer.Start(ctx, "parent")
    defer parentSpan.End()
    
    // call the child function and start a nested span in there
    childFunction(ctx)
    
    // do more work - when this function ends, parentSpan will complete.
    parentFunctionAttributes(ctx)
}

func parentFunctionAttributes(ctx context.Context) {
    // setting attributes at creation...
    ctx, span := tracer.Start(ctx, "parentFunctionAttributes-attributesAtCreation", trace.WithAttributes(attribute.String("WithAttributes-parentFunctionAttributes-k", "0001")))
    // ... and after creation
    span.SetAttributes(attribute.Bool("isTrue", true), attribute.String("parentFunctionAttributes-WithAttributes-SetAttributes", "hi!"))
    defer span.End()
    
}

func childFunction(ctx context.Context) {
    // Create a span to track `childFunction()` - this is a nested span whose parent is `parentSpan`
    ctx, childSpan := tracer.Start(ctx, "child")
    defer childSpan.End()
    
    // do work here, when this function returns, childSpan will complete.
    childFunctionAttributes(ctx)
}

func childFunctionAttributes(ctx context.Context) {
    // setting attributes at creation...
    ctx, span := tracer.Start(ctx, "childFunctionAttributes-attributesAtCreation", trace.WithAttributes(attribute.String("WithAttributes-childFunctionAttributes-k", "0001")))
    // ... and after creation
    span.SetAttributes(attribute.Bool("isTrue", true), attribute.String("childFunctionAttributes-WithAttributes-SetAttributes", "hi!"))
    defer span.End()
}

func InitTrace(ctx context.Context, serviceName, env string) (*sdktrace.TracerProvider, error) {
    // Exporter
    /*exp, err := newExporter()
    if err != nil {
        return nil, fmt.Errorf("InitTrace newExporter err:%v", err)
    }*/
    // ---
    exp, err := newJaegerExporter("http://localhost:14268/api/traces")
    if err != nil {
        return nil, fmt.Errorf("InitTrace newExporter err:%v", err)
    }
    // Exporter -end
    
    tp := sdktrace.NewTracerProvider(
        sdktrace.WithBatcher(exp),
        sdktrace.WithResource(newResource(ctx, serviceName, env)),
        // sdktrace.WithSampler(sdktrace.NeverSample()),
        sdktrace.WithSampler(sdktrace.AlwaysSample()),
    )
    
    otel.SetTracerProvider(tp)
    
    return tp, nil
}

// newResource returns a resource describing this application.
func newResource(ctx context.Context, serviceName, env string) *resource.Resource {
    r, _ := resource.Merge(
        resource.Default(),
        resource.NewWithAttributes(
            semconv.SchemaURL,
            semconv.ServiceNameKey.String(serviceName),
            attribute.String("environment", env),
        ),
    )
    
    return r
}

// newExporter returns a console exporter.
func newExporter() (sdktrace.SpanExporter, error) {
    // Your preferred exporter: console, jaeger, zipkin, OTLP, etc.
    
    // console
    // --- io.Writer: os.Stdout
    // f := os.Stdout
    // --- io.Writer: file
    /*
       // Write telemetry data to a file.
       f, err := os.Create("traces.txt")
       if err != nil {
           return nil, fmt.Errorf("InitTrace os.Create err:%v", err)
       }
    */
    // --- io.Writer: plog
    
    // 获取 plog 已经设置的日志文件及路径
    dir, fileName := plog.GetLogger().GetOutFileInfo()
    file := dir + fileName
    f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
    if err != nil {
        return nil, fmt.Errorf("InitTrace plog.SetOutputFile err:%v", err)
    }
    
    // console -end
    
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
