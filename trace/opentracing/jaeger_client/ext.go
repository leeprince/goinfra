package jaeger_client

import (
    "context"
    "github.com/opentracing/opentracing-go/ext"
    "github.com/opentracing/opentracing-go/log"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/10 上午12:45
 * @Desc:
 */

// --- 额外的 span 功能

// 记录 span 存在的错误
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

// HTTP Url tag 设置请求地址
func TagHTTPUrl(ctx context.Context, urlPath string) {
    span := SpanFromContext(ctx)
    ext.HTTPUrl.Set(span, urlPath)
}

// --- 额外的 span 功能 -end
