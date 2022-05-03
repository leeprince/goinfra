package jaeger_client

import (
    "context"
    "github.com/opentracing/opentracing-go"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/10 上午12:01
 * @Desc:   span 启动与完成
 */

// --- 开始新的 span
// 开始新的初始 span，并设置 span 到上下文 context.Context 中
//      ctx：一般是 context.Background()、context.TODO() 或者其他非分布式调用链的上下文
func StartSpan(ctx context.Context, operationName string, opts ...opentracing.StartSpanOption) context.Context {
    span := opentracing.GlobalTracer().StartSpan(operationName, opts...)
    
    // 传递上下文 context 代替将 span 作为每个函数的第一个参数【最终方案】
    ctx = opentracing.ContextWithSpan(ctx, span)
    
    return ctx
}

// 开始新的子 span，通过上下文 context.Context 开始新的子 span（对应的初始 span 是一样的）
func StartSpanFromContext(ctx context.Context, operationName string, opts ...opentracing.StartSpanOption) context.Context {
    _, ctx = opentracing.StartSpanFromContext(ctx, operationName, opts...)
    return ctx
}

// --- 开始新的 span -end

// --- 获取 span
// 通过上下文 context.Context 获取 span
func SpanFromContext(ctx context.Context) opentracing.Span {
    return opentracing.SpanFromContext(ctx)
}
// --- 获取 span -end

// --- 完成 span
// 通过上下文 context.Context 完成 span
func Finish(ctx context.Context) {
    opentracing.SpanFromContext(ctx).Finish()
}
// --- 完成 span -end

