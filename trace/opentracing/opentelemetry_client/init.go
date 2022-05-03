package opentelemetry_client

import (
    "github.com/leeprince/goinfra/consts"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/9 上午10:27
 * @Desc:   初始化
 */

type pTracer struct {
}

var ptracer pTracer

const (
    packageName = "github.com/leeprince/goinfra/trace/opentracing/opentelemetry_client"
)

// 初始化 Telemetry 客户端
//  - exporter: exporter.go 中支持：NewIOWriterWExporter、NewJaegerExporter 作为导出器
func InitTrace(serviceName string, options ...TracerProviderOptions) {
    tracerProviderOption := &tracerProviderOption{
        serviceName: serviceName,
        env:         consts.ENVLocal,
        exporter:    nil, // 默认为 jaeger 导出器
    }
    for _, optionsFunc := range options {
        optionsFunc(tracerProviderOption)
    }
    
    initTracer(tracerProviderOption)
}
