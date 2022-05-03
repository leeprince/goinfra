package opentelemetry_client

import (
    "context"
    "go.opentelemetry.io/otel/baggage"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/18 下午9:37
 * @Desc:   Baggage 的功能
 */

// 获取 Baggage: 通过 上下文获取 Baggage
func BaggageFromContext(ctx context.Context) baggage.Baggage {
    return baggage.FromContext(ctx)
}

// Baggage 上下文
//  添加 Baggage 的信息到上下文中
func ContextWithBaggage(ctx context.Context, b baggage.Baggage) context.Context {
    return baggage.ContextWithBaggage(ctx, b)
}
// 从上下文中删除 Baggage 的信息
func ContextWithoutBaggage(ctx context.Context) context.Context {
    return baggage.ContextWithoutBaggage(ctx)
}
// Baggage 上下文 -end

// 设置 Baggage 的 member
func SetBaggageItem(ctx context.Context, key, value string, props ...baggage.Property) (spanCtx context.Context, err error) {
    member, err := baggage.NewMember(key, value, props...)
    if err != nil {
        return
    }
    newBaggage, err := BaggageFromContext(ctx).SetMember(member)
    if err != nil {
        return
    }
    
    spanCtx = ContextWithBaggage(ctx, newBaggage)
    return
}

// 获取 Baggage 的 member 对应 key 的值
func BaggageItem(ctx context.Context, key string) string {
    return BaggageFromContext(ctx).Member(key).Value()
}