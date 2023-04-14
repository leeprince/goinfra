package main

import (
	"fmt"
	"github.com/leeprince/goinfra/plog"
	"github.com/leeprince/goinfra/trace/opentracing/jaegerclient"
	"log"
	"net/http"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/5 下午4:12
 * @Desc:
 */

const (
	serviceName = "opentracing_test-rpc_trace_baggage-formatter"
)

func main() {
	jaegerclient.InitTracer(serviceName)
	defer jaegerclient.Close()

	http.HandleFunc("/format", func(w http.ResponseWriter, r *http.Request) {
		spanCtx, err := jaegerclient.ExtractTraceHTTPServer(r.Context(), "publisher@http.HandleFunc", r.Header)
		if err != nil {
			plog.Fatal("jaeger_client.ExtractTraceHTTPServer err:", err)
		}
		defer jaegerclient.Finish(spanCtx)
		plog.LogID(jaegerclient.TraceID(spanCtx)).Info("spanCtx TraceID")

		// 使用 span 的 Baggage 功能
		seq := jaegerclient.BaggageItem(spanCtx, "seq")
		println("BaggageItem:seq", seq)

		jaegerclient.LogKV(spanCtx, "formatter@http.HandleFunc@LogKV001", "println")

		helloTo := r.FormValue("helloTo")
		helloStr := fmt.Sprintf("Hello, %s!", helloTo)
		w.Write([]byte(helloStr))
	})

	log.Fatal(http.ListenAndServe(":8121", nil))
}
