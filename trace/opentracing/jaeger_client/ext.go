package jaeger_client

import (
    "context"
    "github.com/opentracing/opentracing-go"
    "github.com/opentracing/opentracing-go/ext"
    "github.com/opentracing/opentracing-go/log"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/10 上午12:45
 * @Desc:   额外的 span 功能
 */

// 使用 ext.LogError 设置 Tag=error 标记 span 错误
func LogError(ctx context.Context, err error, fields ...log.Field)  {
    span := SpanFromContext(ctx)
    ext.LogError(span, err, fields...)
}

// HTTP Method  tag 设置请求方法
func TagHTTPMethod(ctx context.Context, method string) {
    span := SpanFromContext(ctx)
    ext.HTTPMethod.Set(span, method)
}

// RPC 客户端 tag 设置 span
func RPCClientSetSpan(ctx context.Context)  {
    span := SpanFromContext(ctx)
    ext.SpanKindRPCClient.Set(span)
}

// --- opentracing.StartSpanOption
// RPC 服务端 option
func RPCServerOption(ctx opentracing.SpanContext) opentracing.StartSpanOption {
    return ext.RPCServerOption(ctx)
}
// --- opentracing.StartSpanOption -end

// HTTP Url tag 设置请求地址
func TagHTTPUrl(ctx context.Context, urlPath string) {
    span := SpanFromContext(ctx)
    ext.HTTPUrl.Set(span, urlPath)
}