package jaeger_client

import (
    "github.com/leeprince/goinfra/consts"
    "github.com/leeprince/goinfra/plog"
    "github.com/opentracing/opentracing-go"
    "io"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/9 上午10:20
 * @Desc:
 */

type pTracer struct {
    closer io.Closer
}

var ptracer pTracer

const (
    packageName = "github.com/leeprince/goinfra/trace/opentracing/jaeger_client"
)

// 初始化 Jaeger 客户端，并设置全局分布式链路追踪实例：opentracing.SetGlobalTracer
func InitTracer(serviceName string, options ...JaegerOptions) {
    jaegerOption := &jaegerOption{
        serviceName: serviceName,
        env:         consts.ENVLocal,
        isStdLogger: true,
    }
    for _, optionsFunc := range options {
        optionsFunc(jaegerOption)
    }
    
    opentracingTracer, closer, err := initTracer(jaegerOption)
    if err != nil {
        plog.Fatal("InitTracer err:", err)
    }
    
    ptracer = pTracer{closer: closer}
    
    opentracing.SetGlobalTracer(opentracingTracer)
}

// 获取全局的分布式链路追踪实例
func GlobalTracer() opentracing.Tracer {
    return opentracing.GlobalTracer()
}

// 关闭 opentracing 的所有 io
func Close() {
    ptracer.closer.Close()
}
