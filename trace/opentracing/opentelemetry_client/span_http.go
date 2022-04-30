package opentelemetry_client

import (
    "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
    "net/http"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/16 下午1:27
 * @Desc:
 */

// NewHandler 将传递的处理程序包装在一个以操作命名的 span 内，并带有任何提供的选项，其功能类似于中间件。
func NewHandler(handler http.Handler, operation string, opts ...otelhttp.Option) http.Handler {
    return otelhttp.NewHandler(handler, operation, opts...)
}