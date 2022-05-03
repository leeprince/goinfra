package main

import (
    "fmt"
    "github.com/leeprince/goinfra/plog"
    "github.com/opentracing/opentracing-go"
    "github.com/opentracing/opentracing-go/log"
    "github.com/uber/jaeger-client-go"
    "github.com/uber/jaeger-client-go/config"
    "io"
)

const (
    serviceName = "princeJaeger-lesson01"
)

func main() {
    helloStr := fmt.Sprintf("Hello, world!")
    println(helloStr)
    
    // 注意：function opentracing.GlobalTracer() returns a no-op tracer by default.
    // tracer := opentracing.GlobalTracer()
    
    // tracer, closer := initJaeger(serviceName)
    tracer, closer := initJaegerLog(serviceName)
    defer closer.Close()
    
    println("tracer.StartSpan>>>>")
    span := tracer.StartSpan("say-hello")
    defer span.Finish()
    
    // SetTag
    println("span.SetTag>>>>")
    span.SetTag("SetTag:princeTag001", "value001")
    
    // LogFields
    println("span.LogFields>>>>")
    span.LogFields(
        log.String("LogFields:event", "string-format"),
        log.String("LogFields:value", helloStr),
    )
    // span.LogFields(log.String("LogFields:event01", "string-format"))
    // span.LogFields(log.String("LogFields:value01", helloStr))
    
    // LogKV
    println("span.LogKV>>>>")
    span.LogKV("LogKV:event", "println")
    
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
