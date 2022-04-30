package opentelemetry_client

import (
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/propagation"
    "go.opentelemetry.io/otel/trace"
    sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/15 下午6:36
 * @Desc:
 */

type tracerProviderOption struct {
    // 服务名称
    serviceName string
    
    // 环境变量
    env string
    
    // 导出器
    exporter sdktrace.SpanExporter
}

type TracerProviderOptions func(opt *tracerProviderOption)

func WithTracerProviderOptionENV(env string) TracerProviderOptions {
    return func(opt *tracerProviderOption) {
        opt.env = env
    }
}

func initTracer(option *tracerProviderOption) trace.Tracer {
    tracerProvider := sdktrace.NewTracerProvider(
        sdktrace.WithBatcher(option.exporter),
        sdktrace.WithResource(newResource(option.serviceName, option.env)),
        // sdktrace.WithSampler(sdktrace.NeverSample()),
        sdktrace.WithSampler(sdktrace.AlwaysSample()), // 采样率
    )
    otel.SetTracerProvider(tracerProvider)
    
    // 传播上下文：Propagators and Context.将全局传播器设置为 tracecontext（默认为无操作）
    otel.SetTextMapPropagator(propagation.TraceContext{})
    
    return otel.Tracer(option.serviceName)
}

// --- 获取 TracerProvider
// otel.TracerProvider 转化成
func TracerProvider() *sdktrace.TracerProvider {
    return otel.GetTracerProvider().(*sdktrace.TracerProvider)
}
// --- 获取 TracerProvider

// --- 获取 trace.Tracer
func tracer() trace.Tracer {
    return ptracer.tracer
}
// --- 获取 trace.Tracer -end





