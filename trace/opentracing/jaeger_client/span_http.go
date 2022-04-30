package jaeger_client

import (
    "context"
    "github.com/opentracing/opentracing-go"
    "net/http"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/10 上午1:21
 * @Desc:   使用内部 span 接口包含的方法
 */

// http 客户端远程调用服务时组成一个完整调用链追踪
func HTTPClient(ctx context.Context, method, urlPath string, header http.Header) error {
    RPCClientSetSpan(ctx)
    TagHTTPMethod(ctx, method)
    TagHTTPUrl(ctx, urlPath)
    return HTTPClientHeaderInject(ctx, header)
}

// http 客户端将 http 请求头注入到 span 实例的载体中
func HTTPClientHeaderInject(ctx context.Context, header http.Header) error {
    return Inject(
        ctx,
        opentracing.HTTPHeaders,
        opentracing.HTTPHeadersCarrier(header),
    )
}

// http 服务端从 http 请求头中获取分布式链路追踪的 span 的上下文（非 SpanContext, 而是新的 span 的上下文）
func HTTPServer(ctx context.Context, operationName string, header http.Header) (context.Context, error) {
    spanContext, err := HTTPServerHeaderExtract(header)
    if err != nil {
        return nil, err
    }
    ctx = StartSpan(ctx, operationName, RPCServerOption(spanContext))
    return ctx, nil
}

// http 服务端从 http 请求头中获取分布式链路追踪的 SpanContext
func HTTPServerHeaderExtract(header http.Header) (opentracing.SpanContext, error) {
    return Extract(
        opentracing.HTTPHeaders,
        opentracing.HTTPHeadersCarrier(header),
    )
}
