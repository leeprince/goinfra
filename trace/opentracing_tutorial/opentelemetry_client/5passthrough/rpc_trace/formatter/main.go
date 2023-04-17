package main

import (
	"context"
	"fmt"
	"github.com/leeprince/goinfra/consts"
	"github.com/leeprince/goinfra/plog"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
	"log"
	"net/http"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/5 下午4:12
 * @Desc:
 */

const (
	serviceName = "5passthrough_trace-formatter"
	env         = consts.ENV_LOCAL
)

var (
	tp *sdktrace.TracerProvider
)

func main() {
	ctx := context.Background()

	initPassthroughGlobals()
	tp = nonGlobalTracer(serviceName)
	defer func() { _ = tp.Shutdown(ctx) }()

	// --- 手动通过 otel.GetTextMapPropagator().Extract 提取请求头（header）已注入 span 信息
	headerExtractHandler := http.HandlerFunc(headerExtractHandler)
	http.Handle("/format/headerExtractHandler", headerExtractHandler)

	// --- 通过加入 otelhttp.NewHandler 拦截器中（中间件）中实现从请求头（header）中提取（Extract）已注入 span 信息
	otelHttpHandler := http.HandlerFunc(OTELHttpHandler)
	wrappedHandler := otelhttp.NewHandler(otelHttpHandler, "otelhttp.NewHandler-OTELHttpHandler")
	http.Handle("/format/wrappedHandler", wrappedHandler)

	log.Fatal(http.ListenAndServe(":8203", nil))
}

// r.header 已注入 span 信息
func headerExtractHandler(w http.ResponseWriter, r *http.Request) {
	plog.Info("headerExtractHandler>>>>>header:", r.Header)
	requestCtx := r.Context()
	traceID := trace.SpanContextFromContext(requestCtx).TraceID().String()
	plog.Info("headerExtractHandler>>>>>traceID:", traceID)

	headerCarrier := propagation.HeaderCarrier(r.Header)
	ctx := otel.GetTextMapPropagator().Extract(requestCtx, headerCarrier)
	ctx, span := tp.Tracer("headerExtractHandler-tracer").Start(ctx, "headerExtractHandler start")
	defer span.End()

	plog.Info("headerExtractHandler>>>>>ctx traceID:", trace.SpanContextFromContext(ctx).TraceID().String())

	//  获取 Baggage 的 member 对应 key 的值
	seq := baggage.FromContext(ctx).Member("seq").Value()
	plog.Info("headerExtractHandler>>>>>BaggageItem seq:", seq)

	helloTo := r.FormValue("helloTo")
	helloStr := fmt.Sprintf("Hello, %s!", helloTo)
	w.Write([]byte(helloStr))
}

// r.header 已注入 span 信息
func OTELHttpHandler(w http.ResponseWriter, r *http.Request) {
	plog.Info("OTELHttpHandler>>>>>header:", r.Header)
	requestCtx := r.Context()
	traceID := trace.SpanContextFromContext(requestCtx).TraceID().String()
	plog.Info("OTELHttpHandler>>>>>traceID:", traceID)

	ctx, span := tp.Tracer("OTELHttpHandler-tracer").Start(requestCtx, "OTELHttpHandler start")
	defer span.End()

	//  获取 Baggage 的 member 对应 key 的值
	seq := baggage.FromContext(ctx).Member("seq").Value()
	plog.Info("OTELHttpHandler>>>>>BaggageItem seq:", seq)

	helloTo := r.FormValue("helloTo")
	helloStr := fmt.Sprintf("Hello, %s!", helloTo)
	w.Write([]byte(helloStr))

	plog.Info("OTELHttpHandler plog.Info end")
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
