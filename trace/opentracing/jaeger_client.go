package opentracing

import (
    "fmt"
    "github.com/leeprince/goinfra/env"
    "github.com/leeprince/goinfra/plog"
    "github.com/opentracing/opentracing-go"
    "github.com/uber/jaeger-client-go"
    "github.com/uber/jaeger-client-go/config"
    "io"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/9 上午10:20
 * @Desc:
 */

type jaegerOptions struct {
    // 服务名称
    serverName string
    
    // 环境变量
    // 生产或者沙箱环境不在控制台输出 span 的日志 LogSpans
    env string
    
    // true: 检查 logger；logger != nil 则使用，默认的 jaegerLoggerPlog。
    // false: 标准输出，输出到控制台
    isOutputFile bool
    
    // isOutputFile == true 时才检查 logger
    logger jaeger.Logger
}

type JaegerOption func(opts *jaegerOptions)

// 初始化 Jaeger 客户端
func InitJaegerTracer(serviceName string, options ...JaegerOption) (opentracing.Tracer, io.Closer) {
    jaegerOptions := &jaegerOptions{
        serverName:   serviceName,
        env:          "local",
        isOutputFile: false,
    }
    for _, optionsFunc := range options {
        optionsFunc(jaegerOptions)
    }
    
    tracer, closer, err := initJeager(jaegerOptions)
    if err != nil {
        panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
    }
    
    // opentracing.StartSpanFromContext 依赖 opentracing 的 Tracer
    opentracing.SetGlobalTracer(tracer)
    
    return tracer, closer
}

// jaegerOptions 可选参数设置
func WithJaegerOptionEnv(env string) JaegerOption {
    return func(opts *jaegerOptions) {
        opts.env = env
    }
}
func WithJaegerOptionIsOutputFile(b bool) JaegerOption {
    return func(opts *jaegerOptions) {
        opts.isOutputFile = b
    }
}
func WithJaegerOptionLogger(logger jaeger.Logger) JaegerOption {
    return func(opts *jaegerOptions) {
        opts.logger = logger
    }
}

// jaegerOptions 可选参数设置 -end

// Jaeger 客户端配置
func initJeager(opts *jaegerOptions) (opentracing.Tracer, io.Closer, error) {
    cfg := &config.Configuration{
        ServiceName: opts.serverName,
        Sampler: &config.SamplerConfig{
            Type:  "const",
            Param: 1,
        },
        Reporter: &config.ReporterConfig{
            LogSpans: true,
        },
    }
    
    // 是否在控制台输出 span 的日志
    if env.IsProdOrUat(opts.env) {
        cfg.Reporter.LogSpans = false
    }
    
    // 追踪日志的输出方式
    var logger jaeger.Logger
    logger = jaeger.StdLogger
    if !opts.isOutputFile {
        return cfg.NewTracer(config.Logger(logger))
    }
    
    logger = jaegerLoggerPlog
    if opts.logger != nil {
        logger = opts.logger
    }
    
    return cfg.NewTracer(config.Logger(logger))
}

// Jaeger 客户端配置 -end

// jaeger 记录日志的方式
// jaegerLogger 基于 github.com/leeprince/goinfra/plog 实现 github.com/uber/jaeger-client-go@logger.go 的 Logger 接口
var jaegerLoggerPlog = jaegerLogger{}

type jaegerLogger struct{}

func (l jaegerLogger) Error(msg string) {
    plog.Error(msg)
}
func (l jaegerLogger) Infof(msg string, args ...interface{}) {
    plog.Infof(msg, args...)
}

// jaeger 记录日志的方式 -end
