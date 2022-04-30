package opentelemetry_client

import (
    "context"
    "go.opentelemetry.io/otel/attribute"
    "go.opentelemetry.io/otel/trace"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/16 上午2:04
 * @Desc:
 */

func SetAttributes(ctx context.Context, kv ...attribute.KeyValue) {
    SpanFromContext(ctx).SetAttributes(kv...)
}

// 添加事件
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

// 事件可选是否添加自己的属性
func WithAttributes(attributes ...attribute.KeyValue) trace.SpanStartEventOption {
	return trace.WithAttributes(attributes...)
}