package main

import (
	"context"
	"github.com/leeprince/goinfra/consts"
	"github.com/leeprince/goinfra/http/httpcli"
	"github.com/leeprince/goinfra/plog"
	"github.com/leeprince/goinfra/trace/opentracing/jaeger_client"
	"net/http"
	"net/url"
)

const (
	serviceName = "jaeger_client@opentracing_test"
	env         = consts.ENVLocal
	logDir      = "./"
	logFileName = "application.log"
)

func main() {
	err := plog.SetOutputFile(logDir, logFileName, true)
	if err != nil {
		plog.Fatal("plog.SetOutputFile err:", err)
	}
	plog.SetReportCaller(true)

	jaeger_client.InitTracer(
		serviceName,
		jaeger_client.WithJaegerOptionEnv(env),
		jaeger_client.WithJaegerOptionIsStdLogger(true),
		jaeger_client.WithJaegerReporterLogSpans(true))
	defer jaeger_client.Close()

	ctx := context.Background()
	spanCtx := jaeger_client.StartSpan(ctx, "jaeger_client@opentracing_test")
	defer jaeger_client.Finish(spanCtx)

	// --- 使用 span 的 Baggage 功能
	jaeger_client.SetBaggageItem(spanCtx, "seq", "prince-seq-202204060001")
	seq := jaeger_client.BaggageItem(spanCtx, "seq")
	jaeger_client.PlogInfo(spanCtx, "main jaeger_client.BaggageItem seq:", seq)
	// --- 使用 span 的 Baggage 功能 -end

	// --- span log
	jaeger_client.LogKV(spanCtx, "main:LogKV:001", "v001")
	// --- span tag -end

	// --- span tag
	jaeger_client.SetTag(spanCtx, "main:SetTag:001", "v001")
	// --- span tag -end

	// --- 日志
	// 使用 jaeger_client.Plog(...)、 jaeger_client.Plogf(...) 代替, 或者基于这两个函数的具体函数实现
	plog.LogID(jaeger_client.TraceID(spanCtx)).Infof("main TraceID")

	jaeger_client.Plog(spanCtx, plog.InfoLevel, "main Plog")
	jaeger_client.PlogInfo(spanCtx, "main PlogInfo")
	jaeger_client.PlogInfof(spanCtx, "main PlogInfof, str:%s", "Plogf")
	// --- 日志 -end

	// --- 只有当前服务记录了追踪，暂未将当前服务与远程调用服务的调用链追踪连接成一个完整调用链追踪
	RPCFormatter(spanCtx)
	RPCPublisher(spanCtx)
	RPCFormatterAndRPCPublisher(spanCtx)
	// --- 只有当前服务记录了追踪，暂未将当前服务与远程调用服务的调用链追踪连接成一个完整调用链追踪 -end

	// --- 当前服务与远程调用服务组成一个完整调用链追踪
	RPCTraceFormatter(spanCtx)
	RPCTracePublisher(spanCtx)
	// --- 当前服务与远程调用服务组成一个完整调用链追踪 -end

	// --- 当前服务与远程调用服务组成一个完整调用链追踪，并使用 span 的 Baggage 功能
	// 对应的课程：https://github.com/leeprince/opentracing-tutorial/tree/master/go/lesson04
	RPCTraceBaggageFormatter(spanCtx)
	RPCTraceBaggagePublisher(spanCtx)
	// --- 当前服务与远程调用服务组成一个完整调用链追踪，并使用 span 的 Baggage 功能 -end
	// 传递上下文 context 代替将 span 作为每个函数的第一个参数 -end

}

func RPCFormatter(ctx context.Context) {
	// - 传递上下文 context 代替将 span 作为每个函数的第一个参数
	spanCtx := jaeger_client.StartSpanFromContext(ctx, "RPCFormatter")
	defer jaeger_client.Finish(spanCtx)
	jaeger_client.LogKV(spanCtx, "RPCFormatter@LogKV:event004", "println")

	params := url.Values{
		"helloTo": []string{"prince"},
	}
	urlPath := "http://127.0.0.1:8101/format?" + params.Encode()
	req, err := http.NewRequest(http.MethodGet, urlPath, nil)
	if err != nil {
		jaeger_client.PlogError(spanCtx, "RPCFormatter http.NewRequest err:", err)
		return
	}

	bodyByte, err := httpcli.Do(req)
	if err != nil {
		jaeger_client.PlogError(spanCtx, "RPCFormatter httpcli.Do err:", err)
		return
	}
	jaeger_client.PlogInfo(spanCtx, "bodyString:", string(bodyByte))

	jaeger_client.PlogInfo(spanCtx, "RPCFormatter TraceID:", jaeger_client.TraceID(spanCtx))
}

func RPCPublisher(ctx context.Context) {
	// - 传递上下文 context 代替将 span 作为每个函数的第一个参数
	spanCtx := jaeger_client.StartSpanFromContext(ctx, "RPCPublisher")
	defer jaeger_client.Finish(spanCtx)
	jaeger_client.LogKV(spanCtx, "RPCPublisher@LogKV:event004", "println")

	params := url.Values{
		"helloStr": []string{"hi prince"},
	}
	urlPath := "http://127.0.0.1:8102/publish?" + params.Encode()
	req, err := http.NewRequest(http.MethodGet, urlPath, nil)
	if err != nil {
		jaeger_client.PlogError(spanCtx, "RPCPublisher http.NewRequest err:", err)
		return
	}
	bodyByte, err := httpcli.Do(req)
	if err != nil {
		jaeger_client.PlogError(spanCtx, "RPCPublisher httpcli.Do err:", err)
		return
	}
	jaeger_client.PlogInfo(spanCtx, "bodyString:", string(bodyByte))

	jaeger_client.PlogInfo(spanCtx, "RPCPublisher TraceID:", jaeger_client.TraceID(spanCtx))
}

func RPCFormatterAndRPCPublisher(ctx context.Context) {
	// - 传递上下文 context 代替将 span 作为每个函数的第一个参数
	spanCtx := jaeger_client.StartSpanFromContext(ctx, "RPCFormatterAndRPCPublisher")
	defer jaeger_client.Finish(spanCtx)
	jaeger_client.LogKV(spanCtx, "RPCFormatterAndRPCPublisher@LogKV:event004", "println")

	RPCPublisher(spanCtx)
	RPCPublisher(spanCtx)

	jaeger_client.PlogInfof(spanCtx, "RPCFormatterAndRPCPublisher TraceID")
}

func RPCTraceFormatter(ctx context.Context) {
	// - 传递上下文 context 代替将 span 作为每个函数的第一个参数
	spanCtx := jaeger_client.StartSpanFromContext(ctx, "RPCTraceFormatter")
	defer jaeger_client.Finish(spanCtx)
	jaeger_client.LogKV(spanCtx, "RPCTraceFormatter@LogKV:event004", "println")

	params := url.Values{
		"helloTo": []string{"prince"},
	}
	urlPath := "http://127.0.0.1:8111/format?" + params.Encode()
	req, err := http.NewRequest(http.MethodGet, urlPath, nil)
	if err != nil {
		jaeger_client.PlogError(spanCtx, "RPCTraceFormatter http.NewRequest err:", err)
		return
	}

	// 当前服务与远程调用服务组成一个完整调用链追踪
	err = jaeger_client.InjectTraceHTTPClient(spanCtx, urlPath, http.MethodGet, req.Header)
	if err != nil {
		jaeger_client.PlogError(spanCtx, "RPCTraceFormatter jaeger_client.InjectTraceHTTPClient err:", err)
		return
	}

	bodyByte, err := httpcli.Do(req)
	if err != nil {
		jaeger_client.PlogError(spanCtx, "RPCTraceFormatter httpcli.Do err:", err)
		return
	}
	jaeger_client.PlogInfo(spanCtx, "bodyString:", string(bodyByte))

	jaeger_client.PlogInfo(spanCtx, "RPCTraceFormatter TraceID:", jaeger_client.TraceID(spanCtx))
}

func RPCTracePublisher(ctx context.Context) {
	// - 传递上下文 context 代替将 span 作为每个函数的第一个参数
	spanCtx := jaeger_client.StartSpanFromContext(ctx, "RPCTracePublisher")
	defer jaeger_client.Finish(spanCtx)
	jaeger_client.LogKV(spanCtx, "RPCTracePublisher@LogKV:event004", "println")

	params := url.Values{
		"helloStr": []string{"hi prince"},
	}
	urlPath := "http://127.0.0.1:8112/publish?" + params.Encode()
	req, err := http.NewRequest(http.MethodGet, urlPath, nil)
	if err != nil {
		jaeger_client.PlogError(spanCtx, "RPCTracePublisher http.NewRequest err:", err)
		return
	}

	// 当前服务与远程调用服务组成一个完整调用链追踪
	err = jaeger_client.InjectTraceHTTPClient(spanCtx, urlPath, http.MethodGet, req.Header)
	if err != nil {
		jaeger_client.PlogError(spanCtx, "RPCTraceFormatter jaeger_client.InjectTraceHTTPClient err:", err)
		return
	}

	bodyByte, err := httpcli.Do(req)
	if err != nil {
		jaeger_client.PlogError(spanCtx, "RPCTracePublisher httpcli.Do err:", err)
		return
	}
	jaeger_client.PlogInfo(spanCtx, "bodyString:", string(bodyByte))

	jaeger_client.PlogInfo(spanCtx, "RPCTracePublisher TraceID:", jaeger_client.TraceID(spanCtx))
}

func RPCTraceBaggageFormatter(ctx context.Context) {
	// - 传递上下文 context 代替将 span 作为每个函数的第一个参数
	spanCtx := jaeger_client.StartSpanFromContext(ctx, "RPCTraceBaggageFormatter")
	defer jaeger_client.Finish(spanCtx)
	jaeger_client.LogKV(spanCtx, "RPCTraceBaggageFormatter@LogKV:event004", "println")

	params := url.Values{
		"helloTo": []string{"prince"},
	}
	urlPath := "http://127.0.0.1:8121/format?" + params.Encode()
	req, err := http.NewRequest(http.MethodGet, urlPath, nil)
	if err != nil {
		jaeger_client.PlogError(spanCtx, "RPCTraceBaggageFormatter http.NewRequest err:", err)
		return
	}

	// 当前服务与远程调用服务组成一个完整调用链追踪
	err = jaeger_client.InjectTraceHTTPClient(spanCtx, urlPath, http.MethodGet, req.Header)
	if err != nil {
		jaeger_client.PlogError(spanCtx, "RPCTraceFormatter jaeger_client.InjectTraceHTTPClient err:", err)
		return
	}

	bodyByte, err := httpcli.Do(req)
	if err != nil {
		jaeger_client.PlogError(spanCtx, "RPCTraceBaggageFormatter httpcli.Do err:", err)
		return
	}
	jaeger_client.PlogInfo(spanCtx, "bodyString:", string(bodyByte))

	jaeger_client.PlogInfo(spanCtx, "RPCTraceBaggageFormatter TraceID:", jaeger_client.TraceID(spanCtx))
}

func RPCTraceBaggagePublisher(ctx context.Context) {
	// - 传递上下文 context 代替将 span 作为每个函数的第一个参数
	spanCtx := jaeger_client.StartSpanFromContext(ctx, "RPCTraceBaggagePublisher")
	defer jaeger_client.Finish(spanCtx)
	jaeger_client.LogKV(spanCtx, "RPCTraceBaggagePublisher@LogKV:event004", "println")

	params := url.Values{
		"helloStr": []string{"hi prince"},
	}
	urlPath := "http://127.0.0.1:8122/publish?" + params.Encode()
	req, err := http.NewRequest(http.MethodGet, urlPath, nil)
	if err != nil {
		jaeger_client.PlogError(spanCtx, "RPCTraceBaggagePublisher http.NewRequest err:", err)
		return
	}

	// 当前服务与远程调用服务组成一个完整调用链追踪
	err = jaeger_client.InjectTraceHTTPClient(spanCtx, urlPath, http.MethodGet, req.Header)
	if err != nil {
		jaeger_client.PlogError(spanCtx, "RPCTraceFormatter jaeger_client.InjectTraceHTTPClient err:", err)
		return
	}

	bodyByte, err := httpcli.Do(req)
	if err != nil {
		jaeger_client.PlogError(spanCtx, "RPCTraceBaggagePublisher httpcli.Do err:", err)
		return
	}
	jaeger_client.PlogInfo(spanCtx, "bodyString:", string(bodyByte))

	jaeger_client.PlogInfo(spanCtx, "RPCTraceBaggagePublisher TraceID:", jaeger_client.TraceID(spanCtx))
}
