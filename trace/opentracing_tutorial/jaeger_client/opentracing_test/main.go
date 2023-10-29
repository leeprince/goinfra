package main

import (
	"context"
	"github.com/leeprince/goinfra/consts"
	"github.com/leeprince/goinfra/http/httpcli"
	"github.com/leeprince/goinfra/plog"
	"github.com/leeprince/goinfra/trace/opentracing/jaegerclient"
	"net/http"
	"net/url"
)

const (
	serviceName = "jaegerclient@opentracing_test"
	env         = consts.ENV_LOCAL
	logDir      = "./"
	logFileName = "application.log"

	// gd
	//reporterAgentUrl = "jaeger-agent.monitoring.svc.cluster.local:6831"
	reporterAgentUrl = "10.20.64.11:6831"
)

// 初始化日志
func initLog() {
	err := plog.SetOutputFile(logDir, logFileName, true)
	if err != nil {
		plog.Fatal("plog.SetOutputFile err:", err)
	}
	plog.SetReportCaller(true)
}

// 初始化链路跟踪。注意：程序最后一定要关闭`defer jaegerclient.Close()`
func initTrace(serviceName string) {
	jaegerclient.InitTracer(
		serviceName,
		jaegerclient.WithJaegerOptionEnv(env),
		jaegerclient.WithJaegerOptionIsStdLogger(true),
		jaegerclient.WithReporterAgentUrl(reporterAgentUrl),
		jaegerclient.WithJaegerReporterLogSpans(true))
}

func main() {
	// 初始化日志
	initLog()

	// 初始化链路跟踪
	initTrace(serviceName)
	// 注意：这里一定要关闭Trace
	defer jaegerclient.Close()

	// 开始 span 链路跟踪
	ctx := context.Background()
	spanCtx := jaegerclient.StartSpan(ctx, "main")
	// 注意：这里一定要结束 span
	defer jaegerclient.Finish(spanCtx)

	// 通过 span 继续链路跟踪
	spanTrace(spanCtx)
}

// 开始span跟踪
func spanTrace(spanCtx context.Context) {
	// span 日志
	spanTraceLog(spanCtx)

	// span logKV
	spanTraceLogKV(spanCtx)

	// span Baggage
	spanTraceBaggageItem(spanCtx)

	// span tag
	spanTraceTag(spanCtx)

	// ---

	// span 在RPC场景中的使用
	rpcApplication(spanCtx)
}

// span 日志
func spanTraceLog(spanCtx context.Context) {
	// 使用 jaegerclient.Plog(...)、 jaegerclient.Plogf(...) 代替, 或者基于这两个函数的具体函数实现
	plog.LogID(jaegerclient.TraceID(spanCtx)).Infof("spanTraceLog TraceID")

	jaegerclient.Plog(spanCtx, plog.InfoLevel, "spanTraceLog Plog")
	jaegerclient.PlogInfo(spanCtx, "spanTraceLog PlogInfo")
	jaegerclient.PlogInfof(spanCtx, "spanTraceLog PlogInfof, str:%s", "Plogf")
}

// span LogKV
func spanTraceLogKV(spanCtx context.Context) {
	jaegerclient.LogKV(spanCtx, "spanTraceLogKV:LogKV:001", "v001")
}

// span Baggage
func spanTraceBaggageItem(spanCtx context.Context) {
	jaegerclient.SetBaggageItem(spanCtx, "seq", "prince-spanTraceBaggageItem-seq-202204060001")
	seq := jaegerclient.BaggageItem(spanCtx, "seq")
	jaegerclient.PlogInfo(spanCtx, "spanTraceBaggageItem seq:", seq)
}

// span tag
func spanTraceTag(spanCtx context.Context) {
	jaegerclient.SetTag(spanCtx, "spanTraceTag:SetTag:001", "v001")
}

// span 在RPC场景中的使用
func rpcApplication(spanCtx context.Context) {
	// 只有当前服务记录了追踪，暂未将当前服务与远程调用服务的调用链追踪连接成一个完整调用链追踪
	rpcApplicationLocal(spanCtx)

	// 当前服务与远程调用服务组成一个完整调用链追踪
	rpcApplicationRemote(spanCtx)

	// 当前服务与远程调用服务组成一个完整调用链追踪，并使用 span 的 Baggage 功能
	rpcApplicationRemoteBaggage(spanCtx)
}

// 只有当前服务记录了追踪，暂未将当前服务与远程调用服务的调用链追踪连接成一个完整调用链追踪
func rpcApplicationLocal(spanCtx context.Context) {
	localTraceRPCFormatter(spanCtx)
	localTraceRPCPublisher(spanCtx)
	localTraceRPCFormatterAndRPCPublisher(spanCtx)
}

// 当前服务与远程调用服务组成一个完整调用链追踪
func rpcApplicationRemote(spanCtx context.Context) {
	localRemoteTraceRPCFormatter(spanCtx)
	localRemoteTraceRPCFormatterV1(spanCtx)
	localRemoteTracePublisher(spanCtx)
}

// 当前服务与远程调用服务组成一个完整调用链追踪，并使用 span 的 Baggage 功能
// 	- 对应的课程：https://github.com/leeprince/opentracing-tutorial/tree/master/go/lesson04
func rpcApplicationRemoteBaggage(spanCtx context.Context) {
	localRemoteBaggageTraceRPCFormatter(spanCtx)
	localRemoteBaggageTraceRPCPublisher(spanCtx)
}

func localTraceRPCFormatter(ctx context.Context) {
	jaegerclient.PlogInfo(ctx, "func localTraceRPCFormatter")

	// - 传递上下文 context 代替将 span 作为每个函数的第一个参数
	spanCtx := jaegerclient.StartSpanFromContext(ctx, "localTraceRPCFormatter")
	defer jaegerclient.Finish(spanCtx)
	jaegerclient.LogKV(spanCtx, "localTraceRPCFormatter@LogKV:event004", "println")

	jaegerclient.PlogInfo(ctx, "localTraceRPCFormatter StartSpanFromContext")

	params := url.Values{
		"helloTo": []string{"prince"},
	}
	urlPath := "http://127.0.0.1:8101/format?" + params.Encode()
	req, err := http.NewRequest(http.MethodGet, urlPath, nil)
	if err != nil {
		jaegerclient.PlogError(spanCtx, "localTraceRPCFormatter http.NewRequest err:", err)
		return
	}

	bodyByte, _, err := httpcli.Do(req)
	if err != nil {
		jaegerclient.PlogError(spanCtx, "localTraceRPCFormatter httpcli.Do err:", err)
		return
	}
	jaegerclient.PlogInfo(spanCtx, "bodyString:", string(bodyByte))

	jaegerclient.PlogInfo(spanCtx, "localTraceRPCFormatter TraceID:", jaegerclient.TraceID(spanCtx))
}

func localTraceRPCPublisher(ctx context.Context) {
	jaegerclient.PlogInfo(ctx, "func localTraceRPCPublisher")

	// - 传递上下文 context 代替将 span 作为每个函数的第一个参数
	spanCtx := jaegerclient.StartSpanFromContext(ctx, "localTraceRPCPublisher")
	defer jaegerclient.Finish(spanCtx)
	jaegerclient.LogKV(spanCtx, "localTraceRPCPublisher@LogKV:event004", "println")

	jaegerclient.PlogInfo(ctx, "localTraceRPCPublisher StartSpanFromContext")

	params := url.Values{
		"helloStr": []string{"hi prince"},
	}
	urlPath := "http://127.0.0.1:8102/publish?" + params.Encode()
	req, err := http.NewRequest(http.MethodGet, urlPath, nil)
	if err != nil {
		jaegerclient.PlogError(spanCtx, "localTraceRPCPublisher http.NewRequest err:", err)
		return
	}
	bodyByte, _, err := httpcli.Do(req)
	if err != nil {
		jaegerclient.PlogError(spanCtx, "localTraceRPCPublisher httpcli.Do err:", err)
		return
	}
	jaegerclient.PlogInfo(spanCtx, "bodyString:", string(bodyByte))

	jaegerclient.PlogInfo(spanCtx, "localTraceRPCPublisher TraceID:", jaegerclient.TraceID(spanCtx))
}

func localTraceRPCFormatterAndRPCPublisher(ctx context.Context) {
	jaegerclient.PlogInfo(ctx, "func localTraceRPCFormatter")

	// - 传递上下文 context 代替将 span 作为每个函数的第一个参数
	spanCtx := jaegerclient.StartSpanFromContext(ctx, "localTraceRPCFormatterAndRPCPublisher")
	defer jaegerclient.Finish(spanCtx)
	jaegerclient.LogKV(spanCtx, "localTraceRPCFormatterAndRPCPublisher@LogKV:event004", "println")

	jaegerclient.PlogInfo(ctx, "localTraceRPCFormatter StartSpanFromContext")

	localTraceRPCPublisher(spanCtx)
	localTraceRPCPublisher(spanCtx)

	jaegerclient.PlogInfof(spanCtx, "localTraceRPCFormatterAndRPCPublisher TraceID")
}

func localRemoteTraceRPCFormatter(ctx context.Context) {
	jaegerclient.PlogInfo(ctx, "func localRemoteTraceRPCFormatter")

	// - 传递上下文 context 代替将 span 作为每个函数的第一个参数
	spanCtx := jaegerclient.StartSpanFromContext(ctx, "localRemoteTraceRPCFormatter")
	defer jaegerclient.Finish(spanCtx)
	jaegerclient.LogKV(spanCtx, "localRemoteTraceRPCFormatter@LogKV:event004", "println")

	jaegerclient.PlogInfo(ctx, "localRemoteTraceRPCFormatter StartSpanFromContext")

	params := url.Values{
		"helloTo": []string{"prince"},
	}
	urlPath := "http://127.0.0.1:8111/format?" + params.Encode()
	req, err := http.NewRequest(http.MethodGet, urlPath, nil)
	if err != nil {
		jaegerclient.PlogError(spanCtx, "localRemoteTraceRPCFormatter http.NewRequest err:", err)
		return
	}

	// 当前服务与远程调用服务组成一个完整调用链追踪
	err = jaegerclient.InjectTraceHTTPClient(spanCtx, urlPath, http.MethodGet, req.Header)
	if err != nil {
		jaegerclient.PlogError(spanCtx, "localRemoteTraceRPCFormatter jaegerclient.InjectTraceHTTPClient err:", err)
		return
	}
	jaegerclient.PlogInfo(ctx, "localRemoteTraceRPCFormatter req:", req)

	bodyByte, _, err := httpcli.Do(req)
	if err != nil {
		jaegerclient.PlogError(spanCtx, "localRemoteTraceRPCFormatter httpcli.Do err:", err)
		return
	}
	jaegerclient.PlogInfo(spanCtx, "bodyString:", string(bodyByte))

	jaegerclient.PlogInfo(spanCtx, "localRemoteTraceRPCFormatter TraceID:", jaegerclient.TraceID(spanCtx))
}

// 在 localRemoteTraceRPCFormatter 的基础上，将http请求部分的直接通过已封装的 httpcli 发起请求
func localRemoteTraceRPCFormatterV1(ctx context.Context) {
	jaegerclient.PlogInfo(ctx, "func localRemoteTraceRPCFormatterV1")

	// - 传递上下文 context 代替将 span 作为每个函数的第一个参数
	spanCtx := jaegerclient.StartSpanFromContext(ctx, "localRemoteTraceRPCFormatterV1")
	defer jaegerclient.Finish(spanCtx)
	jaegerclient.LogKV(spanCtx, "localRemoteTraceRPCFormatterV1@LogKV:event004", "println")

	jaegerclient.PlogInfo(ctx, "localRemoteTraceRPCFormatterV1 StartSpanFromContext")

	params := url.Values{
		"helloTo": []string{"prince"},
	}
	hosturl := "http://127.0.0.1:8111/format?" + params.Encode()
	httpClient := httpcli.NewHttpClient().
		//WithIsHttpTrace(true). // logID=重新生成的
		//WithIsHttpTrace(true, ctx). // logID=调用链ID
		WithIsHttpTrace(true, spanCtx). // logID=调用链ID
		WithLogID(jaegerclient.TraceID(spanCtx)).
		WithMethod(http.MethodGet).
		WithURL(hosturl)
	bodyByte, _, err := httpClient.Do()
	if err != nil {
		jaegerclient.PlogError(spanCtx, "localRemoteTraceRPCFormatterV1 httpcli.Do err:", err)
		return
	}
	jaegerclient.PlogInfo(spanCtx, "bodyString:", string(bodyByte))

	jaegerclient.PlogInfo(spanCtx, "localRemoteTraceRPCFormatterV1 TraceID:", jaegerclient.TraceID(spanCtx))
}

func localRemoteTracePublisher(ctx context.Context) {
	jaegerclient.PlogInfo(ctx, "func localRemoteTracePublisher")

	// - 传递上下文 context 代替将 span 作为每个函数的第一个参数
	spanCtx := jaegerclient.StartSpanFromContext(ctx, "localRemoteTracePublisher")
	defer jaegerclient.Finish(spanCtx)
	jaegerclient.LogKV(spanCtx, "localRemoteTracePublisher@LogKV:event004", "println")

	jaegerclient.PlogInfo(ctx, "localRemoteTracePublisher StartSpanFromContext")

	params := url.Values{
		"helloStr": []string{"hi prince"},
	}
	urlPath := "http://127.0.0.1:8112/publish?" + params.Encode()
	req, err := http.NewRequest(http.MethodGet, urlPath, nil)
	if err != nil {
		jaegerclient.PlogError(spanCtx, "localRemoteTracePublisher http.NewRequest err:", err)
		return
	}

	// 当前服务与远程调用服务组成一个完整调用链追踪
	err = jaegerclient.InjectTraceHTTPClient(spanCtx, urlPath, http.MethodGet, req.Header)
	if err != nil {
		jaegerclient.PlogError(spanCtx, "localRemoteTraceRPCFormatter jaegerclient.InjectTraceHTTPClient err:", err)
		return
	}

	bodyByte, _, err := httpcli.Do(req)
	if err != nil {
		jaegerclient.PlogError(spanCtx, "localRemoteTracePublisher httpcli.Do err:", err)
		return
	}
	jaegerclient.PlogInfo(spanCtx, "bodyString:", string(bodyByte))

	jaegerclient.PlogInfo(spanCtx, "localRemoteTracePublisher TraceID:", jaegerclient.TraceID(spanCtx))
}

func localRemoteBaggageTraceRPCFormatter(ctx context.Context) {
	jaegerclient.PlogInfo(ctx, "func localRemoteBaggageTraceRPCFormatter")

	// - 传递上下文 context 代替将 span 作为每个函数的第一个参数
	spanCtx := jaegerclient.StartSpanFromContext(ctx, "localRemoteBaggageTraceRPCFormatter")
	defer jaegerclient.Finish(spanCtx)
	jaegerclient.LogKV(spanCtx, "localRemoteBaggageTraceRPCFormatter@LogKV:event004", "println")

	jaegerclient.PlogInfo(ctx, "localRemoteBaggageTraceRPCFormatter StartSpanFromContext")

	params := url.Values{
		"helloTo": []string{"prince"},
	}
	urlPath := "http://127.0.0.1:8121/format?" + params.Encode()
	req, err := http.NewRequest(http.MethodGet, urlPath, nil)
	if err != nil {
		jaegerclient.PlogError(spanCtx, "localRemoteBaggageTraceRPCFormatter http.NewRequest err:", err)
		return
	}

	// 当前服务与远程调用服务组成一个完整调用链追踪
	err = jaegerclient.InjectTraceHTTPClient(spanCtx, urlPath, http.MethodGet, req.Header)
	if err != nil {
		jaegerclient.PlogError(spanCtx, "localRemoteTraceRPCFormatter jaegerclient.InjectTraceHTTPClient err:", err)
		return
	}

	bodyByte, _, err := httpcli.Do(req)
	if err != nil {
		jaegerclient.PlogError(spanCtx, "localRemoteBaggageTraceRPCFormatter httpcli.Do err:", err)
		return
	}
	jaegerclient.PlogInfo(spanCtx, "bodyString:", string(bodyByte))

	jaegerclient.PlogInfo(spanCtx, "localRemoteBaggageTraceRPCFormatter TraceID:", jaegerclient.TraceID(spanCtx))
}

func localRemoteBaggageTraceRPCPublisher(ctx context.Context) {
	jaegerclient.PlogInfo(ctx, "func localRemoteBaggageTraceRPCPublisher")

	// - 传递上下文 context 代替将 span 作为每个函数的第一个参数
	spanCtx := jaegerclient.StartSpanFromContext(ctx, "localRemoteBaggageTraceRPCPublisher")
	defer jaegerclient.Finish(spanCtx)
	jaegerclient.LogKV(spanCtx, "localRemoteBaggageTraceRPCPublisher@LogKV:event004", "println")

	jaegerclient.PlogInfo(ctx, "localRemoteBaggageTraceRPCPublisher StartSpanFromContext")

	params := url.Values{
		"helloStr": []string{"hi prince"},
	}
	urlPath := "http://127.0.0.1:8122/publish?" + params.Encode()
	req, err := http.NewRequest(http.MethodGet, urlPath, nil)
	if err != nil {
		jaegerclient.PlogError(spanCtx, "localRemoteBaggageTraceRPCPublisher http.NewRequest err:", err)
		return
	}

	// 当前服务与远程调用服务组成一个完整调用链追踪
	err = jaegerclient.InjectTraceHTTPClient(spanCtx, urlPath, http.MethodGet, req.Header)
	if err != nil {
		jaegerclient.PlogError(spanCtx, "localRemoteTraceRPCFormatter jaegerclient.InjectTraceHTTPClient err:", err)
		return
	}

	bodyByte, _, err := httpcli.Do(req)
	if err != nil {
		jaegerclient.PlogError(spanCtx, "localRemoteBaggageTraceRPCPublisher httpcli.Do err:", err)
		return
	}
	jaegerclient.PlogInfo(spanCtx, "bodyString:", string(bodyByte))

	jaegerclient.PlogInfo(spanCtx, "localRemoteBaggageTraceRPCPublisher TraceID:", jaegerclient.TraceID(spanCtx))
}
