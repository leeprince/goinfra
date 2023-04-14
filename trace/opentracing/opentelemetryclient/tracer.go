package opentelemetryclient

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/15 下午6:36
 * @Desc:   Tracer
 */

const (
	// 默认的 trace 名称（可不设置）
	//  底层 go.opentelemetry.io/otel/sdk 也存在一个默认 `defaultTracerName=go.opentelemetry.io/otel/sdk/tracer`
	defaultTracerName = "github.com/leeprince/goinfra/trace"
)

type tracerProviderOption struct {
	// 服务名称
	serviceName string

	// 环境变量
	env string

	// 导出器
	exporter sdktrace.SpanExporter
}

type TracerProviderOptions func(opt *tracerProviderOption)

func WithENV(env string) TracerProviderOptions {
	return func(opt *tracerProviderOption) {
		opt.env = env
	}
}

func WithSpanExporter(exporter sdktrace.SpanExporter) TracerProviderOptions {
	return func(opt *tracerProviderOption) {
		opt.exporter = exporter
	}
}

func initTracer(option *tracerProviderOption) {
	exporter := option.exporter
	if exporter == nil {
		exporter = NewJaegerExporter("")
	}

	resource, _ := newResource(option.serviceName, option.env)

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource),
		// sdktrace.WithSampler(sdktrace.NeverSample()),
		sdktrace.WithSampler(sdktrace.AlwaysSample()), // 采样率
	)
	otel.SetTracerProvider(tracerProvider)

	// 传播上下文：Propagators and Context.将全局传播器设置为 tracecontext（默认为无操作）
	otel.SetTextMapPropagator(
		// propagation.TraceContext{}, // 使用分布式链路追踪的 Trace Span 功能（保证分布式传播 Trace Span：TraceID 统一 ）
		// propagation.Baggage{}, // 使用分布式链路追踪的分布式 Baggate 功能（保证分布式传播 Baggate 功能）
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)
}

// --- 获取 TracerProvider
// otel.TracerProvider 转化成
func TracerProvider() *sdktrace.TracerProvider {
	return otel.GetTracerProvider().(*sdktrace.TracerProvider)
}

// --- 获取 TracerProvider

// --- 获取 trace.Tracer
func Tracer() trace.Tracer {
	return TracerProvider().Tracer(defaultTracerName)
}

// --- 获取 trace.Tracer -end
