package main

import (
    "fmt"
    "github.com/leeprince/goinfra/plog"
    "github.com/leeprince/goinfra/trace/opentracing/jaeger_client"
    "github.com/opentracing/opentracing-go"
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
    serviceName = "opentracing_test-rpc_trace-publisher"
)

func main() {
    jaeger_client.InitTracer(serviceName)
    defer jaeger_client.Close()
    
    http.HandleFunc("/publish", func(w http.ResponseWriter, r *http.Request) {
        spanCtx, err := jaeger_client.HTTPServer(r.Context(), "publisher@http.HandleFunc", r.Header)
        if err != nil {
            plog.Fatal("jaeger_client.HTTPServer err:", err)
        }
        defer jaeger_client.Finish(spanCtx)
        plog.WithFiledLogID(jaeger_client.TraceID(spanCtx)).Info("spanCtx TraceID")
        
        jaeger_client.LogKV(spanCtx, "publisher@http.HandleFunc@LogKV001", "println")
        
        helloStr := r.FormValue("helloStr")
        println(helloStr)
    })
    
    log.Fatal(http.ListenAndServe(":8112", nil))
}

// initJaeger returns an instance of Jaeger pTracer that samples 100% of traces and logs all spans to stdout.
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

// initJaeger returns an instance of Jaeger pTracer that samples 100% of traces and logs all spans to stdout.
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
    
    // opentracing.StartSpanFromContext 依赖 opentracing 的 pTracer
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
