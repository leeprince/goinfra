package main

import (
    "fmt"
    "github.com/leeprince/goinfra/plog"
    "github.com/opentracing/opentracing-go"
    "github.com/opentracing/opentracing-go/ext"
    "github.com/uber/jaeger-client-go"
    "github.com/uber/jaeger-client-go/config"
    "io"
    "log"
    "net/http"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/5 下午4:11
 * @Desc:
 */

const (
    serverName = "princeJaeger-lesson03-rpc-trace-baggage-publisher"
)

func main() {
    // 注意：function opentracing.GlobalTracer() returns a no-op tracer by default.
    // tracer := opentracing.GlobalTracer()
    
    // tracer, closer := initJaeger(serverName)
    tracer, closer := initJaegerLog(serverName)
    defer closer.Close()
    
    http.HandleFunc("/publish", func(w http.ResponseWriter, r *http.Request) {
        spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
        span := tracer.StartSpan("publisher@http.HandleFunc", ext.RPCServerOption(spanCtx))
        defer span.Finish()
        
        // 使用 span 的 Baggage 功能
        seq := span.BaggageItem("seq")
        println("BaggageItem:seq", seq)
        
        span.LogKV("publisher@http.HandleFunc@LogKV001", "println")
        
        helloStr := r.FormValue("helloStr")
        println(helloStr)
    })
    
    log.Fatal(http.ListenAndServe(":8122", nil))
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
    
    // opentracing.StartSpanFromContext 依赖 opentracing 的 Tracer
    opentracing.SetGlobalTracer(tracer)
    
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
