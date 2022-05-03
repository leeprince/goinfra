package opentelemetry_client

import (
    "context"
    "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
    "io"
    "net/http"
    "net/url"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/16 下午1:27
 * @Desc:
 */

// http 服务端：NewHandler 将传递的处理程序包装在一个以操作命名的 span 内，并带有任何提供的选项，其功能类似于中间件。
//      http 服务端从 http 请求头载体中获取分布式链路追踪的 span 的上下文，当 http 请求头载体不存在 span 实例时创建新的 span 上下文
//      otelhttp.Handler 实现 net/http.Handler 接口（实现 ServeHTTP 方法），最终实现拦截器功能
func NewHandler(handler http.Handler, operation string, opts ...otelhttp.Option) http.Handler {
    return otelhttp.NewHandler(handler, operation, opts...)
}
// --- http 服务端 -end

// --- http 客户端:
//      默认的 otelhttp.DefaultClient 的 otelhttp.Transport 实现 http.Transport 接口（实现 RoundTrip 方法），
//      最终将客户端的分布式调用链信息注入（Inject）到 http 请求头部中，
//      而服务端再接入分布式调用链之后通过 NewHandler 实现的拦截器能够从 http 请求头中提取出分布式调用链的 span 信息，
//      最终组成完整的分布式调用链路。
func Get(ctx context.Context, url string)  (resp *http.Response, err error) {
    return otelhttp.Get(ctx, url)
}
func Head(ctx context.Context, url string)  (resp *http.Response, err error) {
    return otelhttp.Head(ctx, url)
}
func Post(ctx context.Context, url string, contentType string, body io.Reader)  (resp *http.Response, err error) {
    return otelhttp.Post(ctx, url, contentType, body)
}
func PostForm(ctx context.Context, url string, data url.Values)  (resp *http.Response, err error) {
    return otelhttp.PostForm(ctx, url, data)
}
// --- http 客户端 -end