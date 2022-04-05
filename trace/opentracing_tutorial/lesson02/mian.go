package main

import (
    "context"
    "fmt"
    "github.com/leeprince/goinfra/plog"
    "github.com/opentracing/opentracing-go"
    "github.com/opentracing/opentracing-go/log"
    "github.com/uber/jaeger-client-go"
    "github.com/uber/jaeger-client-go/config"
    "io"
    "time"
)

const (
    serverName = "princeJaeger-lesson02"
)

func main() {
    helloStr := fmt.Sprintf("Hello, world!")
    println(helloStr)
    
    // 注意：function opentracing.GlobalTracer() returns a no-op tracer by default.
    // tracer := opentracing.GlobalTracer()
    
    // tracer, closer := initJaeger(serverName)
    tracer, closer := initJaegerLog(serverName)
    defer closer.Close()
    
    println("tracer.StartSpan>>>>")
    span := tracer.StartSpan("say-hello")
    defer span.Finish()
    
    // --- span 的方法移动到具体的函数中
    SpanSetTag(span, "SetTag:princeTag001", "value001")
    SpanLogFields(span, map[string]string{
        "LogFields:event001": "string-format",
        "LogFields:value001": helloStr,
    })
    SpanLogKV(span, "LogKV:event001", "println")
    // --- span 的方法移动到具体的函数中 -end
    
    // --- 每个函数包装到它自己的 span 中
    RootSpanSetTag(span, "SetTag:princeTag002", "value001")
    RootSpanLogFields(span, map[string]string{
        "LogFields:event002": "string-format",
        "LogFields:value002": helloStr,
    })
    RootSpanLogKV(span, "LogKV:event002", "println")
    // --- 每个函数包装到它自己的 span 中 -end
    
    // 传递上下文 context 代替将 span 作为每个函数的第一个参数【最终方案】
    opentracing.SetGlobalTracer(tracer)
    ctx := opentracing.ContextWithSpan(context.Background(), span)
    
    ContextSpanSetTag(ctx, "SetTag:princeTag003", "value001")
    ContextSpanLogFields(ctx, map[string]string{
        "LogFields:event003": "string-format",
        "LogFields:value003": helloStr,
    })
    ContextSpanLogKV(ctx, "LogKV:event003", "println")
    // 传递上下文 context 代替将 span 作为每个函数的第一个参数 -end
}

func SpanSetTag(span opentracing.Span, key string, value interface{}) {
    println("SpanSetTag@span.SetTag>>>>", time.Now().Format("2006-01-02 15:04:05.999999999"))
    span.SetTag(key, value)
}

func SpanLogFields(span opentracing.Span, fileds map[string]string) {
    println("SpanLogFields@span.LogFields>>>>", time.Now().Format("2006-01-02 15:04:05.999999999"))
    var logFileds []log.Field
    for key, filed := range fileds {
        logFileds = append(logFileds, log.String(key, filed))
    }
    span.LogFields(logFileds...)
}

func SpanLogKV(span opentracing.Span, alternatingKeyValues ...interface{}) {
    println("SpanLogKV@span.LogKV>>>>", time.Now().Format("2006-01-02 15:04:05.999999999"))
    span.LogKV(alternatingKeyValues...)
}

func RootSpanSetTag(rootSpan opentracing.Span, key string, value interface{}) {
    // 根据根追踪启动一个新的 span, 但是新的的 span 未与根的 span 放在同一个根追踪中
    // span := rootSpan.Tracer().StartSpan("RootSpanSetTag")
    // Combine multiple spans into a single trace(将多个 span 合并到一个根追踪中)
    span := rootSpan.Tracer().StartSpan(
        "RootSpanSetTag",
        opentracing.ChildOf(rootSpan.Context()),
    )
    
    defer span.Finish()
    
    println("RootSpanSetTag@span.SetTag>>>>", time.Now().Format("2006-01-02 15:04:05.999999999"))
    span.SetTag(key, value)
}

func RootSpanLogFields(rootSpan opentracing.Span, fileds map[string]string) {
    // 根据根追踪启动一个新的 span, 但是新的的 span 未与根的 span 放在同一个根追踪中
    // span := rootSpan.Tracer().StartSpan("RootSpanLogFields")
    // Combine multiple spans into a single trace(将多个 span 合并到一个根追踪中)
    span := rootSpan.Tracer().StartSpan(
        "RootSpanLogFields",
        opentracing.ChildOf(rootSpan.Context()),
    )
    defer span.Finish()
    
    println("RootSpanLogFields@span.LogFields>>>>", time.Now().Format("2006-01-02 15:04:05.999999999"))
    var logFileds []log.Field
    for key, filed := range fileds {
        logFileds = append(logFileds, log.String(key, filed))
    }
    span.LogFields(logFileds...)
}

func RootSpanLogKV(rootSpan opentracing.Span, alternatingKeyValues ...interface{}) {
    // 根据根追踪启动一个新的 span, 但是新的的 span 未与根的 span 放在同一个根追踪中
    // span := rootSpan.Tracer().StartSpan("RootSpanLogKV")
    // Combine multiple spans into a single trace(将多个 span 合并到一个根追踪中)
    span := rootSpan.Tracer().StartSpan(
        "RootSpanLogKV",
        opentracing.ChildOf(rootSpan.Context()),
    )
    defer span.Finish()
    
    println("RootSpanLogKV@span.LogKV>>>>", time.Now().Format("2006-01-02 15:04:05.999999999"))
    span.LogKV(alternatingKeyValues...)
}

func ContextSpanSetTag(ctx context.Context, key string, value interface{}) {
    // - 传递上下文 context 代替将 span 作为每个函数的第一个参数
    span, _ := opentracing.StartSpanFromContext(ctx, "ContextSpanSetTag")
    defer span.Finish()
    
    println("ContextSpanSetTag@span.SetTag>>>>", time.Now().Format("2006-01-02 15:04:05.999999999"))
    span.SetTag(key, value)
}

func ContextSpanLogFields(ctx context.Context, fileds map[string]string) {
    // - 传递上下文 context 代替将 span 作为每个函数的第一个参数
    span, _ := opentracing.StartSpanFromContext(ctx, "ContextSpanLogFields")
    defer span.Finish()
    
    println("ContextSpanLogFields@span.LogFields>>>>", time.Now().Format("2006-01-02 15:04:05.999999999"))
    var logFileds []log.Field
    for key, filed := range fileds {
        logFileds = append(logFileds, log.String(key, filed))
    }
    span.LogFields(logFileds...)
}

func ContextSpanLogKV(ctx context.Context, alternatingKeyValues ...interface{}) {
    // - 传递上下文 context 代替将 span 作为每个函数的第一个参数
    span, _ := opentracing.StartSpanFromContext(ctx, "ContextSpanLogKV")
    defer span.Finish()
    
    println("ContextSpanLogKV@span.LogKV>>>>", time.Now().Format("2006-01-02 15:04:05.999999999"))
    span.LogKV(alternatingKeyValues...)
}

// initJaeger returns an instance of Jaeger Tracer that samples 100% of traces and logs all spans to stdout.
func initJaeger(service string) (opentracing.Tracer, io.Closer) {
    cfg := &config.Configuration{
        ServiceName: service,
        Sampler: &config.SamplerConfig{
            Type:  "const",
            Param: 1,
        },
        Reporter: &config.ReporterConfig{
            LogSpans: true,
        },
    }
    tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
    if err != nil {
        panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
    }
    return tracer, closer
}

// initJaeger returns an instance of Jaeger Tracer that samples 100% of traces and logs all spans to stdout.
func initJaegerLog(service string) (opentracing.Tracer, io.Closer) {
    cfg := &config.Configuration{
        ServiceName: service,
        Sampler: &config.SamplerConfig{
            Type:  "const",
            Param: 1,
        },
        Reporter: &config.ReporterConfig{
            LogSpans: true,
        },
    }
    err := plog.SetOutputFile("./", "application.log", true)
    if err != nil {
        panic(fmt.Sprintf("plog.SetOutputFile error:%v", err))
    }
    tracer, closer, err := cfg.NewTracer(config.Logger(jaegerLoggerPlog))
    if err != nil {
        panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
    }
    return tracer, closer
}

var jaegerLoggerPlog = &jaegerLogger{}
type jaegerLogger struct{}
func (l *jaegerLogger) Error(msg string) {
	plog.Error(msg)
}
func (l *jaegerLogger) Infof(msg string, args ...interface{}) {
    plog.Infof(msg, args...)
}
