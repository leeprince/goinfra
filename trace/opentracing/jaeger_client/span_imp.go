package jaeger_client

import (
    "context"
    "github.com/opentracing/opentracing-go"
    "github.com/uber/jaeger-client-go"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/10 上午12:41
 * @Desc:   span 接口包含的方法
 */

// 设置 Baggage
func SetBaggageItem(ctx context.Context, restrictedKey, value string) {
    span := SpanFromContext(ctx)
    span.SetBaggageItem(restrictedKey, value)
}

// 获取 Baggage
func BaggageItem(ctx context.Context, restrictedKey string) string {
    span := SpanFromContext(ctx)
    return span.BaggageItem(restrictedKey)
}

// 记录日志的键值对
func LogKV(ctx context.Context, alternatingKeyValues ...interface{}) {
    span := SpanFromContext(ctx)
    span.LogKV(alternatingKeyValues...)
}

// 注入：获取 span 的分布式链路追踪实例（span.pTracer()），并注入（Inject）到载体（carrier）中（载体的实际类型取决于 format 的值）
func Inject(ctx context.Context, format interface{}, carrier interface{}) error {
    span := SpanFromContext(ctx)
    return span.Tracer().Inject(span.Context(), format, carrier)
}

// 提取：通过获取全局的分布式链路追踪实例提取载体（carrier）中相应的 SpanContext
func Extract(format interface{}, carrier interface{}) (opentracing.SpanContext, error) {
    return GlobalTracer().Extract(format, carrier)
}

// 获取 SpanContext
//  opentracing.SpanContext 转换成 jaeger.SpanContext
func SpanContext(ctx context.Context) jaeger.SpanContext {
    return SpanFromContext(ctx).Context().(jaeger.SpanContext)
}

// --- 标识ID：Reporting span：{初始标识ID}.{当前标识ID}.{父级标识ID}
// 当前 span 的链路追踪标识段
func SpanID(ctx context.Context) string {
    return SpanContext(ctx).SpanID().String()
}

// 当前 span 的初始链路追踪标识段，不包含`0000000000000000`，可用于标识同一个请求
func TraceID(ctx context.Context) string {
    return SpanContext(ctx).TraceID().String()
}

// 当前 span 的父级链路追踪标识段，包含`0000000000000000`
func PrentID(ctx context.Context) string {
    return SpanContext(ctx).ParentID().String()
}
// --- 标识ID -end

