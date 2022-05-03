package main

import (
    "context"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/attribute"
    "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
    "go.opentelemetry.io/otel/sdk/resource"
    "go.opentelemetry.io/otel/sdk/trace"
    semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
    "io"
    "log"
    "os"
    "os/signal"
)

const serviceName = "opentelemetry_client-get_started"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/6 下午5:13
 * @Desc:
 */

func main() {
    l := log.New(os.Stdout, "", 0)
    
    // --- otel
    // io.Writer
    // Write telemetry data to os.Stdout
    /*f := os.Stdout
    exp, err := newExporter(f)
    if err != nil {
        l.Fatal(err)
    }*/
    // io.Writer
    // Write telemetry data to a file.
    f, err := os.Create("traces.txt")
    if err != nil {
        l.Fatal(err)
    }
    
    exp, err := newExporter(f)
    if err != nil {
        l.Fatal(err)
    }
    
    tp := trace.NewTracerProvider(
        trace.WithBatcher(exp),
        trace.WithResource(newResource()),
    )
    // Handle shutdown properly so nothing leaks.
    defer func() {
        if err := tp.Shutdown(context.Background()); err != nil {
            l.Fatal(err)
        }
    }()
    otel.SetTracerProvider(tp)
	// --- otel -end
    
    sigCh := make(chan os.Signal, 1)
    signal.Notify(sigCh, os.Interrupt)
    
    errCh := make(chan error)
    app := NewApp(os.Stdin, l)
    go func() {
        errCh <- app.Run(context.Background())
    }()
    
    select {
    case <-sigCh:
        l.Println("\ngoodbye")
        return
    case err := <-errCh:
        if err != nil {
            l.Fatal(err)
        }
    }
}

// newResource returns a resource describing this application.
func newResource() *resource.Resource {
    r, _ := resource.Merge(
        resource.Default(),
        resource.NewWithAttributes(
            semconv.SchemaURL,
            semconv.ServiceNameKey.String(serviceName),
            semconv.ServiceVersionKey.String("v0.1.0"),
            attribute.String("environment", "demo"),
        ),
    )
    return r
}

// newExporter returns a console exporter.
func newExporter(w io.Writer) (trace.SpanExporter, error) {
    return stdouttrace.New(
        stdouttrace.WithWriter(w),
        // Use human-readable output.
        stdouttrace.WithPrettyPrint(),
        // Do not print timestamps for the demo.
        // stdouttrace.WithoutTimestamps(),
    )
}
