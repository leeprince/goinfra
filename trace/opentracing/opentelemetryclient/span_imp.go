package opentelemetryclient

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/16 上午2:04
 * @Desc:   基于 span 的一些方法
 *              Attributes: 在 Jaeger UI 中对应 Tags
 *              Event: 在 Jaeger UI 中对应 Logs
 */

// 设置属性: 在 Jaeger UI 中对应 Tags
func SetAttributes(ctx context.Context, kv ...attribute.KeyValue) {
	SpanFromContext(ctx).SetAttributes(kv...)
}

// --- 基于 SetAttributes 实现具体传参. Attributes: 在 Jaeger UI 中对应 Tags
func TagBool(ctx context.Context, k string, v bool) {
	SetAttributes(ctx, attribute.Bool(k, v))
}
func TagBoolSlice(ctx context.Context, k string, v []bool) {
	SetAttributes(ctx, attribute.BoolSlice(k, v))
}
func TagInt64(ctx context.Context, k string, v int64) {
	SetAttributes(ctx, attribute.Int64(k, v))
}
func TagInt64Slice(ctx context.Context, k string, v []int64) {
	SetAttributes(ctx, attribute.Int64Slice(k, v))
}
func TagFloat64(ctx context.Context, k string, v float64) {
	SetAttributes(ctx, attribute.Float64(k, v))
}
func TagFloat64Slice(ctx context.Context, k string, v []float64) {
	SetAttributes(ctx, attribute.Float64Slice(k, v))
}
func TagString(ctx context.Context, k string, v string) {
	SetAttributes(ctx, attribute.String(k, v))
}

// --- 基于 SetAttributes 实现具体传参 -end

// 添加事件：在 Jaeger UI 中对应 Logs
//      官方示例：
//          span.AddEvent("Acquiring lock")
//          mutex.Lock()
//          span.AddEvent("Got lock, doing work...")
//          // do stuff
//          span.AddEvent("Unlocking")
//          mutex.Unlock()
func AddEvent(ctx context.Context, name string, options ...trace.EventOption) {
	SpanFromContext(ctx).AddEvent(name, options...)
}

// 添加事件：添加事件时，可选是否添加事件的属性
func WithAttributes(attributes ...attribute.KeyValue) trace.SpanStartEventOption {
	return trace.WithAttributes(attributes...)
}

// 获取 trace.SpanContext
func SpanContextFromContext(ctx context.Context) trace.SpanContext {
	return trace.SpanContextFromContext(ctx)
}

// 标记 span 错误
func LogError(ctx context.Context, err error, options ...trace.EventOption) {
	// 设置 span 状态
	SetStatus(ctx, codes.Error, err.Error())

	// RecordError 底层是 addEvent 方法， Jaeger UI 中记录在 span Logs，event=exception，但是并没有出现感叹号的错误标记，感叹号的错误标记依靠设置 span 状态（SetStatus）
	// options = append(options, trace.WithStackTrace(true)) // 跟踪错误信息
	SpanFromContext(ctx).RecordError(err, options...)
}

// 设置 span 状态：感叹号的错误标记依靠设置 span 的状态码
func SetStatus(ctx context.Context, code codes.Code, description string) {
	SpanFromContext(ctx).SetStatus(code, description)
}

// --- 标识ID
// 当前 span 的链路追踪标识段
func SpanID(ctx context.Context) string {
	return SpanContextFromContext(ctx).SpanID().String()
}

// 当前 span 的初始链路追踪标识段，不包含`0000000000000000`，可用于标识同一个请求
func TraceID(ctx context.Context) string {
	spanContext := SpanContextFromContext(ctx)
	return spanContext.TraceID().String()
}

// --- 标识ID -end
