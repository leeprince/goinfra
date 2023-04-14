package jaeger_client

import (
	"context"
	"errors"
	"github.com/opentracing/opentracing-go"
	"net/http"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/10 上午1:21
 * @Desc:   使用内部 span 接口包含的方法
 */

// http 客户端: http 客户端将 span 实例注入（Inject）到 http 请求头载体中，远程服务从 http 请求头载体中提取（Extract）到 span 的实例，最终组成一个完整调用链追踪
func InjectTraceHTTPClient(ctx context.Context, url, method string, header http.Header) error {
	RPCClientSetSpan(ctx)
	TagHTTPMethod(ctx, method)
	TagHTTPURL(ctx, url)
	return InjectTraceToHTTPClientHeader(ctx, header)
}

// http 客户端将 http 请求头注入到 span 实例的载体中
func InjectTraceToHTTPClientHeader(ctx context.Context, header http.Header) error {
	return Inject(
		ctx,
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(header),
	)
}

// http 服务端：http 服务端从 http 请求头载体中获取分布式链路追踪的 span 的上下文，当 http 请求头载体不存在 span 实例时创建新的 span 上下文
func ExtractTraceHTTPServer(ctx context.Context, operationName string, header http.Header) (context.Context, error) {
	spanContext, err := ExtractHTTPServerHeader(header)
	if err != nil && !errors.Is(err, opentracing.ErrSpanContextNotFound) {
		return nil, err
	}
	ctx = StartSpan(ctx, operationName, RPCServerOption(spanContext))
	return ctx, nil
}

// http 服务端从 http 请求头中获取分布式链路追踪的 SpanContext
func ExtractHTTPServerHeader(header http.Header) (opentracing.SpanContext, error) {
	return Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(header),
	)
}
