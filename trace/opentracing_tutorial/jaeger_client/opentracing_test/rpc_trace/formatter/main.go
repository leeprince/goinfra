package main

import (
    "fmt"
    "github.com/leeprince/goinfra/plog"
    "github.com/leeprince/goinfra/trace/opentracing/jaeger_client"
    "log"
    "net/http"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/5 下午4:12
 * @Desc:
 */

const (
    serviceName = "opentracing_test-rpc_trace-formatter"
)

func main() {
    jaeger_client.InitTracer(serviceName)
    defer jaeger_client.Close()
    
    http.HandleFunc("/format", func(w http.ResponseWriter, r *http.Request) {
        spanCtx, err := jaeger_client.HTTPServer(r.Context(), "formatter@http.HandleFunc", r.Header)
        if err != nil {
            plog.Fatal("jaeger_client.HTTPServer err:", err)
        }
        defer jaeger_client.Finish(spanCtx)
        plog.WithFiledLogID(jaeger_client.TraceID(spanCtx)).Info("spanCtx TraceID")
        
        jaeger_client.LogKV(spanCtx, "formatter@http.HandleFunc@LogKV001", "println")
        plog.WithFiledLogID(jaeger_client.TraceID(spanCtx)).Info("")
        
        helloTo := r.FormValue("helloTo")
        helloStr := fmt.Sprintf("Hello, %s!", helloTo)
        w.Write([]byte(helloStr))
    })
    
    log.Fatal(http.ListenAndServe(":8111", nil))
}
