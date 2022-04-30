package opentelemetry_client

import (
    "github.com/leeprince/goinfra/consts"
    sdktrace "go.opentelemetry.io/otel/sdk/trace"
    "go.opentelemetry.io/otel/trace"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/9 上午10:27
 * @Desc:
 */

type pTracer struct {
    tracer trace.Tracer
}

var ptracer pTracer

// 初始化 Telemetry 客户端
//  - exporter: exporter.go 中支持：NewIOWriterWExporter、NewJaegerExporter 作为导出器
func InitTrace(serviceName string, exporter sdktrace.SpanExporter, options ...TracerProviderOptions) {
    tracerProviderOption := &tracerProviderOption{
        serviceName:  serviceName,
        env:         consts.ENVLocal,
        exporter: exporter,
    }
    for _, optionsFunc := range options {
        optionsFunc(tracerProviderOption)
    }
    
    tracer := initTracer(tracerProviderOption)
    
    ptracer = pTracer{
        tracer: tracer,
    }
}
