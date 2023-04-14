package main

import (
	"context"
	"fmt"
	"github.com/leeprince/goinfra/consts"
	"github.com/leeprince/goinfra/trace/opentracing/opentelemetry_client"
	"log"
	"net/http"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/5 下午4:12
 * @Desc:
 */

const (
	serviceName = "opentracing_test_app-rpc_trace-formatter"
	env         = consts.ENVLocal
)

func main() {
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

	handler := http.HandlerFunc(formatHandler)
	// manualHttpHandler 加入 opentelemetry_client.NewHandler 拦截器中（中间件）
	wrappedHandler := opentelemetryclient.NewHandler(handler, "NewHandler.formatHandler")
	http.Handle("/format", wrappedHandler)

	log.Fatal(http.ListenAndServe(":8202", nil))
}

func formatHandler(w http.ResponseWriter, r *http.Request) {
	spanCtx := opentelemetryclient.StartSpan(r.Context(), "formatHandler")
	defer opentelemetryclient.Finish(spanCtx)
	opentelemetryclient.TagString(spanCtx, "formatter@TagString01", "println")
	seq := opentelemetryclient.BaggageItem(spanCtx, "seq")
	opentelemetryclient.PlogInfo(spanCtx, "formatHandler@opentelemetry_client.BaggageItem seq:", seq)

	helloTo := r.FormValue("helloTo")
	helloStr := fmt.Sprintf("Hello, %s!", helloTo)
	w.Write([]byte(helloStr))

	opentelemetryclient.PlogInfo(spanCtx, "formatHandler plog.Info end")
}
