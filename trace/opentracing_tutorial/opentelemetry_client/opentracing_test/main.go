package main

import (
	"context"
	"fmt"
	"github.com/leeprince/goinfra/consts"
	"github.com/leeprince/goinfra/plog"
	"github.com/leeprince/goinfra/trace/opentracing/opentelemetryclient"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"sync"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/15 下午5:30
 * @Desc:
 */

const (
	serviceName = "opentracing_test"
	env         = consts.ENVLocal
	logDir      = "./"
	logFileName = "application.log"
	port        = ":8084"
)

func main() {
	err := plog.SetOutputFile(logDir, logFileName, true)
	if err != nil {
		plog.Fatal("plog.SetOutputFile err:", err)
	}
	plog.SetReportCaller(true)

	ctx := context.Background()
	opentelemetryclient.InitTrace(
		serviceName,
		// opentelemetry_client.WithSpanExporter(opentelemetry_client.NewIOWriterWExporter(opentelemetry_client.IOWriterExporterStdout)),
		// opentelemetry_client.WithSpanExporter(opentelemetry_client.NewIOWriterWExporter(opentelemetry_client.IOWriterExporterCreate)),
		// opentelemetry_client.WithSpanExporter(opentelemetry_client.NewIOWriterWExporter(opentelemetry_client.IOWriterExporterPlog)),
		// opentelemetry_client.WithSpanExporter(opentelemetry_client.NewJaegerExporter("")), // 默认
		opentelemetryclient.WithENV(env),
	)
	defer opentelemetryclient.Shutdown(ctx, time.Second*5)

	// manualHttpHandler 加入 opentelemetry_client.NewHandler 拦截器中（中间件）
	handler := http.HandlerFunc(manualHttpHandler)
	wrappedHandler := opentelemetryclient.NewHandler(handler, "main:NewHandler")
	http.Handle("/hello-export", wrappedHandler)

	// And start the HTTP serve.
	plog.Info("port:", port)
	plog.Fatal(http.ListenAndServe(port, nil))
}

// myHttpHandler is an HTTP handler function that is going to be instrumented.
func manualHttpHandler(w http.ResponseWriter, r *http.Request) {
	spanCtx := opentelemetryclient.StartSpan(r.Context(), "httpHandler")
	defer opentelemetryclient.Finish(spanCtx)

	fmt.Fprintf(w, "Hello, World! I am instrumented automatically!")
	parentFunction(spanCtx)

	// Attribute keys can be precomputed
	var myKey = attribute.Key("httpHandler-SetAttributes")
	opentelemetryclient.SetAttributes(spanCtx, myKey.String("a value"))

	// --- Events
	mutex := sync.Mutex{}
	opentelemetryclient.AddEvent(spanCtx, "Acquiring lock")
	mutex.Lock()
	opentelemetryclient.AddEvent(spanCtx, "Got lock, doing work...")
	// do stuff
	opentelemetryclient.AddEvent(spanCtx, "Unlocking")
	mutex.Unlock()

	// Events can also have attributes of their own
	opentelemetryclient.AddEvent(spanCtx, "Cancelled wait due to external signal", opentelemetryclient.WithAttributes(attribute.Int("pid", 4328)))

	// --- 日志
	// 使用 jaeger_client.Plog(...)、 jaeger_client.Plogf(...) 代替, 或者基于这两个函数的具体函数实现
	plog.LogID(opentelemetryclient.TraceID(spanCtx)).Infof("main TraceID")

	opentelemetryclient.Plog(spanCtx, plog.InfoLevel, "4export_data:manualHttpHandler:end Plog")
	opentelemetryclient.PlogInfo(spanCtx, "4export_data:manualHttpHandler:end PlogInfo")
	opentelemetryclient.PlogInfof(spanCtx, "4export_data:manualHttpHandler:end PlogInfof:%s", "PlogInfof")
	// --- 日志 -end
}

func parentFunction(ctx context.Context) {
	spanCtx := opentelemetryclient.StartSpan(ctx, "parent")
	defer opentelemetryclient.Finish(spanCtx)

	// call the child function and start a nested span in there
	childFunction(spanCtx)

	// do more work - when this function ends, parentSpan will complete.
	parentFunctionAttributes(spanCtx)
}

func parentFunctionAttributes(ctx context.Context) {
	// setting attributes at creation...
	spanCtx := opentelemetryclient.StartSpan(ctx, "parentFunctionAttributes-attributesAtCreation", trace.WithAttributes(attribute.String("WithAttributes-parentFunctionAttributes-k", "0001")))
	// ... and after creation
	opentelemetryclient.SetAttributes(spanCtx, attribute.Bool("isTrue", true), attribute.String("parentFunctionAttributes-WithAttributes-SetAttributes", "hi!"))
	defer opentelemetryclient.Finish(spanCtx)

}

func childFunction(ctx context.Context) {
	// Create a span to track `childFunction()` - this is a nested span whose parent is `parentSpan`
	spanCtx := opentelemetryclient.StartSpan(ctx, "child")
	defer opentelemetryclient.Finish(spanCtx)

	// do work here, when this function returns, childSpan will complete.
	childFunctionAttributes(spanCtx)
}

func childFunctionAttributes(ctx context.Context) {
	// setting attributes at creation...
	spanCtx := opentelemetryclient.StartSpan(ctx, "childFunctionAttributes-attributesAtCreation", trace.WithAttributes(attribute.String("WithAttributes-childFunctionAttributes-k", "0001")))
	// ... and after creation
	opentelemetryclient.SetAttributes(spanCtx, attribute.Bool("isTrue", true), attribute.String("childFunctionAttributes-WithAttributes-SetAttributes", "hi!"))
	defer opentelemetryclient.Finish(spanCtx)
}
