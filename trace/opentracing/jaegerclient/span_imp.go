package jaegerclient

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

// 设置 span 日志: 必须是键值对(key, value) 的形式记录日志，即 alternatingKeyValues 的参数格式必须是2的倍数
func LogKV(ctx context.Context, alternatingKeyValues ...interface{}) {
	span := SpanFromContext(ctx)
	span.LogKV(alternatingKeyValues...)
}

// 设置 span Tag
func SetTag(ctx context.Context, key string, value interface{}) {
	span := SpanFromContext(ctx)
	span.SetTag(key, value)
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

// --- 标识ID：Reporting span：{TraceID初始标识ID}.{SpanID当前标识ID}.{PrentID父级标识ID}
// traceID表示跟踪的全局唯一ID。通常作为随机数生成，不包含`0000000000000000`，可用于标识同一个请求
func TraceID(ctx context.Context) string {
	return SpanContext(ctx).TraceID().String()
}

// spanID表示在其跟踪中必须是唯一的span ID，但不必是全局唯一的。
func SpanID(ctx context.Context) string {
	return SpanContext(ctx).SpanID().String()
}

// parentID指的是父跨度的ID。如果当前跨度是根跨度，则应为0，包含`0000000000000000`
func PrentID(ctx context.Context) string {
	return SpanContext(ctx).ParentID().String()
}

// --- 标识ID：Reporting span：{TraceID初始标识ID}.{SpanID当前标识ID}.{PrentID父级标识ID}-end
