// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
    "context"
    "github.com/leeprince/goinfra/http/httpcli"
    "github.com/leeprince/goinfra/plog"
    "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
    "go.opentelemetry.io/otel/attribute"
    "go.opentelemetry.io/otel/baggage"
    "go.opentelemetry.io/otel/example/passthrough/handler"
    "go.opentelemetry.io/otel/exporters/jaeger"
    "go.opentelemetry.io/otel/sdk/resource"
    "log"
    "net/http"
    "net/url"
    
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/propagation"
    sdktrace "go.opentelemetry.io/otel/sdk/trace"
    semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
    "go.opentelemetry.io/otel/trace"
)

const (
    serviceName = "5passthrough"
)

var (
    tp *sdktrace.TracerProvider
)

func main() {
    ctx := context.Background()
    
    initPassthroughGlobals()
    tp = nonGlobalTracer(serviceName)
    defer func() { _ = tp.Shutdown(ctx) }()
    
    // This is roughly what an instrumented http client does.
    plog.Info("The \"make outer request\" span should be recorded, because it is recorded with a Tracer from the SDK TracerProvider")
    var span trace.Span
    ctx, span = tp.Tracer("example/passthrough/outer").Start(ctx, "make outer request")
    defer span.End()
    
    // 设置 Baggage 的 member
    member, err := baggage.NewMember("seq", "prince-seq-20220429")
    if err == nil {
        newBaggage, err := baggage.FromContext(ctx).SetMember(member)
        if err != nil {
            plog.Panic("baggage.FromContext(ctx).SetMember err", err)
        }
        ctx = baggage.ContextWithBaggage(ctx, newBaggage)
    } else {
        plog.Panic("baggage.NewMember err", err)
    }
    
    // make an initial http request
    r, err := http.NewRequest("", "", nil)
    if err != nil {
        panic(err)
    }
    r = r.WithContext(ctx)
    otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(r.Header))
    
    backendFunc := func(r *http.Request) {
        // This is roughly what an instrumented http server does.
        ctx := otel.GetTextMapPropagator().Extract(r.Context(), propagation.HeaderCarrier(r.Header))
        plog.Info("The \"handle inner request\" span should be recorded, because it is recorded with a Tracer from the SDK TracerProvider")
        ctx, span := tp.Tracer("example/passthrough/inner").Start(ctx, "handle inner request")
        defer span.End()
        
        //  获取 Baggage 的 member 对应 key 的值
        seq := baggage.FromContext(ctx).Member("seq").Value()
        plog.Info("backendFunc>>>>>BaggageItem seq:", seq)
        
        // Do "backend work"
        // time.Sleep(time.Second)
    }
    // This handler will be a passthrough, since we didn't set a global TracerProvider
    passthroughHandler := handler.New(backendFunc)
    passthroughHandler.HandleHTTPReq(r)
    
    RPCTraceFormatterHttpCli(ctx)
    RPCTraceFormatteOTELrHttpGet(ctx)
}

func RPCTraceFormatterHttpCli(ctx context.Context) {
    // - 传递上下文 context 代替将 span 作为每个函数的第一个参数
    ctx, span := tp.Tracer("RPCTraceFormatter-tracer").Start(ctx, "RPCTraceFormatter start")
    defer span.End()
    
    span.SetAttributes(attribute.String("RPCTraceFormatter:SetAttributes", "001"))
    
    //  获取 Baggage 的 member 对应 key 的值
    seq := baggage.FromContext(ctx).Member("seq").Value()
    plog.Info("RPCTraceFormatterHttpCli>>>>>BaggageItem seq:", seq)
    
    params := url.Values{
        "helloTo": []string{"prince"},
    }
    urlPath := "http://127.0.0.1:8203/format/headerExtractHandler?" + params.Encode()
    
    // make an initial http request
    r, err := http.NewRequest(http.MethodGet, urlPath, nil)
    if err != nil {
        panic(err)
    }
    
    r = r.WithContext(ctx)
    otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(r.Header))
    
    bodyByte, err := httpcli.Do(r)
    if err != nil {
        plog.Info("RPCTraceFormatter httpcli.Do err:", err)
        return
    }
    
    plog.Info("bodyString:", string(bodyByte))
    plog.Info("RPCTraceFormatter TraceID:", trace.SpanContextFromContext(ctx).TraceID().String())
}

func RPCTraceFormatteOTELrHttpGet(ctx context.Context) {
    // - 传递上下文 context 代替将 span 作为每个函数的第一个参数
    ctx, span := tp.Tracer("RPCTraceFormatteOTELrHttpGet-tracer").Start(ctx, "RPCTraceFormatteOTELrHttpGet start")
    defer span.End()
    
    span.SetAttributes(attribute.String("RPCTraceFormatteOTELrHttpGet:SetAttributes", "001"))
    
    //  获取 Baggage 的 member 对应 key 的值
    seq := baggage.FromContext(ctx).Member("seq").Value()
    plog.Info("RPCTraceFormatteOTELrHttpGet>>>>>BaggageItem seq:", seq)
    
    params := url.Values{
        "helloTo": []string{"prince"},
    }
    urlPath := "http://127.0.0.1:8203/format/wrappedHandler?" + params.Encode()
    
    resp, err := otelhttp.Get(ctx, urlPath)
    if err != nil {
        plog.Info("RPCTraceFormatteOTELrHttpGet otelhttp.Get err:", err)
        return
    }
    bodyByte, err := httpcli.ResponseToBytes(resp)
    if err != nil {
        plog.Info("RPCTraceFormatteOTELrHttpGet httpcli.ResponseToBytes err:", err)
        return
    }
    
    plog.Info("bodyString:", string(bodyByte))
    plog.Info("RPCTraceFormatteOTELrHttpGet TraceID:", trace.SpanContextFromContext(ctx).TraceID().String())
}

func initPassthroughGlobals() {
    // We explicitly DO NOT set the global TracerProvider using otel.SetTracerProvider().
    // The unset TracerProvider returns a "non-recording" span, but still passes through context.
    plog.Info("Register a global TextMapPropagator, but do not register a global TracerProvider to be in \"passthrough\" mode.")
    plog.Info("The \"passthrough\" mode propagates the TraceContext and Baggage, but does not record spans.")
    otel.SetTextMapPropagator(
        // propagation.TraceContext{}, // 使用分布式链路追踪的 Trace Span 功能（保证分布式传播 Trace Span：TraceID 统一 ）
        // propagation.Baggage{}, // 使用分布式链路追踪的分布式 Baggate 功能（保证分布式传播 Baggate 功能）
        propagation.NewCompositeTextMapPropagator(
            propagation.TraceContext{},
            propagation.Baggage{},
        ),
    )
}

// nonGlobalTracer creates a trace provider instance for testing, but doesn't
// set it as the global tracer provider
func nonGlobalTracer(serviceName string) *sdktrace.TracerProvider {
    var err error
    exp, err := jaeger.New(jaeger.WithCollectorEndpoint())
    // exp, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
    if err != nil {
        log.Panicf("failed to initialize stdouttrace exporter %v\n", err)
    }
    resourceNew, err := newResource(serviceName, "local")
    if err != nil {
        log.Panicf("failed to initialize stdouttrace exporter %v\n", err)
    }
    bsp := sdktrace.NewBatchSpanProcessor(exp)
    tp := sdktrace.NewTracerProvider(
        sdktrace.WithResource(resourceNew),
        sdktrace.WithSampler(sdktrace.AlwaysSample()),
        sdktrace.WithSpanProcessor(bsp),
    )
    
    otel.SetTracerProvider(tp)
    
    return tp
}

func newResource(serviceName, env string) (*resource.Resource, error) {
    return resource.Merge(
        resource.Default(),
        resource.NewWithAttributes(
            semconv.SchemaURL,
            semconv.ServiceNameKey.String(serviceName),
            attribute.String("environment", env),
        ),
    )
}