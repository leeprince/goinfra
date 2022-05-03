package jaeger_client

import "context"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/18 下午9:36
 * @Desc:   span Baggage 的功能: 同一个分布式调用链中都可以获取到已设置的 Baggage 数据
 */

// 设置 Baggage
//  Baggage 也会在 Jaeger UI 中的 Logs 中展示(event=baggage)
func SetBaggageItem(ctx context.Context, restrictedKey, value string) {
    span := SpanFromContext(ctx)
    span.SetBaggageItem(restrictedKey, value)
}

// 获取 Baggage
func BaggageItem(ctx context.Context, restrictedKey string) string {
    span := SpanFromContext(ctx)
    return span.BaggageItem(restrictedKey)
}
