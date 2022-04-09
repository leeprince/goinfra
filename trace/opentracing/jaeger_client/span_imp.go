package jaeger_client

import (
    "context"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/10 上午12:41
 * @Desc:   span 接口包含的方法
 */

// Baggage
func SetBaggageItem(ctx context.Context, restrictedKey, value string) {
    span := SpanFromContext(ctx)
    span.SetBaggageItem(restrictedKey, value)
}

// 记录日志的键值对
func LogKV(ctx context.Context, alternatingKeyValues ...interface{}) {
    span := SpanFromContext(ctx)
    span.LogKV(alternatingKeyValues...)
}

// 注入：获取分布式链路追踪的 span 实例，并注入（Inject）到载体（carrier）中
func Inject(ctx context.Context, format interface{}, carrier interface{}) error {
    span := SpanFromContext(ctx)
    return span.Tracer().Inject(span.Context(), format, carrier)
}
