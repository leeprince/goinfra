package main

import (
	"github.com/leeprince/goinfra/plog"
	"github.com/leeprince/goinfra/trace/opentracing/jaegerclient"
	"log"
	"net/http"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/5 下午4:11
 * @Desc:
 */

const (
	serviceName = "princeJaeger-lesson03-rpc-trace-baggage-publisher"
)

func main() {
	jaegerclient.InitTracer(serviceName)
	defer jaegerclient.Close()

	http.HandleFunc("/publish", func(w http.ResponseWriter, r *http.Request) {
		spanCtx, err := jaegerclient.ExtractTraceHTTPServer(r.Context(), "publisher@http.HandleFunc", r.Header)
		if err != nil {
			plog.Fatal("jaeger_client.ExtractTraceHTTPServer err:", err)
		}
		defer jaegerclient.Finish(spanCtx)
		plog.LogID(jaegerclient.TraceID(spanCtx)).Info("spanCtx TraceID")

		// 使用 span 的 Baggage 功能
		seq := jaegerclient.BaggageItem(spanCtx, "seq")
		println("BaggageItem:seq", seq)

		jaegerclient.LogKV(spanCtx, "publisher@http.HandleFunc@LogKV001", "println")

		helloStr := r.FormValue("helloStr")
		println(helloStr)
	})

	log.Fatal(http.ListenAndServe(":8122", nil))
}
