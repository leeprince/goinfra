package main

import (
    "github.com/leeprince/goinfra/plog"
    "github.com/leeprince/goinfra/trace/opentracing/jaeger_client"
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
    jaeger_client.InitTracer(serviceName)
    defer jaeger_client.Close()
    
    http.HandleFunc("/publish", func(w http.ResponseWriter, r *http.Request) {
        spanCtx, err := jaeger_client.HTTPServer(r.Context(), "publisher@http.HandleFunc", r.Header)
        if err != nil {
            plog.Fatal("jaeger_client.HTTPServer err:", err)
        }
        defer jaeger_client.Finish(spanCtx)
        plog.WithFiledLogID(jaeger_client.TraceID(spanCtx)).Info("spanCtx TraceID")
        
        // 使用 span 的 Baggage 功能
        seq := jaeger_client.BaggageItem(spanCtx, "seq")
        println("BaggageItem:seq", seq)
        
        jaeger_client.LogKV(spanCtx, "publisher@http.HandleFunc@LogKV001", "println")
        
        helloStr := r.FormValue("helloStr")
        println(helloStr)
    })
    
    log.Fatal(http.ListenAndServe(":8122", nil))
}