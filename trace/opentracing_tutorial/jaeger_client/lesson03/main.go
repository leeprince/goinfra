package main

import (
    "context"
    "fmt"
    "github.com/leeprince/goinfra/http/httpcli"
    "github.com/leeprince/goinfra/plog"
    "github.com/opentracing/opentracing-go"
    "github.com/opentracing/opentracing-go/ext"
    "github.com/uber/jaeger-client-go"
    "github.com/uber/jaeger-client-go/config"
    "io"
    "net/http"
    "net/url"
    "time"
)

const (
    serverName = "princeJaeger-lesson03"
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
    
    // 使用 span 的 Baggage 功能
    // after starting the span
    span.SetBaggageItem("seq", "prince-seq-202204060001")
    
    // 传递上下文 context 代替将 span 作为每个函数的第一个参数【最终方案】
    ctx := opentracing.ContextWithSpan(context.Background(), span)
    
    // --- 只有当前服务记录了追踪，暂未将当前服务与远程调用服务的调用链追踪连接成一个完整调用链追踪
    RPCFormatter(ctx)
    RPCPublisher(ctx)
    // --- 只有当前服务记录了追踪，暂未将当前服务与远程调用服务的调用链追踪连接成一个完整调用链追踪 -end
    
    // --- 当前服务与远程调用服务组成一个完整调用链追踪
    RPCTraceFormatter(ctx)
    RPCTracePublisher(ctx)
    // --- 当前服务与远程调用服务组成一个完整调用链追踪 -end
    
    // --- 当前服务与远程调用服务组成一个完整调用链追踪，并使用 span 的 Baggage 功能
    // 对应的课程：https://github.com/leeprince/opentracing-tutorial/tree/master/go/lesson04
    RPCTraceBaggageFormatter(ctx)
    RPCTraceBaggagePublisher(ctx)
    // --- 当前服务与远程调用服务组成一个完整调用链追踪，并使用 span 的 Baggage 功能 -end
    // 传递上下文 context 代替将 span 作为每个函数的第一个参数 -end
}

func RPCFormatter(ctx context.Context) {
    // - 传递上下文 context 代替将 span 作为每个函数的第一个参数
    span, _ := opentracing.StartSpanFromContext(ctx, "RPCFormatter")
    defer span.Finish()
    println("RPCFormatter@span.LogKV>>>>", time.Now().Format("2006-01-02 15:04:05.999999999"))
    span.LogKV("RPCFormatter@LogKV:event004", "println")
    
    params := url.Values{
        "helloTo": []string{"prince"},
    }
    urlPath := "http://127.0.0.1:8101/format?" + params.Encode()
    req, err := http.NewRequest(http.MethodGet, urlPath, nil)
    if err != nil {
        plog.Error("RPCFormatter http.NewRequest err:", err)
        return
    }
    
    // 使用 ext.LogError 设置 Tag=error 标记 span 错误
    // ext.LogError(span, errors.New("RPCFormatter>>>>>>>>>>>>>>>>>>"))
    
    bodyByte, err := httpcli.Do(req)
    if err != nil {
        plog.Error("RPCFormatter httpcli.Do err:", err)
        return
    }
    fmt.Println("bodyString:", string(bodyByte))
}

func RPCPublisher(ctx context.Context) {
    // - 传递上下文 context 代替将 span 作为每个函数的第一个参数
    span, _ := opentracing.StartSpanFromContext(ctx, "RPCPublisher")
    defer span.Finish()
    println("RPCPublisher@span.LogKV>>>>", time.Now().Format("2006-01-02 15:04:05.999999999"))
    span.LogKV("RPCPublisher@LogKV:event004", "println")
    
    params := url.Values{
        "helloStr": []string{"hi prince"},
    }
    urlPath := "http://127.0.0.1:8102/publish?" + params.Encode()
    req, err := http.NewRequest(http.MethodGet, urlPath, nil)
    if err != nil {
        plog.Error("RPCPublisher http.NewRequest err:", err)
        return
    }
    bodyByte, err := httpcli.Do(req)
    if err != nil {
        plog.Error("RPCPublisher httpcli.Do err:", err)
        return
    }
    fmt.Println("bodyString:", string(bodyByte))
}

func RPCTraceFormatter(ctx context.Context) {
    // - 传递上下文 context 代替将 span 作为每个函数的第一个参数
    span, _ := opentracing.StartSpanFromContext(ctx, "RPCTraceFormatter")
    defer span.Finish()
    println("RPCTraceFormatter@span.LogKV>>>>", time.Now().Format("2006-01-02 15:04:05.999999999"))
    span.LogKV("RPCTraceFormatter@LogKV:event004", "println")
    
    params := url.Values{
        "helloTo": []string{"prince"},
    }
    urlPath := "http://127.0.0.1:8111/format?" + params.Encode()
    req, err := http.NewRequest(http.MethodGet, urlPath, nil)
    if err != nil {
        // 使用 ext.LogError 设置 Tag=error 标记 span 错误
        ext.LogError(span, err)
        plog.Error("RPCTraceFormatter http.NewRequest err:", err)
        return
    }
    
    // 当前服务与远程调用服务组成一个完整调用链追踪
    ext.SpanKindRPCClient.Set(span)
    ext.HTTPUrl.Set(span, urlPath)
    ext.HTTPMethod.Set(span, "GET")
    span.Tracer().Inject(
        span.Context(),
        opentracing.HTTPHeaders,
        opentracing.HTTPHeadersCarrier(req.Header),
    )
    
    bodyByte, err := httpcli.Do(req)
    if err != nil {
        // 使用 ext.LogError 设置 Tag=error 标记 span 错误
        ext.LogError(span, err)
        plog.Error("RPCTraceFormatter httpcli.Do err:", err)
        return
    }
    fmt.Println("bodyString:", string(bodyByte))
}

func RPCTracePublisher(ctx context.Context) {
    // - 传递上下文 context 代替将 span 作为每个函数的第一个参数
    span, _ := opentracing.StartSpanFromContext(ctx, "RPCTracePublisher")
    defer span.Finish()
    println("RPCTracePublisher@span.LogKV>>>>", time.Now().Format("2006-01-02 15:04:05.999999999"))
    span.LogKV("RPCTracePublisher@LogKV:event004", "println")
    
    params := url.Values{
        "helloStr": []string{"hi prince"},
    }
    urlPath := "http://127.0.0.1:8112/publish?" + params.Encode()
    req, err := http.NewRequest(http.MethodGet, urlPath, nil)
    if err != nil {
        // 使用 ext.LogError 设置 Tag=error 标记 span 错误
        ext.LogError(span, err)
        plog.Error("RPCTracePublisher http.NewRequest err:", err)
        return
    }
    
    // 当前服务与远程调用服务组成一个完整调用链追踪
    ext.SpanKindRPCClient.Set(span)
    ext.HTTPUrl.Set(span, urlPath)
    ext.HTTPMethod.Set(span, "GET")
    span.Tracer().Inject(
        span.Context(),
        opentracing.HTTPHeaders,
        opentracing.HTTPHeadersCarrier(req.Header),
    )
    
    bodyByte, err := httpcli.Do(req)
    if err != nil {
        // 使用 ext.LogError 设置 Tag=error 标记 span 错误
        ext.LogError(span, err)
        plog.Error("RPCTracePublisher httpcli.Do err:", err)
        return
    }
    fmt.Println("bodyString:", string(bodyByte))
}

func RPCTraceBaggageFormatter(ctx context.Context) {
    // - 传递上下文 context 代替将 span 作为每个函数的第一个参数
    span, _ := opentracing.StartSpanFromContext(ctx, "RPCTraceBaggageFormatter")
    defer span.Finish()
    println("RPCTraceBaggageFormatter@span.LogKV>>>>", time.Now().Format("2006-01-02 15:04:05.999999999"))
    span.LogKV("RPCTraceBaggageFormatter@LogKV:event004", "println")
    
    params := url.Values{
        "helloTo": []string{"prince"},
    }
    urlPath := "http://127.0.0.1:8121/format?" + params.Encode()
    req, err := http.NewRequest(http.MethodGet, urlPath, nil)
    if err != nil {
        // 使用 ext.LogError 设置 Tag=error 标记 span 错误
        ext.LogError(span, err)
        plog.Error("RPCTraceBaggageFormatter http.NewRequest err:", err)
        return
    }
    
    // 当前服务与远程调用服务组成一个完整调用链追踪
    ext.SpanKindRPCClient.Set(span)
    ext.HTTPUrl.Set(span, urlPath)
    ext.HTTPMethod.Set(span, "GET")
    span.Tracer().Inject(
        span.Context(),
        opentracing.HTTPHeaders,
        opentracing.HTTPHeadersCarrier(req.Header),
    )
    
    bodyByte, err := httpcli.Do(req)
    if err != nil {
        // 使用 ext.LogError 设置 Tag=error 标记 span 错误
        ext.LogError(span, err)
        plog.Error("RPCTraceBaggageFormatter httpcli.Do err:", err)
        return
    }
    fmt.Println("bodyString:", string(bodyByte))
}

func RPCTraceBaggagePublisher(ctx context.Context) {
    // - 传递上下文 context 代替将 span 作为每个函数的第一个参数
    span, _ := opentracing.StartSpanFromContext(ctx, "RPCTraceBaggagePublisher")
    defer span.Finish()
    println("RPCTraceBaggagePublisher@span.LogKV>>>>", time.Now().Format("2006-01-02 15:04:05.999999999"))
    span.LogKV("RPCTraceBaggagePublisher@LogKV:event004", "println")
    
    params := url.Values{
        "helloStr": []string{"hi prince"},
    }
    urlPath := "http://127.0.0.1:8122/publish?" + params.Encode()
    req, err := http.NewRequest(http.MethodGet, urlPath, nil)
    if err != nil {
        // 使用 ext.LogError 设置 Tag=error 标记 span 错误
        ext.LogError(span, err)
        plog.Error("RPCTraceBaggagePublisher http.NewRequest err:", err)
        return
    }
    
    // 当前服务与远程调用服务组成一个完整调用链追踪
    ext.SpanKindRPCClient.Set(span)
    ext.HTTPUrl.Set(span, urlPath)
    ext.HTTPMethod.Set(span, "GET")
    span.Tracer().Inject(
        span.Context(),
        opentracing.HTTPHeaders,
        opentracing.HTTPHeadersCarrier(req.Header),
    )
    
    bodyByte, err := httpcli.Do(req)
    if err != nil {
        // 使用 ext.LogError 设置 Tag=error 标记 span 错误
        ext.LogError(span, err)
        plog.Error("RPCTraceBaggagePublisher httpcli.Do err:", err)
        return
    }
    fmt.Println("bodyString:", string(bodyByte))
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
    
    opentracing.SetGlobalTracer(tracer)
    
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
    
    // opentracing.StartSpanFromContext 依赖 opentracing 的 Tracer
    opentracing.SetGlobalTracer(tracer)
    
    return tracer, closer
}

// jaegerLogger 基于 github.com/leeprince/goinfra/plog 实现 github.com/uber/jaeger-client-go@logger.go 的 Logger 接口
var jaegerLoggerPlog = &jaegerLogger{}
type jaegerLogger struct{}
func (l *jaegerLogger) Error(msg string) {
    plog.Error(msg)
}
func (l *jaegerLogger) Infof(msg string, args ...interface{}) {
    plog.Infof(msg, args...)
}
