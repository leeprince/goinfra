package opentelemetry_client

import (
    "context"
    "go.opentelemetry.io/otel/trace"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/16 上午1:19
 * @Desc:   span 启动与完成
 */

// --- 获取 span
// 通过上下文 context.Context 获取 span
func SpanFromContext(ctx context.Context) trace.Span {
    return trace.SpanFromContext(ctx)
}
// --- 启动新的 span -end

// --- 启动新的 span
func StartSpan(ctx context.Context, spanName string, opts ...trace.SpanStartOption) context.Context {
    ctx, _ = Tracer().Start(ctx, spanName, opts...)
    return ctx
}
// --- 获取 span -end

// --- 完成 span
// 通过上下文 context.Context 完成 span
func Finish(ctx context.Context, options ...trace.SpanEndOption) {
    SpanFromContext(ctx).End(options...)
}
// --- 完成 span -end