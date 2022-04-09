package jaeger_client

import (
    "context"
    "github.com/opentracing/opentracing-go"
    "net/http"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/10 上午1:21
 * @Desc:   内部使用 span 接口包含的方法
 */

// http 客户端远程调用服务时组成一个完整调用链追踪
func HTTPClient(ctx context.Context, method, urlPath string, header http.Header) error {
    RPCClientSetSpan(ctx)
    TagHTTPMethod(ctx, method)
    TagHTTPUrl(ctx, urlPath)
    return HTTPClientHeaderInject(ctx, header)
}

// 将 http 请求头注入到 span 实例的载体中
func HTTPClientHeaderInject(ctx context.Context, header http.Header) error {
    return Inject(
        ctx,
        opentracing.HTTPHeaders,
        opentracing.HTTPHeadersCarrier(header),
    )
}
