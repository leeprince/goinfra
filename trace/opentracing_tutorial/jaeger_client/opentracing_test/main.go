package main

import (
    "context"
    "fmt"
    "github.com/leeprince/goinfra/consts"
    "github.com/leeprince/goinfra/http/httpcli"
    "github.com/leeprince/goinfra/plog"
    "github.com/leeprince/goinfra/trace/opentracing/jaeger_client"
    "net/http"
    "net/url"
    "time"
)

const (
    serverName = "opentracing_test"
    env        = consts.EnvLocal
)

func main() {
    helloStr := fmt.Sprintf("Hello, world!")
    println(helloStr)
    
    _, closer := jaeger_client.InitJaegerTracer(serverName)
    defer closer.Close()
    
    println("tracer.StartSpan>>>>")
    ctx := context.Background()
    spanCtx := jaeger_client.StartSpan(ctx, "say-hello")
    defer jaeger_client.Finish(spanCtx)
    
    // 使用 span 的 Baggage 功能
    // after starting the span
    jaeger_client.SetBaggageItem(spanCtx, "seq", "prince-seq-202204060001")
    
    // --- 只有当前服务记录了追踪，暂未将当前服务与远程调用服务的调用链追踪连接成一个完整调用链追踪
    RPCFormatter(spanCtx)
    RPCPublisher(spanCtx)
    // --- 只有当前服务记录了追踪，暂未将当前服务与远程调用服务的调用链追踪连接成一个完整调用链追踪 -end
    
    // --- 当前服务与远程调用服务组成一个完整调用链追踪
    RPCTraceFormatter(spanCtx)
    RPCTracePublisher(spanCtx)
    // --- 当前服务与远程调用服务组成一个完整调用链追踪 -end
    
    // --- 当前服务与远程调用服务组成一个完整调用链追踪，并使用 span 的 Baggage 功能
    // 对应的课程：https://github.com/leeprince/opentracing-tutorial/tree/master/go/lesson04
    RPCTraceBaggageFormatter(spanCtx)
    RPCTraceBaggagePublisher(spanCtx)
    // --- 当前服务与远程调用服务组成一个完整调用链追踪，并使用 span 的 Baggage 功能 -end
    // 传递上下文 context 代替将 span 作为每个函数的第一个参数 -end
}

func RPCFormatter(ctx context.Context) {
    // - 传递上下文 context 代替将 span 作为每个函数的第一个参数
    spanCtx := jaeger_client.StartSpanFromContext(ctx, "RPCFormatter")
    defer jaeger_client.Finish(spanCtx)
    println("RPCFormatter@span.LogKV>>>>", time.Now().Format("2006-01-02 15:04:05.999999999"))
    jaeger_client.LogKV(spanCtx, "RPCFormatter@LogKV:event004", "println")
    
    params := url.Values{
        "helloTo": []string{"prince"},
    }
    urlPath := "http://127.0.0.1:8101/format?" + params.Encode()
    req, err := http.NewRequest(http.MethodGet, urlPath, nil)
    if err != nil {
        plog.Error("RPCFormatter http.NewRequest err:", err)
        return
    }
    
    bodyByte, err := httpcli.Do(req)
    if err != nil {
        plog.Error("RPCFormatter httpcli.Do err:", err)
        return
    }
    fmt.Println("bodyString:", string(bodyByte))
}

func RPCPublisher(ctx context.Context) {
    // - 传递上下文 context 代替将 span 作为每个函数的第一个参数
    spanCtx := jaeger_client.StartSpanFromContext(ctx, "RPCPublisher")
    defer jaeger_client.Finish(spanCtx)
    println("RPCPublisher@span.LogKV>>>>", time.Now().Format("2006-01-02 15:04:05.999999999"))
    jaeger_client.LogKV(spanCtx, "RPCPublisher@LogKV:event004", "println")
    
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
    spanCtx := jaeger_client.StartSpanFromContext(ctx, "RPCTraceFormatter")
    defer jaeger_client.Finish(spanCtx)
    println("RPCTraceFormatter@span.LogKV>>>>", time.Now().Format("2006-01-02 15:04:05.999999999"))
    jaeger_client.LogKV(spanCtx, "RPCTraceFormatter@LogKV:event004", "println")
    
    params := url.Values{
        "helloTo": []string{"prince"},
    }
    urlPath := "http://127.0.0.1:8111/format?" + params.Encode()
    req, err := http.NewRequest(http.MethodGet, urlPath, nil)
    if err != nil {
        // 使用 ext.LogError 设置 Tag=error 标记 span 错误
        jaeger_client.LogError(spanCtx, err)
        plog.Error("RPCTraceFormatter http.NewRequest err:", err)
        return
    }
    
    // 当前服务与远程调用服务组成一个完整调用链追踪
    err = jaeger_client.HTTPClient(spanCtx, http.MethodGet, urlPath, req.Header)
    if err != nil {
        // 使用 ext.LogError 设置 Tag=error 标记 span 错误
        jaeger_client.LogError(spanCtx, err)
        plog.Error("RPCTraceFormatter jaeger_client.HTTPClient err:", err)
        return
    }
    
    bodyByte, err := httpcli.Do(req)
    if err != nil {
        // 使用 ext.LogError 设置 Tag=error 标记 span 错误
        jaeger_client.LogError(spanCtx, err)
        plog.Error("RPCTraceFormatter httpcli.Do err:", err)
        return
    }
    fmt.Println("bodyString:", string(bodyByte))
}

func RPCTracePublisher(ctx context.Context) {
    // - 传递上下文 context 代替将 span 作为每个函数的第一个参数
    spanCtx := jaeger_client.StartSpanFromContext(ctx, "RPCTracePublisher")
    defer jaeger_client.Finish(spanCtx)
    println("RPCTracePublisher@span.LogKV>>>>", time.Now().Format("2006-01-02 15:04:05.999999999"))
    jaeger_client.LogKV(spanCtx, "RPCTracePublisher@LogKV:event004", "println")
    
    params := url.Values{
        "helloStr": []string{"hi prince"},
    }
    urlPath := "http://127.0.0.1:8112/publish?" + params.Encode()
    req, err := http.NewRequest(http.MethodGet, urlPath, nil)
    if err != nil {
        // 使用 ext.LogError 设置 Tag=error 标记 span 错误
        jaeger_client.LogError(spanCtx, err)
        plog.Error("RPCTracePublisher http.NewRequest err:", err)
        return
    }
    
    // 当前服务与远程调用服务组成一个完整调用链追踪
    err = jaeger_client.HTTPClient(spanCtx, http.MethodGet, urlPath, req.Header)
    if err != nil {
        // 使用 ext.LogError 设置 Tag=error 标记 span 错误
        jaeger_client.LogError(spanCtx, err)
        plog.Error("RPCTraceFormatter jaeger_client.HTTPClient err:", err)
        return
    }
    
    bodyByte, err := httpcli.Do(req)
    if err != nil {
        // 使用 ext.LogError 设置 Tag=error 标记 span 错误
        jaeger_client.LogError(spanCtx, err)
        plog.Error("RPCTracePublisher httpcli.Do err:", err)
        return
    }
    fmt.Println("bodyString:", string(bodyByte))
}

func RPCTraceBaggageFormatter(ctx context.Context) {
    // - 传递上下文 context 代替将 span 作为每个函数的第一个参数
    spanCtx := jaeger_client.StartSpanFromContext(ctx, "RPCTraceBaggageFormatter")
    defer jaeger_client.Finish(spanCtx)
    println("RPCTraceBaggageFormatter@span.LogKV>>>>", time.Now().Format("2006-01-02 15:04:05.999999999"))
    jaeger_client.LogKV(spanCtx, "RPCTraceBaggageFormatter@LogKV:event004", "println")
    
    params := url.Values{
        "helloTo": []string{"prince"},
    }
    urlPath := "http://127.0.0.1:8121/format?" + params.Encode()
    req, err := http.NewRequest(http.MethodGet, urlPath, nil)
    if err != nil {
        // 使用 ext.LogError 设置 Tag=error 标记 span 错误
        jaeger_client.LogError(spanCtx, err)
        plog.Error("RPCTraceBaggageFormatter http.NewRequest err:", err)
        return
    }
    
    // 当前服务与远程调用服务组成一个完整调用链追踪
    err = jaeger_client.HTTPClient(spanCtx, http.MethodGet, urlPath, req.Header)
    if err != nil {
        // 使用 ext.LogError 设置 Tag=error 标记 span 错误
        jaeger_client.LogError(spanCtx, err)
        plog.Error("RPCTraceFormatter jaeger_client.HTTPClient err:", err)
        return
    }
    
    bodyByte, err := httpcli.Do(req)
    if err != nil {
        // 使用 ext.LogError 设置 Tag=error 标记 span 错误
        jaeger_client.LogError(spanCtx, err)
        plog.Error("RPCTraceBaggageFormatter httpcli.Do err:", err)
        return
    }
    fmt.Println("bodyString:", string(bodyByte))
}

func RPCTraceBaggagePublisher(ctx context.Context) {
    // - 传递上下文 context 代替将 span 作为每个函数的第一个参数
    spanCtx := jaeger_client.StartSpanFromContext(ctx, "RPCTraceBaggagePublisher")
    defer jaeger_client.Finish(spanCtx)
    println("RPCTraceBaggagePublisher@span.LogKV>>>>", time.Now().Format("2006-01-02 15:04:05.999999999"))
    jaeger_client.LogKV(spanCtx, "RPCTraceBaggagePublisher@LogKV:event004", "println")
    
    params := url.Values{
        "helloStr": []string{"hi prince"},
    }
    urlPath := "http://127.0.0.1:8122/publish?" + params.Encode()
    req, err := http.NewRequest(http.MethodGet, urlPath, nil)
    if err != nil {
        // 使用 ext.LogError 设置 Tag=error 标记 span 错误
        jaeger_client.LogError(spanCtx, err)
        plog.Error("RPCTraceBaggagePublisher http.NewRequest err:", err)
        return
    }
    
    // 当前服务与远程调用服务组成一个完整调用链追踪
    err = jaeger_client.HTTPClient(spanCtx, http.MethodGet, urlPath, req.Header)
    if err != nil {
        // 使用 ext.LogError 设置 Tag=error 标记 span 错误
        jaeger_client.LogError(spanCtx, err)
        plog.Error("RPCTraceFormatter jaeger_client.HTTPClient err:", err)
        return
    }
    
    bodyByte, err := httpcli.Do(req)
    if err != nil {
        // 使用 ext.LogError 设置 Tag=error 标记 span 错误
        jaeger_client.LogError(spanCtx, err)
        plog.Error("RPCTraceBaggagePublisher httpcli.Do err:", err)
        return
    }
    fmt.Println("bodyString:", string(bodyByte))
}
