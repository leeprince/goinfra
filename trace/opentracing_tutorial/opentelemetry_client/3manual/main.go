package main

import (
    "context"
    "fmt"
    "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/attribute"
    "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
    "go.opentelemetry.io/otel/propagation"
    "go.opentelemetry.io/otel/sdk/resource"
    sdktrace "go.opentelemetry.io/otel/sdk/trace"
    semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
    "go.opentelemetry.io/otel/trace"
    "log"
    "net/http"
    "os"
    "sync"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/7 下午3:50
 * @Desc:
 */

const serverName = "opentelemetry_client-manual"

var tracer trace.Tracer

func main() {
    ctx := context.Background()
    
    tracerProvider, err := InitTrace(ctx, serverName)
    if err != nil {
        log.Fatal(err)
    }
    
    // Handle shutdown properly so nothing leaks.
    defer func() { _ = tracerProvider.Shutdown(ctx) }()
    
    // 二者都可以，推荐 tracerProvider.Tracer
    // tracer = otel.Tracer(serverName)
    tracer = tracerProvider.Tracer(serverName)
    
    // Wrap your httpHandler function.
    handler := http.HandlerFunc(manualHttpHandler)
    wrappedHandler := otelhttp.NewHandler(handler, "main:NewHandler")
    http.Handle("/hello-manual", wrappedHandler)
    
    // And start the HTTP serve.
    log.Fatal(http.ListenAndServe(":8200", nil))
}

// myHttpHandler is an HTTP handler function that is going to be instrumented.
func manualHttpHandler(w http.ResponseWriter, r *http.Request) {
    ctx, span := tracer.Start(r.Context(), "httpHandler")
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

func InitTrace(ctx context.Context, serverName string) (*sdktrace.TracerProvider, error) {
    exp, err := newExporter()
    if err != nil {
        return nil, fmt.Errorf("InitTrace newExporter err:%v", err)
    }
    
    tp := sdktrace.NewTracerProvider(
        sdktrace.WithBatcher(exp),
        sdktrace.WithResource(newResource(serverName)),
    )
    
    otel.SetTracerProvider(tp)
    
    return tp, nil
}

// newResource returns a resource describing this application.
func newResource(serverName string) *resource.Resource {
    r, _ := resource.Merge(
        resource.Default(),
        resource.NewWithAttributes(
            semconv.SchemaURL,
            semconv.ServiceNameKey.String(serverName),
            attribute.String("environment", "local"),
        ),
    )
    return r
}

// newExporter returns a console exporter.
func newExporter() (sdktrace.SpanExporter, error) {
    // Your preferred exporter: console, jaeger, zipkin, OTLP, etc.
    
    // io.Writer
    // Write telemetry data to os.Stdout
    /*f := os.Stdout
      exp, err := newExporter(f)
      if err != nil {
          l.Fatal(err)
      }*/
    // io.Writer
    // Write telemetry data to a file.
    f, err := os.Create("traces.txt")
    if err != nil {
        return nil, fmt.Errorf("InitTrace os.Create err:%v", err)
    }
    
    return stdouttrace.New(
        stdouttrace.WithWriter(f),
        // Use human-readable output.
        stdouttrace.WithPrettyPrint(),
        // Do not print timestamps for the demo.
        // stdouttrace.WithoutTimestamps(),
    )
}
