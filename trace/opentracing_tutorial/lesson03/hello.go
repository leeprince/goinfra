package main

import (
    "context"
    "fmt"
    "github.com/leeprince/goinfra/http/httpcli"
    "github.com/leeprince/goinfra/plog"
    "github.com/opentracing/opentracing-go"
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
    opentracing.SetGlobalTracer(tracer)
    
    defer closer.Close()
    
    println("tracer.StartSpan>>>>")
    span := tracer.StartSpan("say-hello")
    defer span.Finish()
    
    // 传递上下文 context 代替将 span 作为每个函数的第一个参数【最终方案】
    ctx := opentracing.ContextWithSpan(context.Background(), span)
    
    RPCFormatter(ctx)
    RPCPublisher(ctx)
    // 传递上下文 context 代替将 span 作为每个函数的第一个参数 -end
}

func RPCFormatter(ctx context.Context) {
    // - 传递上下文 context 代替将 span 作为每个函数的第一个参数
    span, _ := opentracing.StartSpanFromContext(ctx, "ContextSpanLogKV")
    defer span.Finish()
    
    println("ContextSpanLogKV@span.LogKV>>>>", time.Now().Format("2006-01-02 15:04:05.999999999"))
    span.LogKV("LogKV:event004", "println")
    
    httpcli.Do()
}

func RPCPublisher(ctx context.Context) {

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
    tracer, closer, err := cfg.NewTracer(config.Logger(plog.JaegerLogger))
    if err != nil {
        panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
    }
    return tracer, closer
}
