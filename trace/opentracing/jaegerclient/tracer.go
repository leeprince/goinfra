package jaegerclient

import (
	"github.com/leeprince/goinfra/plog"
	"github.com/leeprince/goinfra/utils"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/15 下午6:44
 * @Desc:   jaeger 实例
 */

type jaegerOption struct {
	// 服务名称
	serviceName string

	// 环境变量
	// 生产或者沙箱环境不在控制台输出 span 的日志 LogSpans
	env string

	// jaeger 内置的日志输出方式。如：reporterLogSpans（Reporting span） 日志
	// true: 标准输出，输出到 jaeger.StdLogger
	// false（默认，推荐）: false 时先检查 logger；logger != nil 则使用默认的 jaegerLoggerPlog(已实现jaeger.Logger的接口), 否则使用 输入的 logger
	isStdLogger bool

	// jaeger 内置日志是否输出 span 的日志
	reporterLogSpans bool

	// isStdLogger == true 时才检查 logger
	logger jaeger.Logger

	// localAgentHostPort 指示 reporter 将 spans 发送到此地址的 jaeger 代理
	//      默认：fmt.Sprintf("%s:%d", jaeger.DefaultUDPSpanServerHost, jaeger.DefaultUDPSpanServerPort),
	localAgentHostPort string
}

type JaegerOptions func(opt *jaegerOption)

// --- jaegerOption 可选参数设置
func WithJaegerOptionEnv(env string) JaegerOptions {
	return func(opt *jaegerOption) {
		opt.env = env
	}
}
func WithJaegerOptionIsStdLogger(b bool) JaegerOptions {
	return func(opt *jaegerOption) {
		opt.isStdLogger = b
	}
}
func WithJaegerReporterLogSpans(b bool) JaegerOptions {
	return func(opt *jaegerOption) {
		opt.reporterLogSpans = b
	}
}
func WithJaegerOptionLogger(logger jaeger.Logger) JaegerOptions {
	return func(opt *jaegerOption) {
		opt.logger = logger
	}
}
func WithJaegerLocalAgentHostPort(url string) JaegerOptions {
	return func(opt *jaegerOption) {
		opt.localAgentHostPort = url
	}
}

// --- jaegerOption 可选参数设置 -end

// --- 初始化 Tracer
// 初始化 opentracing.pTracer
func initTracer(opt *jaegerOption) (opentracing.Tracer, io.Closer, error) {
	cfg := &config.Configuration{
		ServiceName: opt.serviceName,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: time.Second * 1,
		},
	}

	if utils.IsProdOrSandbox(opt.env) || opt.reporterLogSpans == false {
		cfg.Reporter.LogSpans = false
	}

	if opt.localAgentHostPort != "" {
		cfg.Reporter.LocalAgentHostPort = opt.localAgentHostPort
	}

	// jaeger 内置的日志输出方式
	var logger jaeger.Logger
	logger = jaeger.StdLogger
	if opt.isStdLogger {
		return cfg.NewTracer(config.Logger(logger))
	}

	logger = jaegerLoggerPlog
	if opt.logger != nil {
		logger = opt.logger
	}

	return cfg.NewTracer(config.Logger(logger))
}

// --- 初始化 Tracer -end

// --- jaeger 内置的日志输出方式
// jaegerLogger 基于 github.com/leeprince/goinfra/plog 实现 github.com/uber/jaeger-client-go@logger.go 的 Logger 接口
var jaegerLoggerPlog = jaegerLogger{}

type jaegerLogger struct{}

func (l jaegerLogger) Error(msg string) {
	plog.Error(msg)
}
func (l jaegerLogger) Infof(msg string, args ...interface{}) {
	plog.Infof(msg, args...)
}

// --- jaeger 内置的日志输出方式 -end
