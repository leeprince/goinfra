package main

import (
    "context"
    "github.com/leeprince/goinfra/consts"
    "github.com/leeprince/goinfra/http/httpcli"
    "github.com/leeprince/goinfra/plog"
    "github.com/leeprince/goinfra/trace/opentracing/opentelemetry_client"
    "net/http"
    "net/url"
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/17 下午12:46
 * @Desc:
 */

const (
    serviceName = "opentracing_test_app"
    env         = consts.ENVLocal
    logDir      = "./"
    logFileName = "application.log"
)

func main() {
    err := plog.SetOutputFile(logDir, logFileName, true)
    if err != nil {
        plog.Fatal("plog.SetOutputFile err:", err)
    }
    plog.SetReportCaller(true)
    
    defer func() {
        err := recover()
        if err != nil {
            plog.Trace("defer recover err:", err)
        }
    }()
    
    ctx := context.Background()
    opentelemetry_client.InitTrace(
        serviceName,
        // opentelemetry_client.WithSpanExporter(opentelemetry_client.NewIOWriterWExporter(opentelemetry_client.IOWriterExporterStdout)),
        // opentelemetry_client.WithSpanExporter(opentelemetry_client.NewIOWriterWExporter(opentelemetry_client.IOWriterExporterCreate)),
        // opentelemetry_client.WithSpanExporter(opentelemetry_client.NewIOWriterWExporter(opentelemetry_client.IOWriterExporterPlog)),
        // opentelemetry_client.WithSpanExporter(opentelemetry_client.NewJaegerExporter("")), // 默认
        opentelemetry_client.WithENV(env),
    )
    defer opentelemetry_client.Shutdown(ctx, time.Second*5)
    
    spanCtx := opentelemetry_client.StartSpan(ctx, "opentracing_test_app main")
    defer opentelemetry_client.Finish(spanCtx)
    
    // --- Baggage 功能
    spanCtx, err = opentelemetry_client.SetBaggageItem(spanCtx, "seq", "prince-seq-202204060002")
    if err != nil {
        plog.Fatal("main opentelemetry_client.SetBaggageItem err:", err)
    }
    seq := opentelemetry_client.BaggageItem(spanCtx, "seq")
    opentelemetry_client.PlogInfo(spanCtx, "opentelemetry_client.BaggageItem seq:", seq)
    // --- Baggage 功能
    
    // --- span tag
    opentelemetry_client.TagString(spanCtx, "main@TagString:001", "println")
    opentelemetry_client.TagBool(spanCtx, "main@TagBool:001", false)
    opentelemetry_client.TagBoolSlice(spanCtx, "main@TagBoolSlice:001", []bool{true, false})
    opentelemetry_client.TagBoolSlice(spanCtx, "main@TagBoolSlice:001", []bool{true, false})
    // --- span tag -end
    
    // --- span log
    opentelemetry_client.AddEvent(spanCtx, "main@AddEvent:001")
    opentelemetry_client.AddEvent(spanCtx, "main@AddEvent:002")
    // --- span log -end
    
    // --- 只有当前服务记录了追踪，暂未将当前服务与远程调用服务的调用链追踪连接成一个完整调用链追踪
    RPCFormatter(spanCtx)
    // --- 只有当前服务记录了追踪，暂未将当前服务与远程调用服务的调用链追踪连接成一个完整调用链追踪 -end
    
    // --- 当前服务与远程调用服务组成一个完整调用链追踪
    RPCTraceFormatter(spanCtx)
    // --- 当前服务与远程调用服务组成一个完整调用链追踪 -end
}

func RPCFormatter(ctx context.Context) {
    // - 传递上下文 context 代替将 span 作为每个函数的第一个参数
    spanCtx := opentelemetry_client.StartSpan(ctx, "RPCFormatter")
    defer opentelemetry_client.Finish(spanCtx)
    opentelemetry_client.TagString(spanCtx, "RPCFormatter@TagString:001", "println")
    
    params := url.Values{
        "helloTo": []string{"prince"},
    }
    urlPath := "http://127.0.0.1:8201/format?" + params.Encode()
    req, err := http.NewRequest(http.MethodGet, urlPath, nil)
    if err != nil {
        opentelemetry_client.PlogError(spanCtx, "RPCFormatter http.NewRequest err:", err)
        return
    }
    
    bodyByte, err := httpcli.Do(req)
    if err != nil {
        opentelemetry_client.PlogError(spanCtx, "RPCFormatter httpcli.Do err:", err)
        return
    }
    
    opentelemetry_client.PlogInfo(spanCtx, "bodyString:", string(bodyByte))
    opentelemetry_client.PlogInfo(spanCtx, "RPCFormatter TraceID:", opentelemetry_client.TraceID(spanCtx))
}

func RPCTraceFormatter(ctx context.Context) {
    // - 传递上下文 context 代替将 span 作为每个函数的第一个参数
    spanCtx := opentelemetry_client.StartSpan(ctx, "RPCTraceFormatter")
    defer opentelemetry_client.Finish(spanCtx)
    opentelemetry_client.TagString(spanCtx, "RPCTraceFormatter@TagString:event004", "println")
    
    params := url.Values{
        "helloTo": []string{"prince"},
    }
    urlPath := "http://127.0.0.1:8202/format?" + params.Encode()
    
    resp, err := opentelemetry_client.Get(spanCtx, urlPath)
    if err != nil {
        opentelemetry_client.PlogError(spanCtx, "RPCTraceFormatter opentelemetry_client.Get err:", err)
        return
    }
    bodyByte, err := httpcli.ResponseToBytes(resp)
    if err != nil {
        opentelemetry_client.PlogError(spanCtx, "RPCTraceFormatter httpcli.ResponseToBytes err:", err)
        return
    }
    
    opentelemetry_client.PlogInfo(spanCtx, "bodyString:", string(bodyByte))
    opentelemetry_client.PlogInfo(spanCtx, "RPCTraceFormatter TraceID:", opentelemetry_client.TraceID(spanCtx))
}
